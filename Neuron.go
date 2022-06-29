package waioNEAT

const (
	INPUT  int = 0
	HIDDEN int = 1
	OUTPUT int = 2
)

type Neuron struct {
	Id          int
	Type        int
	Bias        float64
	Value       float64
	Param       string
	Connections map[*Neuron]float64
	Activation  ActivationFn

	active bool
}

func MakeNeuron(g *NeuronGene) *Neuron {
	act := ""
	switch g.Type {
	case INPUT:
		act = defInputAct
	case OUTPUT:
		act = defOutputAct
	case HIDDEN:
		act = defHiddenAct

	}

	b := 0.0
	if IntegratedBias {
		b = RndGen.Float64()
	}
	n := &Neuron{
		Id:          g.Id,
		Type:        g.Type,
		Param:       g.Param,
		Value:       0.0,
		Bias:        b,
		Connections: make(map[*Neuron]float64),
		Activation:  ActivationMap[act],
		active:      false,
	}
	return n
}

func (neuron *Neuron) Activate() float64 {
	if neuron.active || neuron.Type == INPUT {
		return neuron.Value
	}

	neuron.active = true
	var sum float64 = 0

	for sender, weight := range neuron.Connections {
		sum += sender.Activate() * weight
	}

	sum += neuron.Bias
	neuron.Value = neuron.Activation.fn(sum)
	return neuron.Value
}
