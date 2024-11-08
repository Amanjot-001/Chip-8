package cpu

import (
	"chip-8/pkg/memory"
	"log"
)

type CPU struct {
	Memory    *memory.Memory // 4096 bytes of memory
	Registers [16]uint8      // 16 8-bit registers
	I         uint16         // addr register
	PC        uint16         // program counter
	Stack     [16]uint16     // chip8 hax max 16 levels depth for stack
	SP        uint8          // stack pointer
}

func NewCPU(gamePath string) (*CPU, error) {
	cpu := &CPU{
		Memory: memory.NewMemory(),
		PC:     0x200, // Programs start at 0x200
	}
	cpu.Reset()
	err := cpu.LoadGame(gamePath)
	return cpu, err
}

func (cpu *CPU) Reset() {
	cpu.I = 0                      // addr register to 0
	cpu.PC = 0x200                 // starting addr of all games
	for i := range cpu.Registers { // initialize all registers to 0
		cpu.Registers[i] = 0
	}

	cpu.Memory.Clear() // mem clear
}

func (cpu *CPU) LoadGame(gamePath string) error {
	err := cpu.Memory.LoadGame(gamePath, 0x200)
	return err
}

// 2 bytes opcode but 1 byte mem size
// combining PC and PC+1 to get opcode (big-endian)
func (cpu *CPU) GetNextOpcode() uint16 {
	var res uint16 = 0
	res = uint16(cpu.Memory.Read(cpu.PC))      // first byte
	res <<= 8                                  // left shift 8
	res |= uint16(cpu.Memory.Read(cpu.PC + 1)) // second byte
	cpu.PC += 2                                // increment PC

	return res
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
