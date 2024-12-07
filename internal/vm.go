package vm

import (
	"bufio"
	"os"
)

type VM struct {
	Memory       Memory
	Registers    *Registers
	Devices      map[int]Device // device number -> device
	Loader       *Loader
	ProgramName  string
	CodeAddress  int
	StartAddress int
	CodeLength   int
}

func NewVM(reader *bufio.Reader) *VM {
	vm := &VM{
		Loader:    NewLoader(reader),
		Devices:   make(map[int]Device),
		Registers: NewRegisters(),
	}
	vm.SetDevice(0, NewInputDevice(os.Stdin))
	vm.SetDevice(1, NewInputDevice(os.Stdout))
	vm.SetDevice(2, NewInputDevice(os.Stderr))
	return vm
}

func (vm *VM) Load() error {
	return vm.Loader.Load(vm)
}

// TODO: bounds check?
func (vm *VM) GetDevice(num int) Device {
	device := vm.Devices[num]
	return device
}

func (vm *VM) SetDevice(num int, device Device) {
	vm.Devices[num] = device
}

// !!! warning: sign extension when using operands !!!

func (vm *VM) Run() error {
	vm.Registers.PC = vm.StartAddress

	for {
		// TODO: implement HALT
		// if detect HALT: break loop
		opcode, err := vm.fetch()
		if err != nil {
			return err
		}
		if err := vm.execute(opcode); err != nil {
			return err
		}
	}
	return nil
}

func (vm *VM) fetch() (byte, error) {
	value, err := vm.Memory.GetByte(vm.Registers.PC)
	if err != nil {
		return 0, err
	}
	vm.Registers.PC++
	return value, nil
}

func (vm *VM) execute(opcode byte) error {

	return nil
}

func (vm *VM) tryExecuteType1(opcode byte) (bool, error) {
	return false, nil
}
func (vm *VM) tryExecuteType2(opcode byte) error {
	return nil
}
