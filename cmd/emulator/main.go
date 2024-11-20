package main

import (
	"chip-8/pkg/cpu"
	"fmt"
	"log"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	gamePath := "./games/IBM-Logo.ch8"

	chip8, err := cpu.NewCPU(gamePath)
	if err != nil {
		log.Fatalf("Failed to initialize CPU: %v\n", err)
	}

	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		log.Fatalf("Failed to initialize SDL: %v\n", err)
	}
	defer sdl.Quit()

	// Main emulator loop
	fps := 60
	interval := time.Second / time.Duration(fps)
	fmt.Println(interval)
	numOfOpcodes := 600
	numFrame := numOfOpcodes / fps
	fmt.Println(numFrame)
	
	quit := false
	for !quit {
		startTime := time.Now()

		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch e := event.(type) {
			case *sdl.QuitEvent:
				quit = true
			case *sdl.KeyboardEvent:
				chip8.Keys.HandleKeyPress(e)
			}
		}

		for i := 0; i < numFrame; i++ {
			opcode := chip8.GetNextOpcode()
			chip8.DecodeAndExecute(opcode)
		}

		chip8.DecreaseTimers()
		chip8.Display.Render()

		elapsed := time.Since(startTime)
		if elapsed < interval {
			sdl.Delay(uint32(interval - elapsed))
		}
	}
}
