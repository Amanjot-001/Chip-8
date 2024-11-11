package display

import "fmt"

const (
	Width  = 64
	Height = 32
)

type Display struct {
	Pixels [Height][Width]uint8 // 2D array representing the screen state
}

func NewDisplay() *Display {
	d := &Display{}
	d.Clear()
	return d
}

func (d *Display) Clear() {
	for y := range d.Pixels {
		for x := range d.Pixels[y] {
			d.Pixels[y][x] = 0
		}
	}
}

func (d *Display) DrawSprite(x, y uint8, sprite []uint8) bool {
	collision := false
	for row := 0; row < len(sprite); row++ {
		spriteRow := sprite[row]
		for col := 0; col < 8; col++ {
			pixelX := (x + uint8(col)) % Width
			pixelY := (y + uint8(row)) % Height

			pixelState := (spriteRow >> (7 - col)) & 1

			// both are 1 then resulting pixel will be 0 (means VF = 1)
			if d.Pixels[pixelY][pixelX] == 1 && pixelState == 1 {
				collision = true
			}

			// pixel toggle
			d.Pixels[pixelY][pixelX] ^= pixelState
		}
	}

	return collision
}

func (d *Display) Render() {
	for y := 0; y < Height; y++ {
		for x := 0; x < Width; x++ {
			if d.Pixels[y][x] == 1 {
				fmt.Print("#") // Represents "on" pixel
			} else {
				fmt.Print(" ") // Represents "off" pixel
			}
		}
		fmt.Println("")
	}
}
