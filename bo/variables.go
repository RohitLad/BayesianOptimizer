package bo

import (
	"errors"
	"math/rand"
)

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
	xList     [][]float64
	yList     []float64
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

func (vm VariableManager) VerifyBounds(samples []float64) (bool, error) {

	for i := 0; i < len(samples); i++ {
		vType := vm.variables[vm.names[i]]
		if !(samples[i] <= vType.Max && samples[i] >= vType.Min) {
			errStr := "Bounds not satisfied for sample vector index" + string(i)
			return false, errors.New(errStr)
		}
	}
	return false, nil
}

func (vm *VariableManager) SaveMapping(samples []float64, feval float64) {
	vm.xList = append(vm.xList, samples)
	vm.yList = append(vm.yList, feval)
}

func (vm VariableManager) PeepMapping(index int) (x []float64, y float64, err error) {
	if len(vm.xList)-1 < index {
		return []float64{0.}, 0., errors.New("Improper index")
	}
	return vm.xList[index], vm.yList[index], nil
}
