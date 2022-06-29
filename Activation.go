package waioNEAT

import "math"

type ActivationFn struct {
	name string
	fn   func(float64) float64
}

func Linear(x float64) float64 {
	return x
}

func BinStep(x float64) float64 {
	if x >= 1 {
		return 1
	} else {
		return 0
	}
}

func ReLU(x float64) float64 {
	return math.Max(x, 0)
}

func LeakyReLU(x float64) float64 {
	return math.Max(x, x*0.01)
}

func Sigmoid(x float64) float64 {
	return 1 / (1 + math.Pow(math.E, -x))
}

func Tanh(x float64) float64 {
	return math.Tanh(x)
}

var ActivationMap = map[string]ActivationFn{
	"linear":    {"linear", Linear},
	"binstep":   {"binstep", BinStep},
	"relu":      {"relu", ReLU},
	"leakyrelu": {"leakyrelu", LeakyReLU},
	"sigmoid":   {"sigmoid", Sigmoid},
	"tanh":      {"tanh", Tanh},
}

var ActivationSlice = []ActivationFn{
	{"linear", Linear},
	{"binstep", BinStep},
	{"relu", ReLU},
	{"leakyrelu", LeakyReLU},
	{"sigmoid", Sigmoid},
	{"tanh", Tanh},
}

func RegisterActivation(name string, fn func(float64) float64) {
	ActivationMap[name] = ActivationFn{name, fn}
	ActivationSlice = append(ActivationSlice, ActivationFn{name, fn})
}
