# Gomoku Game Usage Example

## Introduction
This document provides an overview of how to play the Gomoku game, including setup instructions and gameplay rules.

## Setup
1. Clone the repository:
   ```
   git clone https://github.com/yourusername/gomoku.git
   ```
2. Navigate to the project directory:
   ```
   cd gomoku
   ```
3. Install the necessary dependencies:
   ```
   go mod tidy
   ```

## Running the Game
### Command Line Interface
To start the game using the command line interface, run:
```
go run internal/ui/cli/main.go
```

### Web Interface
To start the game using the web interface, run:
```
go run internal/ui/web/server.go
```
Then open your web browser and navigate to `http://localhost:8080`.

## Gameplay Instructions
1. The game is played on a 15x15 board.
2. Players take turns placing their stones (black and white) on the board.
3. The first player to align five stones horizontally, vertically, or diagonally wins the game.
4. Players can choose to play against each other or against the AI.

## Rules
- Players must place their stones on empty intersections.
- The game ends when one player achieves five in a row or when the board is full (draw).

## Conclusion
Enjoy playing Gomoku! For more information, refer to the README.md file or the documentation in the `pkg` directory.