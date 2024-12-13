package tui

import (
	"fmt"
	"maps"
	logger "sic_vm"
	"sic_vm/simulator"
	"sic_vm/vm"
	"slices"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

const MEM_TABLE_WIDTH = 16

var (
	app            *tview.Application
	registersTable *tview.Table
	memoryTable    *tview.Table
	startStopBtn   *tview.Button
	stepBtn        *tview.Button
	sim            *simulator.Simulator
)

func Setup(simulator *simulator.Simulator) {
	app = tview.NewApplication()
	sim = simulator

	root := tview.NewFlex()
	root.SetBorder(true).SetTitle("SIX/XE Simulator")

	leftPane := tview.NewFlex().SetDirection(tview.FlexRow)
	leftPane.SetBorder(true)

	controls := createControls()
	leftPane.AddItem(controls, 3, 1, true)
	registersTable = createRegistersTable()
	leftPane.AddItem(registersTable, 0, 1, false)

	rightPane := createMemoryTable()
	mainSplit := tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(leftPane, 0, 4, false).
		AddItem(rightPane, 0, 5, false)

	root.AddItem(mainSplit, 0, 1, false)

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == 'q' {
			app.Stop()
		}
		return event
	})

	simulator.Subscribe(func(vm *vm.VM) {
		// update registers, memory, ...
		app.QueueUpdateDraw(func() {
			updateRegisters(vm.Registers)
			updateMemory(&vm.Memory)
			updateControls()
		})
	})

	if err := app.SetRoot(root, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}

func createControls() *tview.Flex {
	startStopBtn = tview.NewButton("Start").SetSelectedFunc(func() {
		go sim.Toggle()
	})
	stepBtn = tview.NewButton("Step").SetSelectedFunc(func() {
		go sim.Step()
	})
	controls := tview.NewFlex()
	controls.SetBorderPadding(0, 0, 1, 1)
	controls.AddItem(startStopBtn, 20, 1, true)
	controls.AddItem(stepBtn, 20, 1, false)
	controls.AddItem(tview.NewTextView().SetText("Speed: "), 10, 1, false)
	controls.AddItem(tview.NewInputField(), 30, 1, false)
	return controls
}

func createRegistersTable() *tview.Table {
	regTable := tview.NewTable().
		SetBorders(true)
	idx := 0
	for _, regName := range vm.RegisterNames {
		regTable.SetCellSimple(0, idx, regName)
		regTable.GetCell(0, idx).SetAlign(tview.AlignCenter)
		regTable.SetCellSimple(1, idx, "0x000000")
		idx += 1
	}
	return regTable
}

// TODO: use VirtualTable to support entire memory
// https://github.com/rivo/tview/wiki/VirtualTable
func createMemoryTable() *tview.Flex {
	rightPane := tview.NewFlex().SetDirection(tview.FlexRow)
	rightPane.SetBorder(true)
	rightPane.AddItem(tview.NewTextView().SetText("Memory"), 1, 1, false)
	memoryTable = tview.NewTable().SetBorders(true)
	for rowIdx := 0; rowIdx < 100; rowIdx++ {
		memoryTable.SetCellSimple(rowIdx, 0, fmt.Sprintf(" 0x%X ", rowIdx*16))
		memoryTable.GetCell(rowIdx, 0).SetBackgroundColor(tcell.ColorDarkBlue)
	}
	updateMemory(&sim.Vm.Memory)
	rightPane.AddItem(memoryTable, 0, 1, false)
	return rightPane
}

func updateRegisters(registers *vm.Registers) {
	for idx, regIdx := range slices.Collect(maps.Keys(vm.RegisterNames)) {
		if regIdx != 6 {
			formatted := fmt.Sprintf("0x%06X", registers.GetReg(regIdx))
			registersTable.SetCellSimple(1, idx, formatted)
		}
	}
}

func updateControls() {
	logger.Log.Printf("sim: isRunning=%v isDone=%v\n", sim.IsRunning, sim.IsDone)
	if sim.IsRunning {
		startStopBtn.SetLabel("Stop")
		stepBtn.SetLabel("Running")
	} else {
		startStopBtn.SetLabel("Start")
		stepBtn.SetLabel("Step")
	}
	if sim.IsDone {
		startStopBtn.SetLabel("Done")
		stepBtn.SetLabel("Done")
	}
}

func updateMemory(memory *vm.Memory) {
	for row := 0; row < 100; row++ {
		for col := 0; col < MEM_TABLE_WIDTH; col++ {
			addr := row*MEM_TABLE_WIDTH + col
			value, _ := memory.GetByte(int32(addr))
			memoryTable.SetCellSimple(row, col+1, fmt.Sprintf(" 0x%02X ", value))
			memoryTable.GetCell(row, col+1).SetTextColor(tcell.ColorWhite)
		}
	}
	pc := sim.Vm.Registers.PC
	row := int(pc / MEM_TABLE_WIDTH)
	col := int(pc % MEM_TABLE_WIDTH)
	memoryTable.GetCell(row, col+1).SetTextColor(tcell.ColorRed)
}

// TODO: optional, add a window on left to show stdout
