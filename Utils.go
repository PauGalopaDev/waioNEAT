package waioNEAT

import (
	"fmt"
)

func (n *Neuron) String() string {
	synapses := ""
	for k, v := range n.Connections {
		synapses += fmt.Sprintf("(-- %f -->%d)\n", v, k.Id)
	}
	return fmt.Sprintf(
		"neuron:(\n"+
			"\tid:%d\n"+
			"\ttype:%d\n"+
			"\tparam:%s\n"+
			"\tvalue:%f\n"+
			"\tsynapses:%s\n"+
			"\tactivation:%s\n)",
		n.Id, n.Type, n.Param, n.Value, synapses, n.Activation.name)
}

func (n *Network) String() string {
	result := "Network:\n\tInputs:\n"
	for _, i := range n.Input {
		result += fmt.Sprintf("%s: %f\n", i.Param, i.Value)
	}

	result += "\n\tOutputs:\n"
	for _, o := range n.Output {
		result += fmt.Sprintf("%s: %f\n", o.Param, o.Value)
	}
	return result
}

func (g *Genome) String() string {
	result := "Genome: \n\tNeurons:\n"
	for _, ng := range g.NeuronGenes {
		result += fmt.Sprintf("%v\n", ng)
	}

	result += "\n\tSynapses:\n"
	for _, sg := range g.SynapseGenes {
		result += fmt.Sprintf("%v\n", sg)
	}
	return result
}

func (sg *SynapseGene) String() string {
	active := "X"
	if sg.Active {
		active = ""
	}
	return fmt.Sprintf("%s(%d -- %f --> %d)%s", active, sg.Sender, sg.Weight, sg.Reciver, active)
}

func (ng *NeuronGene) String() string {
	t := "hidden"
	if ng.Type == INPUT {
		t = "input"
	} else if ng.Type == OUTPUT {
		t = "output"
	}
	return fmt.Sprintf("(%d, %s, %s, %s, %s)", ng.Id, t, ng.Param, ng.Activation, ng.Innov)
}
