package cpu

import (
	"chip-8/pkg/display"
	"chip-8/pkg/input"
	"chip-8/pkg/memory"
	"log"
)

type CPU struct {
	Memory     *memory.Memory   // 4096 bytes of memory
	Registers  [16]uint8        // 16 8-bit registers
	I          uint16           // addr register
	PC         uint16           // program counter
	Stack      [16]uint16       // chip8 hax max 16 levels depth for stack
	SP         uint8            // stack pointer
	Display    *display.Display // 64x32 display
	Keys       *input.Input     // 16 keys
	DelayTimer uint8
	SoundTimer uint8
}

func NewCPU(gamePath string) (*CPU, error) {
	cpu := &CPU{
		Memory:     memory.NewMemory(),
		PC:         0x200, // Programs start at 0x200
		Display:    display.NewDisplay(),
		Keys:       input.NewInput(),
		DelayTimer: 0,
		SoundTimer: 0,
	}
	cpu.Reset()
	cpu.Memory.LoadFontset()
	err := cpu.LoadGame(gamePath)
	return cpu, err
}

func (cpu *CPU) Reset() {
	cpu.I = 0                      // addr register to 0
	cpu.PC = 0x200                 // starting addr of all games
	for i := range cpu.Registers { // initialize all registers to 0
		cpu.Registers[i] = 0
	}

	cpu.Memory.Clear()  // mem clear
	cpu.Display.Clear() // clear display
	cpu.Keys.Clear()    // clear input
}

func (cpu *CPU) LoadGame(gamePath string) error {
	err := cpu.Memory.LoadGame(gamePath, 0x200)
	return err
}

func (cpu *CPU) PushToStack(value uint16) {
	if cpu.SP >= 16 {
		log.Fatalf("Stack Overflow")
	}

	cpu.Stack[cpu.SP] = value
	cpu.SP++
}

func (cpu *CPU) PopFromStack() uint16 {
	if cpu.SP < 0 {
		log.Fatalf("Stack Underflow")
	}

	cpu.SP--
	return cpu.Stack[cpu.SP]
}

func (cpu *CPU) DecreaseTimers() {
	if cpu.DelayTimer > 0 {
		cpu.DelayTimer--
	}

	if cpu.SoundTimer > 0 {
		cpu.SoundTimer--
		// Optional: Add sound handling logic here if your emulator supports sound
		// Example: Beep or play sound if the sound timer reaches zero
	}
}
