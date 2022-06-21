package waioNEAT

import "math"

type ActivationFn struct {
	name string
	fn   func(...float64) float64
}

func BinStep(x ...float64) float64 {
	if x[0] >= 1 {
		return 1
	} else {
		return 0
	}
}

func ReLU(x ...float64) float64 {
	return math.Max(x[0], 0)
}

func LeakyReLU(x ...float64) float64 {
	return math.Max(x[0], x[0]*0.01)
}

func Sigmoid(x ...float64) float64 {
	return 1 / (1 + math.Pow(math.E, -x[0]))
}

func Tanh(x ...float64) float64 {
	return math.Tanh(x[0])
}

var ActivationMap = map[string]ActivationFn{
	"binstep":   {"binstep", BinStep},
	"relu":      {"relu", ReLU},
	"leakyrelu": {"leakyrelu", LeakyReLU},
	"sigmoid":   {"sigmoid", Sigmoid},
	"tanh":      {"tanh", Tanh},
}

var ActivationSlice = []ActivationFn{
	{"binstep", BinStep},
	{"relu", ReLU},
	{"leakyrelu", LeakyReLU},
	{"sigmoid", Sigmoid},
	{"tanh", Tanh},
}

func RegisterActivation(name string, fn func(...float64) float64) {
	ActivationMap[name] = ActivationFn{name, fn}
	ActivationSlice = append(ActivationSlice, ActivationFn{name, fn})
}
