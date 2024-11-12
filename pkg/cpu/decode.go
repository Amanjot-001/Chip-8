package cpu

import (
	"log"
	"math/rand"
)

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

func (cpu *CPU) DecodeAndExecute(opcode uint16) {
	first := opcode & 0xF000
	second := opcode & 0x0F00
	third := opcode & 0x00F0
	fourth := opcode & 0x000F
	last3 := opcode & 0x0FFF
	last2 := opcode & 0x00FF

	switch first {
	case 0x0000:
		switch last3 {
		case 0x00E0:
			cpu.Display.Clear()
		case 0x00EE:
			cpu.PC = cpu.PopFromStack()
		}
	case 0x1000:
		cpu.PC = last3
	case 0x2000:
		cpu.PushToStack(cpu.PC)
		cpu.PC = last3
	case 0x3000:
		regX := second
		regX >>= 8
		if cpu.Registers[regX] == uint8(last2) {
			cpu.PC += 2
		}
	case 0x4000:
		regX := second
		regX >>= 8
		if cpu.Registers[regX] != uint8(last2) {
			cpu.PC += 2
		}
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
		case 0x0000:
			regX := second
			regX >>= 8
			regY := third
			regY >>= 4

			cpu.Registers[regX] = cpu.Registers[regY]
		case 0x0001:
			regX := second
			regX >>= 8
			regY := third
			regY >>= 4

			cpu.Registers[regX] |= cpu.Registers[regY]
		case 0x0002:
			regX := second
			regX >>= 8
			regY := third
			regY >>= 4

			cpu.Registers[regX] &= cpu.Registers[regY]
		case 0x0003:
			regX := second
			regX >>= 8
			regY := third
			regY >>= 4

			cpu.Registers[regX] ^= cpu.Registers[regY]
		case 0x0004:
			cpu.Registers[0xF] = 0
			regX := second
			regX >>= 8
			regY := third
			regY >>= 4

			xval := cpu.Registers[regX]
			yval := cpu.Registers[regY]

			if uint16(xval+yval) > 255 {
				cpu.Registers[0xF] = 1
			}
			cpu.Registers[regX] += yval
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
		case 0x0006: // according to original chip8 (wikipedia)
			regX := second
			regX >>= 8

			cpu.Registers[0xF] = cpu.Registers[regX] & 0x0F
			cpu.Registers[regX] >>= 1
		case 0x0007:
			cpu.Registers[0xF] = 1
			regX := second
			regX >>= 8
			regY := third
			regY >>= 4

			xval := cpu.Registers[regX]
			yval := cpu.Registers[regY]

			if xval > yval {
				cpu.Registers[0xF] = 0
			}
			cpu.Registers[regX] = yval - xval
		case 0x000E: // according to original chip8 (wikipedia)
			regX := second
			regX >>= 8

			cpu.Registers[0xF] = cpu.Registers[regX] & 0xF0
			cpu.Registers[regX] <<= 1
		}
	case 0x9000:
		regX := second
		regX >>= 8
		regY := third
		regY >>= 4

		if cpu.Registers[regX] != cpu.Registers[regY] {
			cpu.PC += 2 // skip next instruction
		}
	case 0xA000:
		cpu.I = last3
	case 0xB000:
		cpu.PC = uint16(cpu.Registers[0x00]) + last3
	case 0xC000:
		regX := second
		regX >>= 8
		randomNumber := rand.Intn(256)
		cpu.Registers[regX] = uint8(randomNumber) & uint8(last2)
	case 0xD000:
		regX := second
		regX >>= 8
		regY := third
		regY >>= 4
		height := fourth

		xval := cpu.Registers[regX]
		yval := cpu.Registers[regY]

		cpu.Registers[0xF] = 0

		var sprite []uint8
		for row := 0; row < int(height); row++ {
			sprite = append(sprite, uint8(cpu.I+uint16(row)))
		}

		if cpu.Display.DrawSprite(xval, yval, sprite) {
			cpu.Registers[0xF] = 1
		}
	case 0xE000:
		switch last2 {
		case 0x009E:
			regX := second
			regX >>= 8
			key := cpu.Registers[regX]

			if cpu.keys.IsKeyPressed(key) {
				cpu.PC += 2
			}
		case 0x00A1:
			regX := second
			regX >>= 8
			key := cpu.Registers[regX]

			if !cpu.keys.IsKeyPressed(key) {
				cpu.PC += 2
			}
		}
	case 0xF000:
		switch last2 {
		case 0x0055:
			regX := second
			regX >>= 8

			for i := 0; i <= int(regX); i++ {
				cpu.Memory.Write(cpu.I+uint16(i), cpu.Registers[i])
			}

			cpu.I += regX + 1 // but wiki says not to change
		case 0x0065:
			regX := second
			regX >>= 8

			for i := 0; i <= int(regX); i++ {
				cpu.Registers[uint8(regX)+uint8(i)] = cpu.Memory.Read(cpu.I + uint16(i))
			}

			cpu.I += regX + 1 // but wiki says not to change
		}
	default:
		log.Fatalf("Unknown opcode: 0x%X\n", opcode)
	}
}
