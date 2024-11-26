package main

import (
	"chip-8/pkg/cpu"
	"chip-8/pkg/debugger"
	"log"

	"time"

	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	gamePath := "./tests/5-quirks.ch8"

	chip8, err := cpu.NewCPU(gamePath)
	if err != nil {
		log.Fatalf("Failed to initialize CPU: %v\n", err)
	}

	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		log.Fatalf("Failed to initialize SDL: %v\n", err)
	}
	defer sdl.Quit()

	/* _________________ debugger init _________________ */
	debugger := debugger.NewDebugger(chip8)
	// debugger.PrintMemory(chip8.PC, chip8.PC+34)

	/* _________________ Main emulator loop _________________ */

	// frames per second supported
	fps := 60

	// time for each frame to elapse => 16.67 ms
	interval := time.Second / time.Duration(fps)

	// opcodes to execute per second
	numOfOpcodes := 600

	// opcodes to execute per frame
	numFrame := numOfOpcodes / fps

	quit := false
	stepMode := false

	for !quit {
		startTime := time.Now()
		// log.Println("start time", startTime)

		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			// log.Printf("Event captured: %T\n", event)
			switch e := event.(type) {
			case *sdl.QuitEvent:
				quit = true
				log.Printf("quit")
			case *sdl.KeyboardEvent:
				log.Printf("Key event received: %+v\n", e)
				chip8.Keys.HandleKeyPress(e)
			}
		}

		// executionLoopStart := time.Now()
		for i := 0; i < numFrame; i++ {
			if stepMode {
				debugger.WaitForKeyPress(&quit)
			}

			opcode := chip8.GetNextOpcode()
			chip8.DecodeAndExecute(opcode)

			// instStart := time.Now()
			if chip8.DrawFlag {
				chip8.Display.Render()
				chip8.DrawFlag = false // Reset after rendering
			}
			// log.Println("instruction since", time.Since(instStart))

			// log.Printf("Executing Opcode: 0x%X\n", opcode)

			if stepMode {
				log.Printf("Executing Opcode: 0x%X\n", opcode)
				// debugger.PrintState()
			}

		}
		// log.Println("loop since", time.Since(executionLoopStart))

		chip8.DecreaseTimers()

		elapsed := time.Since(startTime)
		// log.Println("elapsed since", elapsed)

		if elapsed < interval {
			// delaystarttime := time.Now()
			// log.Println("delaystarttime", delaystarttime)
			time.Sleep(interval - elapsed)
			// log.Println("delay since", time.Since(delaystarttime))
		}

		// log.Printf("Frame elapsed: %v, Target interval: %v\n", elapsed, interval)
	}
}
