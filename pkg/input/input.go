package input

type Input struct {
	keys [16]bool // 0x0 - 0xF
}

func NewInput() *Input {
	i := &Input{}
	return i
}