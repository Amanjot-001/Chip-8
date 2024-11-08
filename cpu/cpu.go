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

func Reset(cpu *CPU, gamePath string) {
	cpu.I = 0                      // addr register to 0
	cpu.PC = 0x200                 // starting addr of all games
	for i := range cpu.Registers { // initialize all registers to 0
		cpu.Registers[i] = 0
	}
	for i := range cpu.Memory { // mem cleared
		cpu.Memory[i] = 0
	}
	cpu.Stack = make([]uint16, 16) // reinit stack with 0

	gameData, err := os.ReadFile(gamePath) // game read
	if err != nil {
		log.Fatalf("Failed to load game: %v", err)
	}

	if len(gameData) > len(cpu.Memory)-0x200 { // not fit in memory
		log.Fatalf("Game too large to fit in memory.")
	}

	copy(cpu.Memory[0x200:], gameData)
}

// 2 bytes opcode but 1 byte mem size
// combining PC and PC+1 to get opcode (big-endian)
func (cpu *CPU) GetNextOpcode() uint16 {
	var res uint16 = 0
	res = uint16(cpu.Memory[cpu.PC])    // first byte
	res <<= 8                           // left shift 8
	res |= uint16(cpu.Memory[cpu.PC+1]) // second byte
	cpu.PC += 2                         // increment PC

	return res
}
