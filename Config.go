package waioNEAT

import (
	"math/rand"
	"time"
)

var (
	defInputAct            string
	defHiddenAct           string
	defOutputAct           string
	IntegratedBias         bool
	AlwaysAddBiasNeuron    bool
	CrossLargeParentChance float64
	Seed                   int64
	RndActivationHidden    bool
	RndActivationInput     bool
	RndActivationOutput    bool
	RndGen                 *rand.Rand
)

func Init() {
	defInputAct = "linear"
	defHiddenAct = "sigmoid"
	defOutputAct = "tanh"
	AlwaysAddBiasNeuron = false
	IntegratedBias = false
	CrossLargeParentChance = 0.5 // 0.3 would mean 30% chance for the larger and 70% for the smaller
	RndActivationHidden = false
	RndActivationInput = false
	RndActivationOutput = false
	Seed = time.Hour.Nanoseconds()
	RndGen = rand.New(rand.NewSource(Seed))
}

/*
Configuration format:
{
	"defInputAct":"linear",
	"defHiddenAct":"sigmoid",
	"defOutputAct":"tanh",
	"IntegratedBias":false,
	"CrossLargeParentChance":0.5,
	"RndActivationHidden" = false
	"RndActivationInput" = false
	"RndActivationOutput" = false
	"Seed":null,
	""
}
*/
