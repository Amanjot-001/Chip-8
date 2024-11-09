package cpu

import "log"

func (cpu *CPU) DecodeAndExecute(opcode uint16) {
	first := opcode & 0xF000
	second := opcode & 0x0F00
	third := opcode & 0x00F0
	fourth := opcode & 0x000F
	last3 := opcode & 0x0FFF
	last2 := opcode & 0x00FF

	switch first {
	case 0x0000:
		switch fourth {
		case 0x000E:

		case 0x0000:

		default:
			log.Printf("Unknown opcode: 0x%X\n", opcode)
		}
	case 0x1000:
		cpu.PC = last3
	case 0x2000:
		cpu.PushToStack(cpu.PC)
		cpu.PC = last3
	case 0x5000:
		regX := second
		regX >>= 8
		regY := third
		regY >>= 4

		if cpu.Registers[regX] == cpu.Registers[regY] {
			cpu.PC += 2 // skip next instruction
		}
	case 0x6000:
		register := (second) >> 8
		value := uint8(last2)
		cpu.Registers[register] = value
	case 0x7000:
		register := (second) >> 8
		addVal := uint8(last2)
		cpu.Registers[register] += addVal
	case 0x8000:
		switch fourth {
		case 0x0005:
			cpu.Registers[0xF] = 1
			regX := second
			regX >>= 8
			regY := third
			regY >>= 4

			xval := cpu.Registers[regX]
			yval := cpu.Registers[regY]
			
			if yval > xval {
				cpu.Registers[0xF] = 0
			}
			cpu.Registers[regX] -= yval
		}
	case 0xA000:
		cpu.I = last3
	default:
		log.Printf("Unknown opcode: 0x%X\n", opcode)
	}
}
