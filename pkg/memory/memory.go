package memory

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
