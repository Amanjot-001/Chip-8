package timers

import (
	"chip-8/pkg/cpu"
	"context"
	"time"

	"github.com/veandco/go-sdl2/mix"
)

// StartTimers runs a goroutine to decrement DelayTimer and SoundTimer at 60 Hz.
func StartTimers(ctx context.Context, beepSound *mix.Chunk, cpu *cpu.CPU) {
	ticker := time.NewTicker(time.Second / 60) // 60 Hz
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return // Exit the goroutine when context is canceled
		case <-ticker.C:
			if cpu.DelayTimer > 0 {
				cpu.DelayTimer--
			}
			if cpu.SoundTimer > 0 {
				cpu.SoundTimer--
				beepSound.Play(-1, 0) // 0 = loop count (0 means play once)
			} else {
				mix.HaltMusic()
			}
		}
	}
}

// func (cpu *CPU) DecreaseTimers() {
// 	if cpu.DelayTimer > 0 {
// 		cpu.DelayTimer--
// 	}

// 	if cpu.SoundTimer > 0 {
// 		cpu.SoundTimer--
// 	}
// }
