package input

import (
	"log"

	"github.com/veandco/go-sdl2/sdl"
)

type Input struct {
	Keys [16]uint8 // 0x0 - 0xF
}

func NewInput() *Input {
	in := &Input{}
	return in
}

func (in *Input) Clear() {
	for i := 0x0; i <= 0xF; i++ {
		in.Keys[i] = 0
	}
}

var KeyMap = map[sdl.Keycode]uint8{
	sdl.K_1: 0x1, sdl.K_2: 0x2, sdl.K_3: 0x3, sdl.K_4: 0xC,
	sdl.K_q: 0x4, sdl.K_w: 0x5, sdl.K_e: 0x6, sdl.K_r: 0xD,
	sdl.K_a: 0x7, sdl.K_s: 0x8, sdl.K_d: 0x9, sdl.K_f: 0xE,
	sdl.K_z: 0xA, sdl.K_x: 0x0, sdl.K_c: 0xB, sdl.K_v: 0xF,
}

func (in *Input) SetKey(key uint8, pressed uint8) {
	if key < 16 {
		in.Keys[key] = pressed
	}
}

func (in *Input) IsKeyPressed(key uint8) bool {
	if key < 16 {
		return in.Keys[key] == 1
	}

	return false
}

func (in *Input) HandleKeyPress(event sdl.Event) {
	switch t := event.(type) {
	case *sdl.KeyboardEvent:
		key, ok := KeyMap[t.Keysym.Sym] // val, bool
		if ok {
			if t.State == sdl.PRESSED {
				log.Printf("Key pressed: %v -> %v\n", t.Keysym.Sym, key)
				in.SetKey(key, 1)
			} else if t.State == sdl.RELEASED {
				log.Printf("Key released: %v -> %v\n", t.Keysym.Sym, key)
				in.SetKey(key, 0)
			} else {
				log.Printf("Unhandled key: %v\n", t.Keysym.Sym)
			}
		}
	}
}
