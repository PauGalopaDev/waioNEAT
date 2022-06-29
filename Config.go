package waioNEAT

import (
	"math/rand"
	"time"
)

var (
	defInputAct           string
	defHiddenAct          string
	defOutputAct          string
	CrossoverParentChance float64
	RndActivation         bool
	RndGen                *rand.Rand
)

func Init() {
	defInputAct = "linear"
	defHiddenAct = "sigmoid"
	defOutputAct = "tanh"
	CrossoverParentChance = 0.5
	RndActivation = false
	RndGen = rand.New(rand.NewSource(time.Hour.Nanoseconds()))
}

/*
To implement
Integrated Bias and Bias as Input + Qty of Binputs (Not exclusive)
*/
