package display

const (
	Width  = 64
	Height = 32
)

type Display struct {
	Pixels [Height][Width]uint8 // 2D array representing the screen state
}

func NewDisplay() *Display {
	return &Display{}
}

func (d *Display) Clear() {
	for y := range d.Pixels {
		for x := range d.Pixels[y] {
			d.Pixels[y][x] = 0
		}
	}
}
