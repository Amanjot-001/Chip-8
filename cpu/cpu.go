package cpu

type CPU struct {
	Memory    [4096]uint8
	Registers [16]uint8
	I         uint16
	PC        uint16
	Stack     []uint16
}
