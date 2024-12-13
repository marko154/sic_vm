package simulator

import (
	"io"
	logger "sic_vm"
	"sic_vm/vm"
	"time"
)

type Simulator struct {
	Vm        *vm.VM
	Speed     int
	IsRunning bool
	IsDone    bool
	callbacks []func(*vm.VM)
	quit      chan bool
}

func NewSimulator(reader io.Reader) *Simulator {
	machine := vm.NewVM(reader)
	machine.Load()
	return &Simulator{
		Vm:        machine,
		callbacks: []func(*vm.VM){},
		IsRunning: false,
	}
}

func (s *Simulator) Subscribe(callback func(*vm.VM)) func() {
	s.callbacks = append(s.callbacks, callback)
	return func() {
		// TODO: remove callback from sliceo
		// slices.DeleteFunc(s.callbacks, func(el func(*vm.VM)) bool {
		// return el == callback
		// })
	}
}

func (s *Simulator) Step() {
	done, err := s.Vm.Step()
	s.IsDone = done
	logger.Log.Printf("done=%v\n", done)
	if err != nil {
		panic(err)
	}
	if done {
		s.IsRunning = false
		s.Stop()
	}
	s.runCallbacks()
}

// TODO: check where/if mutexes are needed

func (s *Simulator) Toggle() {
	if s.IsRunning {
		s.Stop()
	} else {
		s.Start()
	}
}

func (s *Simulator) Start() {
	ticker := time.NewTicker(100 * time.Millisecond)
	s.IsRunning = true
	s.runCallbacks()
	s.quit = make(chan bool)
	go func() {
		for {
			select {
			case <-ticker.C:
				s.Step()
			case <-s.quit:
				logger.Log.Printf("quitting\n")
				s.IsRunning = false
				ticker.Stop()
				s.runCallbacks()
				return
			}
		}
	}()
}

func (s *Simulator) runCallbacks() {
	logger.Log.Printf("running callbacks\n")
	for _, callback := range s.callbacks {
		callback(s.Vm)
	}
}

func (s *Simulator) Stop() {
	logger.Log.Printf("Stop()\n")
	close(s.quit)
}

func (s *Simulator) SetSpeed(speed int) {
	s.Speed = speed
}
