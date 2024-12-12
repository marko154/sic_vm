package simulator

import (
	"io"
	"sic_vm/vm"
	"time"
)

type Simulator struct {
	vm        *vm.VM
	speed     int
	callbacks []func(*vm.VM)
	quit      chan struct{}
	isRunning bool
}

func NewSimulator(reader io.Reader) *Simulator {
	machine := vm.NewVM(reader)
	machine.Load()
	return &Simulator{
		vm:        machine,
		callbacks: []func(*vm.VM){},
		quit:      make(chan struct{}),
	}
}

func (s *Simulator) Subscribe(callback func(*vm.VM)) {
	s.callbacks = append(s.callbacks, callback)
}

func (s *Simulator) Step() {
	_, err := s.vm.Step()
	s.isRunning = true
	if err != nil {
		panic(err)
	}
	for _, callback := range s.callbacks {
		callback(s.vm)
	}
}

func (s *Simulator) Start() {
	ticker := time.NewTicker(1 * time.Second)
	go func() {
		for {
			select {
			case <-ticker.C:
				s.Step()
			case <-s.quit:
				return
			}
		}
	}()
}

func (s *Simulator) Stop() {
	s.quit <- struct{}{}
}

func (s *Simulator) SetSpeed(speed int) {
	s.speed = speed
}
