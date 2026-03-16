# Gomoku Game

## Overview
Gomoku, also known as Five in a Row, is a classic board game where two players take turns placing their pieces on a grid. The objective is to be the first to get five of their pieces in a row, either horizontally, vertically, or diagonally.

## Features
- Play against another player or a simple AI.
- Command-line interface for easy interaction.
- Web interface for a more visual experience.
- Game rules implemented to ensure fair play.

## Installation
To install the Gomoku game, follow these steps:

1. Clone the repository:
   ```
   git clone <repository-url>
   ```
2. Navigate to the project directory:
   ```
   cd gomoku
   ```
3. Install the necessary dependencies:
   ```
   go mod tidy
   ```

## Usage
To start the game, you can choose between the command-line interface or the web interface.

### Command-Line Interface
Run the following command to start the CLI version:
```
go run internal/ui/cli/main.go
```

### Web Interface
To start the web server, run:
```
go run internal/ui/web/server.go
```
Then, open your web browser and navigate to `http://localhost:8080` to play the game.

## Contributing
Contributions are welcome! Please feel free to submit a pull request or open an issue for any suggestions or improvements.

## License
This project is licensed under the MIT License. See the LICENSE file for more details.