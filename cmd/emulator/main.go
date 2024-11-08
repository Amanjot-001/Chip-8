package main

import (
	"chip-8/pkg/cpu"
	"fmt"
	"log"
)

func main() {
	gamePath := "./games/pong.rom"

	chip8, err := cpu.NewCPU(gamePath)
	if err != nil {
		log.Fatalf("Failed to initialize CPU: %v\n", err)
	}

	// Basic tests to confirm the initial state
	fmt.Printf("Program Counter: 0x%X\n", chip8.PC) // Should print: Program Counter: 0x200
	fmt.Printf("Index Register: 0x%X\n", chip8.I)   // Should print: Index Register: 0x0
	fmt.Println("Registers:", chip8.Registers)      // Should print all zeros
	fmt.Println("Stack:", chip8.Stack)              // Should print 16 zeros

	// Basic test for memory loading (assuming game data exists)
	for i := 0; i < 10; i++ { // Print first 10 bytes after 0x200
		fmt.Printf("Memory[0x%X]: 0x%X\n", 0x200+i, chip8.Memory.Read(uint16(0x200+i)))
	}
}
