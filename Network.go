package pneat

type Network struct {
	Neurons map[int]*Neuron
	Input   map[string]*Neuron
	Output  map[string]*Neuron
}

func (n *Network) Feed() {
	for _, n := range n.Output {
		n.Activate()
	}
	for _, n := range n.Neurons {
		n.Value = 0.0
		n.active = false
	}
}

func MakeNetwork(g *Genome) *Network {
	network := &Network{
		Neurons: make(map[int]*Neuron, len(g.NeuronGenes)),
		Input:   make(map[string]*Neuron),
		Output:  make(map[string]*Neuron),
	}

	for _, ng := range g.NeuronGenes {
		n := MakeNeuron(ng)

		if n.Type == INPUT {
			network.Input[n.Param] = n
		} else if n.Type == OUTPUT {
			network.Output[n.Param] = n
		}
		network.Neurons[n.Id] = n
	}

	for _, sg := range g.SynapseGenes {
		network.Neurons[sg.Reciver].Connections[network.Neurons[sg.Sender]] = sg.Weight
	}

	return network
}
