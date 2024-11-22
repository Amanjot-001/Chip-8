# CHIP-8 Emulator in Go

Welcome to the CHIP-8 Emulator! This project is an implementation of the CHIP-8 virtual machine written in Go. It aims to accurately emulate the CHIP-8 architecture, supporting its instruction set, input handling, graphics, and timers. This emulator is designed with modularity and performance in mind, featuring SDL for graphics and input, and includes a step mode for debugging.

## 🚀 Features

1. CHIP-8 Instruction Set Support
   - Full support for the CHIP-8 opcode set.
   - Handles arithmetic, logical, control, and memory instructions.
   - Implements sprite drawing and collision detection.
  
2. Graphical Display
   - Utilizes SDL2 for rendering.
   - Renders at a resolution of **64x32 pixels**, scaled for modern displays.
   - Efficient handling of drawing sprites with XOR-based pixel toggling.

3. Input Handling
   - Maps CHIP-8 hexadecimal keys **(0x0-0xF)** to a standard keyboard layout.
   - Real-time handling of key presses and releases using SDL events.

4. Timers
   - Emulates the CHIP-8 Delay and Sound timers.
   - Timers decrease at 60 Hz, synchronized with the emulator’s frame rate.

5. Debugger Mode
   - Step-by-step execution for debugging.
   - Displays CPU state, memory, and current opcode for each step.

6. ROM Support
   - Load and play classic CHIP-8 ROMs such as Pong, Tetris, and Space Invaders.
   - Built-in support for **.ch8** files.

## 🖥 How to Use

### Requirements
- Go 1.20+
- SDL2 library
	- Install SDL2 via your package manager or follow SDL2 installation guide.

### Setup

## 🎮 Controls

| CHIP-8 Key | Keyboard Key |
|------------|--------------|
| `1`        | `1`          |
| `2`        | `2`          |
| `3`        | `3`          |
| `C`        | `4`          |
| `4`        | `Q`          |
| `5`        | `W`          |
| `6`        | `E`          |
| `D`        | `R`          |
| `7`        | `A`          |
| `8`        | `S`          |
| `9`        | `D`          |
| `E`        | `F`          |
| `A`        | `Z`          |
| `0`        | `X`          |
| `B`        | `C`          |
| `F`        | `V`          |


## Contributing
Contributions are welcome! Please fork the repository, create a new branch, and submit a pull request.

## License
This emulator is open-sourced under the MIT License. See the LICENSE file for more information.
