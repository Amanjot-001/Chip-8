package display

import (
	"github.com/veandco/go-sdl2/sdl"
)

const (
	Width       = 64
	Height      = 32
	ScaleFactor = 20
)

type Display struct {
	window   *sdl.Window
	renderer *sdl.Renderer
	Pixels   [Height][Width]uint8 // 2D array representing the screen state
}

func NewDisplay() (*Display, error) {
	window, err := sdl.CreateWindow("CHIP-8 Emulator", sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED, Width*ScaleFactor, Height*ScaleFactor, sdl.WINDOW_SHOWN)
	if err != nil {
		return nil, err
	}

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		window.Destroy()
		return nil, err
	}

	d := &Display{
		window:   window,
		renderer: renderer,
	}
	d.Clear()
	return d, nil
}

func (d *Display) Clear() {
	for y := range d.Pixels {
		for x := range d.Pixels[y] {
			d.Pixels[y][x] = 0
		}
	}
}

func (d *Display) DrawSprite(x, y uint8, sprite []uint8) (bool, bool) {
	collision := false
	updated := false
	for row := 0; row < len(sprite); row++ {
		spriteRow := sprite[row]
		if spriteRow == 0 { // No pixels to draw in this row
			continue
		}
		for col := 0; col < 8; col++ {
			pixelX := (x + uint8(col)) % Width
			pixelY := (y + uint8(row)) % Height

			pixelState := (spriteRow >> (7 - col)) & 1

			// both are 1 then resulting pixel will be 0 (means VF = 1)
			if d.Pixels[pixelY][pixelX] == 1 && pixelState == 1 {
				collision = true
			}

			if pixelState == 1 {
				// pixel toggle
				d.Pixels[pixelY][pixelX] ^= pixelState
				updated = true
			}
		}
	}

	return collision, updated
}

// func (d *Display) Render() {
// 	for y := 0; y < Height; y++ {
// 		for x := 0; x < Width; x++ {
// 			if d.Pixels[y][x] == 1 {
// 				fmt.Print("#") // Represents "on" pixel
// 			} else {
// 				fmt.Print(" ") // Represents "off" pixel
// 			}
// 		}
// 		fmt.Println("")
// 	}
// }

func (d *Display) Render() {
	d.renderer.SetDrawColor(0, 0, 0, 255) // Set draw color to black (clears the screen)
	d.renderer.Clear()

	for y := 0; y < Height; y++ {
		for x := 0; x < Width; x++ {
			if d.Pixels[y][x] == 1 {
				d.renderer.SetDrawColor(255, 255, 255, 255) // White color for "on" pixel
			} else {
				d.renderer.SetDrawColor(0, 0, 0, 255) // Black color for "off" pixel
			}

			// Draw a rectangle representing a pixel
			rect := sdl.Rect{
				X: int32(x * ScaleFactor),
				Y: int32(y * ScaleFactor),
				W: ScaleFactor,
				H: ScaleFactor,
			}
			d.renderer.FillRect(&rect)
		}
	}

	// Present the updated rendering
	d.renderer.Present()
}
