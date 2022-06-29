package waioNEAT

import (
	hash "crypto/sha256"
	bs64 "encoding/base64"
	"fmt"
)

// 'Gene' describing a Neuron
type NeuronGene struct {
	Id         int
	Type       int    `json:"type"`
	Param      string `json:"param"`
	Activation string `json:"activation"`
	Innov      string
}

func (ng *NeuronGene) Checksum() string {
	h := hash.New()
	h.Write([]byte(fmt.Sprintf("%d%s%s", ng.Type, ng.Param, ng.Activation)))
	return bs64.StdEncoding.EncodeToString([]byte(h.Sum(nil)))
}

func MakeNeuronGene(Id int, Type int, Param string, Activation string) *NeuronGene {
	r := &NeuronGene{
		Id:         Id,
		Type:       Type,
		Param:      Param,
		Activation: Activation,
	}
	r.Innov = r.Checksum()
	return r
}

type SynapseGene struct {
	Sender  int     `json:"sender"`
	Reciver int     `json:"reciver"`
	Weight  float64 `json:"weight"`
	Active  bool    `json:"active"`
	Innov   string
}

func MakeSynapseGene(Sender int, Reciver int, Weight float64, Active bool) *SynapseGene {
	r := &SynapseGene{
		Sender:  Sender,
		Reciver: Reciver,
		Weight:  Weight,
		Active:  Active,
	}
	r.Innov = r.Checksum()
	return r
}

func (sg *SynapseGene) Checksum() string {
	h := hash.New()
	h.Write([]byte(fmt.Sprintf("%d%d", sg.Sender, sg.Reciver)))
	return bs64.StdEncoding.EncodeToString([]byte(h.Sum(nil)))
}

type Genome struct {
	NeuronGenes  []*NeuronGene
	SynapseGenes []*SynapseGene
	id_count     int
}

func (g *Genome) nextId() int {
	g.id_count += 1
	return g.id_count
}

func (g *Genome) Copy(t *Genome) {
	g.id_count = t.id_count
	g.NeuronGenes = make([]*NeuronGene, 0, len(t.NeuronGenes))
	for _, ng := range t.NeuronGenes {
		g.NeuronGenes = append(g.NeuronGenes, MakeNeuronGene(ng.Id, ng.Type, ng.Param, ng.Activation))
	}
	g.SynapseGenes = make([]*SynapseGene, 0, len(t.SynapseGenes))
	for _, sg := range t.SynapseGenes {
		g.SynapseGenes = append(g.SynapseGenes, MakeSynapseGene(sg.Sender, sg.Reciver, sg.Weight, sg.Active))
	}
}

// Creates a Genome with the given input and output neuron genes.
// No hidden neuron genes are included.
// Accepts a map that relates each parameter with it's chance of apearing.
// Chance goes from 0 to 1.
func MakeGenome(InParams map[string]float64, OutParams map[string]float64) *Genome {
	genome := &Genome{
		id_count: -1,
	}

	actName := ""
	for param, chance := range InParams {
		if RndGen.Float64() <= chance {
			if RndActivation {
				actName = ActivationSlice[RndGen.Intn(len(ActivationSlice))].name

			} else {
				actName = defInputAct
			}
			n := &NeuronGene{
				Id:         genome.nextId(),
				Type:       INPUT,
				Param:      param,
				Activation: actName,
				Innov:      "",
			}
			n.Innov = n.Checksum()
			genome.NeuronGenes = append(genome.NeuronGenes, n)
		}
	}

	for param, chance := range OutParams {
		if RndGen.Float64() <= chance {
			if RndActivation {
				actName = ActivationSlice[RndGen.Intn(len(ActivationSlice))].name

			} else {
				actName = defOutputAct
			}
			n := &NeuronGene{
				Id:         genome.nextId(),
				Type:       OUTPUT,
				Param:      param,
				Activation: actName,
				Innov:      "",
			}
			n.Innov = n.Checksum()
			genome.NeuronGenes = append(genome.NeuronGenes, n)
		}
	}
	return genome
}

/*
Accepts two genomes and returns the union of both.
For equivalent genes, there is a 50% chance for each ancestor gene.
(Equivalent genes still differ in some factors)
*/
func Crossover(g1 *Genome, g2 *Genome) *Genome {
	g := &Genome{}

	// Make g1 the larger genome
	var swap *Genome = nil
	if len(g1.NeuronGenes) < len(g2.NeuronGenes) {
		swap = g1
		g1 = g2
		g2 = swap
	}

	// Get the larger genome's neuron genes
	bNeuronGenes := make(map[string]*NeuronGene, len(g1.NeuronGenes))
	for _, ng := range g1.NeuronGenes {
		bNeuronGenes[ng.Checksum()] = ng
	}

	// For each of smaller genome's neuron genes...
	for _, sng := range g2.NeuronGenes {
		// if its present on the larger genome...
		if bng, f := bNeuronGenes[sng.Checksum()]; f {
			if c := RndGen.Float64(); c <= 0.5 {
				// there is a 50% chance for each equivalent
				g.NeuronGenes = append(g.NeuronGenes, sng)
			} else {
				g.NeuronGenes = append(g.NeuronGenes, bng)
			}
		}
	}

	// Repeat the proces for the synaptic genes
	swap = nil
	if len(g1.SynapseGenes) < len(g2.SynapseGenes) {
		swap = g1
		g1 = g2
		g2 = swap
	}

	bSynapseGenes := make(map[string]*SynapseGene, len(g1.SynapseGenes))
	for _, sg := range g1.SynapseGenes {
		bSynapseGenes[sg.Checksum()] = sg
	}

	for _, ssg := range g2.SynapseGenes {
		if bsg, f := bSynapseGenes[ssg.Checksum()]; f {
			if c := RndGen.Float64(); c <= 0.5 {
				g.SynapseGenes = append(g.SynapseGenes, ssg)
			} else {
				g.SynapseGenes = append(g.SynapseGenes, bsg)
			}
		}
	}
	return g
}

/*
An existing connection is split and the new node placed where the old connection used to be. The
old connection is disabled and two new connections are added to the genome. The new connection
leading into the new node receives a weight of 1, and the new connection leading out receives the
same weight as the old connection. This method of adding nodes was chosen in order to minimize the
initial effect of the mutation.
(K. O. Stanley and R. Miikkulainen; Evolutionary Computation Volume 10, Number 2)
*/
func (g *Genome) MutateAddNode() {
	if len(g.SynapseGenes) <= 0 {
		return
	}

	// Get a random Activation and Synapse
	rActivation := ""
	if RndActivation {
		rActivation = ActivationSlice[RndGen.Intn(len(ActivationSlice))].name

	} else {
		rActivation = defHiddenAct
	}
	rSynapse := g.SynapseGenes[RndGen.Intn(len(g.SynapseGenes))]
	rSynapse.Active = false

	// Create the new Hidden Neuron gene
	n := &NeuronGene{
		Id:         g.nextId(),
		Type:       HIDDEN,
		Param:      "",
		Activation: rActivation,
		Innov:      "",
	}
	// Create the first half of the splitted synapse
	s1 := &SynapseGene{
		Sender:  rSynapse.Sender,
		Reciver: n.Id,
		Weight:  1,
		Innov:   "",
		Active:  true, // Leave to chance too?
	}
	// Create the second half of the splitted synapse
	s2 := &SynapseGene{
		Sender:  n.Id,
		Reciver: rSynapse.Reciver,
		Weight:  rSynapse.Weight,
		Innov:   "",
		Active:  true, // Leave to chance too?
	}

	n.Innov = n.Checksum()
	s1.Innov = s1.Checksum()
	s2.Innov = s2.Checksum()

	// Add the new genes to the genome
	g.NeuronGenes = append(g.NeuronGenes, n)
	g.SynapseGenes = append(g.SynapseGenes, s1, s2)
}

// A single new connection gene with a random weight is added connecting two previously unconnected
// nodes.
func (g *Genome) MutateAddSynapse() {
	n1 := g.NeuronGenes[RndGen.Intn(len(g.NeuronGenes))]
	n2 := g.NeuronGenes[RndGen.Intn(len(g.NeuronGenes))]

	// Get 2 andom Neuron Genes until both are valid
	// Valid: Diferent nodes + the reciver can't be an input and the sender can't be an output node
	for n1.Id == n2.Id || n1.Type == OUTPUT || n2.Type == INPUT {
		n1 = g.NeuronGenes[RndGen.Intn(len(g.NeuronGenes))]
		n2 = g.NeuronGenes[RndGen.Intn(len(g.NeuronGenes))]
	}

	// If the found nodes are already connected the mutation doesn't take place
	for _, sy := range g.SynapseGenes {
		if sy.Sender == n1.Id && sy.Reciver == n2.Id {
			return
		}
	}
	sg := &SynapseGene{
		Sender:  n1.Id,
		Reciver: n2.Id,
		Weight:  RndGen.NormFloat64(),
		Active:  true,
		Innov:   "",
	}
	sg.Innov = sg.Checksum()
	g.SynapseGenes = append(g.SynapseGenes, sg)
}

/*
Implement mutations:
	- Weight/s mutation
*/

/*
REFERENCE:
http://nn.cs.utexas.edu/downloads/papers/stanley.ec02.pdf

Implement:
	- [ ] Speciation:
		This factor is calculated by taking in account dijoint and excess genes and wheigt
		mean, and it's used to calculate the compatibility of genomes so they can be
		clustered in species, so only genomes with enough compatibility shoud be able to
		reproduce. This helps new species, that may end up being more fit than others but
		aren't optimized yet, to survive.

		Compatibility Distrance = (c_1 · E)/N + (c_2 · D)/N + c_3 · Ŵ
		Where...
			c_i : Importance factor
			E   : Number of Excess genes
			D	: Number of Disjoint genes
			Ŵ	: Average weight differences of matching genes
			N	: number of genes in the larger genome

	- [ ] Minimized Dimensionality:
		NEAT starts with a uniform population and without hidden nodes. New changes are
		introduced by _Growth_ and protected by _Speciation_ but ultimately, fitness
		evaluation determines if those changes are useful, therefore complexity increases
		when needed.
*/

// New change in 0.0.1 branch
