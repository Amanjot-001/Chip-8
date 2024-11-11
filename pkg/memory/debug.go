package memory

import "fmt"

func (m *Memory) Extract() {
	var instructions []uint16
	for i := 0x200; i < 0x1000; i += 2 {
		byte1 := m.Read(uint16(i))
		byte2 := m.Read(uint16(i + 1))
		opcode := (uint16(byte1) << 8) + uint16(byte2)
		instructions = append(instructions, opcode)
	}

	for _, opcode := range instructions {
		if opcode == 0 {
			break
		}
		fmt.Printf("0x%04X\n", opcode)
	}
}
