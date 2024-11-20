package debugger

import (
	"chip-8/pkg/cpu"
	"fmt"
)

type Debugger struct {
	CPU *cpu.CPU
}

func NewDebugger(cpu *cpu.CPU) *Debugger {
	return &Debugger{CPU: cpu}
}

func (d *Debugger) PrintState() {
	fmt.Println("---- CPU State ----")
	fmt.Printf("PC: 0x%04X  I: 0x%04X  SP: 0x%02X\n", d.CPU.PC, d.CPU.I, d.CPU.SP)
	fmt.Println("Registers:")
	for i, reg := range d.CPU.Registers {
		fmt.Printf("V%X: 0x%02X  ", i, reg)
		if (i+1)%4 == 0 {
			fmt.Println()
		}
	}
	fmt.Printf("\nDelayTimer: %d  SoundTimer: %d\n", d.CPU.DelayTimer, d.CPU.SoundTimer)
	fmt.Println("\nStack:")
	for i := 0; i < int(d.CPU.SP); i++ {
		fmt.Printf("0x%04X ", d.CPU.Stack[i])
	}
	fmt.Println("")
}

func (d *Debugger) PrintMemory(start, end uint16) {
	fmt.Printf("---- Memory Dump (0x%04X to 0x%04X) ----\n", start, end)
	for addr := start; addr <= end; addr++ {
		fmt.Printf("0x%04X: 0x%02X\n", addr, d.CPU.Memory.Read(addr))
	}
}
