package bo

import "math/rand"

type Variable struct {
	Min     float64
	Max     float64
	Integer bool
	Name    string
}

func (v Variable) Sample() float64 {
	return rand.Float64()*(v.Max-v.Min) + v.Min
}

type VariableManager struct {
	variables map[string]Variable
	names     []string
}

func NewVariableManager() *VariableManager {
	vm := VariableManager{variables: make(map[string]Variable)}
	return &vm
}

func (vm *VariableManager) InsertNewVariable(Name string, Min, Max float64, IsInt bool) {
	if _, ok := vm.variables[Name]; !ok {
		vm.names = append(vm.names, Name)
	}
	vm.variables[Name] = Variable{Min: Min, Max: Max, Name: Name, Integer: IsInt}
}

func (vm *VariableManager) RamdomSample() []float64 {
	dims := len(vm.variables)
	var sample []float64
	for i := 0; i < dims; i++ {
		name := vm.names[i]
		sample = append(sample, vm.variables[name].Sample())
	}
	return sample
}
