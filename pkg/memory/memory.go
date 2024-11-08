package memory

import "os"

type Memory struct {
	data [4096]uint8
}

func NewMemory() *Memory {
	return &Memory{}
}

func (m *Memory) Clear() {
	for i := range m.data {
		m.data[i] = 0
	}
}

func (m *Memory) Read(address uint16) uint8 {
	if address < 4096 {
		return m.data[address]
	}
	return 0
}

func (m *Memory) Write(address uint16, value uint8) {
	if address < 4096 {
		m.data[address] = value
	}
}

func (m *Memory) LoadGame(gamePath string, startAddress uint16) error {
	gameData, err := os.ReadFile(gamePath)
	if err != nil {
		return err
	}

	if len(gameData) > len(m.data)-int(startAddress) {
		return err
	}

	copy(m.data[startAddress:], gameData)
	return nil
}
