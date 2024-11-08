package cpu

import (
	"log"
	"os"
)

type CPU struct {
	Memory    [4096]uint8 // 4096 bytes of memory
	Registers [16]uint8   // 16 8-bit registers
	I         uint16      // addr register
	PC        uint16      // program counter
	Stack     []uint16    // stack
}

func Reset(cpu *CPU) {
	cpu.I = 0                      // addr register to 0
	cpu.PC = 0x200                 // starting addr of all games
	for i := range cpu.Registers { // initialize all registers to 0
		cpu.Registers[i] = 0
	}
	for i := range cpu.Memory { // mem cleared
		cpu.Memory[i] = 0
	}
	cpu.Stack = make([]uint16, 16) // reinit stack with 0

	gameData, err := os.ReadFile("")
	if err != nil {
		log.FatalF("Failed to load game: %v", err)
	}

	if len(gameData) > len(cpu.Memory)-0x200 {
		log.Fatalf("Game too large to fit in memory.")
	}

	copy(cpu.Memory[ox200:], gameData)
}
