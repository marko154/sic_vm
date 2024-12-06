package vm

import "bufio"

type VM struct {
	Memory       Memory
	Registers    Registers
	Devices      map[int]Device // device number -> device
	Loader       *Loader
	ProgramName  string
	CodeAddress  int
	StartAddress int
	CodeLength   int
}

func NewVM(reader *bufio.Reader) *VM {
	return &VM{
		Loader: NewLoader(reader),
	}
}

func (vm *VM) Load() error {
	return vm.Loader.Load(vm)
}

// !!! warning: sign extension when using operands !!!

func (vm *VM) Run() {
}
