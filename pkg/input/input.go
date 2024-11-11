package input

type Input struct {
	keys [16]uint8 // 0x0 - 0xF
}

func NewInput() *Input {
	in := &Input{}
	return in
}

func (in *Input) Clear() {
	for i := 0x0; i<0xF; i++ {
		in.keys[i] = 0
	}
}