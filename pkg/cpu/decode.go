package cpu

import (
	"chip-8/pkg/input"
	"log"
	"math/rand"

	"github.com/veandco/go-sdl2/sdl"
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

			or := cpu.Registers[regX] | cpu.Registers[regY]
			cpu.Registers[regX] = or
			cpu.Registers[0xF] = 0
		case 0x0002:
			regX := second
			regX >>= 8
			regY := third
			regY >>= 4

			and := cpu.Registers[regX] & cpu.Registers[regY]
			cpu.Registers[regX] = and
			cpu.Registers[0xF] = 0
		case 0x0003:
			regX := second
			regX >>= 8
			regY := third
			regY >>= 4

			xor := cpu.Registers[regX] ^ cpu.Registers[regY]
			cpu.Registers[regX] = xor
			cpu.Registers[0xF] = 0
		case 0x0004:
			regX := second
			regX >>= 8
			regY := third
			regY >>= 4

			add := uint16(cpu.Registers[regX]) + uint16(cpu.Registers[regY])
			cpu.Registers[regX] = uint8(add & 0xFF)
			if add > 0xFF {
				cpu.Registers[0xF] = 1
			} else {
				cpu.Registers[0xF] = 0
			}
		case 0x0005:
			regX := second
			regX >>= 8
			regY := third
			regY >>= 4

			xval := cpu.Registers[regX]
			yval := cpu.Registers[regY]

			cpu.Registers[regX] -= yval
			cpu.Registers[0xF] = 1
			if yval > xval {
				cpu.Registers[0xF] = 0
			}
		case 0x0006:
			regX := second
			regX >>= 8
			regY := third
			regY >>= 4

			lsb := cpu.Registers[regY] & 0x0001
			cpu.Registers[regX] = (cpu.Registers[regY] >> 1)
			cpu.Registers[0xF] = lsb
		case 0x0007:
			regX := second
			regX >>= 8
			regY := third
			regY >>= 4

			xval := cpu.Registers[regX]
			yval := cpu.Registers[regY]

			cpu.Registers[regX] = yval - xval
			cpu.Registers[0xF] = 1
			if xval > yval {
				cpu.Registers[0xF] = 0
			}
		case 0x000E:
			regX := second
			regX >>= 8
			regY := third
			regY >>= 4

			msb := (cpu.Registers[regY] & 0b10000000) >> 7
			cpu.Registers[regX] = (cpu.Registers[regY] << 1)
			cpu.Registers[0xF] = msb
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

		sprite := make([]uint8, height)
		for row := 0; row < int(height); row++ {
			sprite[row] = cpu.Memory.Read(cpu.I + uint16(row))
		}

		collision, updated := cpu.Display.DrawSprite(xval, yval, sprite)
		cpu.Registers[0xF] = 0x00
		if collision {
			cpu.Registers[0xF] = 0x01
		}
		cpu.DrawFlag = updated
	case 0xE000:
		switch last2 {
		case 0x009E:
			regX := second
			regX >>= 8
			key := cpu.Registers[regX] & 0x0F

			if cpu.Keys.IsKeyPressed(key) {
				cpu.PC += 2
			}
		case 0x00A1:
			regX := second
			regX >>= 8
			key := cpu.Registers[regX] & 0x0F

			if !cpu.Keys.IsKeyPressed(key) {
				cpu.PC += 2
			}
		}
	case 0xF000:
		switch last2 {
		case 0x0007:
			regX := second
			regX >>= 8

			cpu.Registers[regX] = cpu.DelayTimer
		case 0x000A:
			regX := second
			regX >>= 8
			keyPressed := false

			for !keyPressed {
				event := sdl.WaitEvent()
				switch e := event.(type) {
				case *sdl.QuitEvent:
					// quit := true
					return
				case *sdl.KeyboardEvent:
					if e.State == sdl.PRESSED {
						key := input.KeyMap[e.Keysym.Sym]
						if key != 0 {
							cpu.Registers[regX] = uint8(key)
							keyPressed = true
						}
					}
				}
			}
		case 0x0015:
			regX := second
			regX >>= 8

			cpu.DelayTimer = cpu.Registers[regX]
		case 0x0018:
			regX := second
			regX >>= 8

			cpu.SoundTimer = cpu.Registers[regX]
		case 0x001E:
			regX := second
			regX >>= 8

			cpu.I += uint16(cpu.Registers[regX])
		case 0x0029:
			regX := second
			regX >>= 8

			// register is 8 bits and can store 2 chars
			// so we use char in last nibble
			character := cpu.Registers[regX] & 0x0F

			cpu.I = uint16(0x50 + character)
		case 0x0033: // Binary-coded
			regX := second
			regX >>= 8

			value := cpu.Registers[regX]

			hundreds := value / 100
			tens := (value / 10) % 10
			units := value % 10

			cpu.Memory.Write(cpu.I, hundreds)
			cpu.Memory.Write(cpu.I+1, tens)
			cpu.Memory.Write(cpu.I+2, units)
		case 0x0055:
			regX := second
			regX >>= 8

			for i := 0; i <= int(regX); i++ {
				cpu.Memory.Write(cpu.I+uint16(i), cpu.Registers[i])
			}

			cpu.I += regX + 1 // only for old games
		case 0x0065:
			regX := second
			regX >>= 8

			for i := 0; i <= int(regX); i++ {
				cpu.Registers[uint8(i)] = cpu.Memory.Read(cpu.I + uint16(i))
			}

			cpu.I += regX + 1 // only for old games
		}
	default:
		log.Fatalf("Unknown opcode: 0x%X\n", opcode)
	}
}
