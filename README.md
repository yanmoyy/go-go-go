# Go-Go-Go: Real-time Multiplayer Strategy Game

A full-stack multiplayer [Alkkagi (Stone-Shooting)](https://en.wikipedia.org/wiki/Alkkagi) game built entirely in Go, featuring a `terminal user interface (TUI)`, `real-time WebSocket communication` and `game history recording`.

![Game Screenshot](img/mainPage.png)

## Project Overview
### 🖥️ **Terminal User Interface (TUI)**
- **Modern TUI**: Built with [Bubble Tea](https://github.com/charmbracelet/bubbletea) framework for a gui-like experience in the terminal.
- **Real-time Rendering**: Render game state and animations by using 
- **Intuitive Controls**: Easy-to-use navigation and game controls
- **Cross-platform**: Works on any terminal that supports ANSI colors

### 🌐 **Real-time Multiplayer**
- **WebSocket Communication**: Instant game state synchronization between players
- **Live Chat System**: In-game messaging for player interaction
- **Session Management**: Automatic player matching and game session handling
- **Event-driven Architecture**: Efficient real-time event broadcasting

### 🎮 **Game Mechanics**
- **Physics Engine**: Realistic stone collision and movement simulation
- **Turn-based Strategy**: Strategic gameplay with alternating turns
- **Win Conditions**: Eliminate opponent's stones to win
- **Real-time Animations**: Visual feedback for all game actions

### �� **Data Persistence**
- **Game History**: Automatic recording of all game sessions
- **PostgreSQL Integration**: Reliable data storage with JSONB for flexible game records
- **Replay System**: Store complete game state for analysis and replay
- **Session Tracking**: Comprehensive logging of player actions and game events

## 🏗️ Technical Architecture

### **Backend (Server)**
- **Go 1.24.4**: Modern Go with latest features
- **Gorilla WebSocket**: Robust WebSocket implementation
- **PostgreSQL**: Reliable data persistence with JSONB support
- **Environment Configuration**: Secure configuration management
- **Structured Logging**: Comprehensive logging with slog

### **Frontend (Client)**
- **Bubble Tea**: Elegant TUI framework for terminal applications
- **Lip Gloss**: Beautiful styling and layout system
- **Real-time Updates**: Live game state synchronization
- **Error Handling**: Graceful error recovery and user feedback

### **Game Engine**
- **Physics Simulation**: Realistic collision detection and movement
- **Event System**: Comprehensive game event management
- **State Management**: Immutable game state with event sourcing
- **Multiplayer Sync**: Deterministic game state across clients

## 🚀 Getting Started

### Prerequisites
- Go 1.24.4 or later
- PostgreSQL database
- Terminal with ANSI color support

### Installation

1. **Clone the repository**
   ```bash
   git clone https://github.com/yanmoyy/go-go-go.git
   cd go-go-go
   ```

2. **Set up environment variables**
   ```bash
   cp .env.example .env
   # Edit .env with your database credentials
   ```

3. **Run database migrations**
   ```bash
   psql -d your_database -f sql/schema/001_record.sql
   ```

4. **Start the server**
   ```bash
   go run cmd/server/main.go
   ```

5. **Start the client**
   ```bash
   go run cmd/client/main.go
   ```

## �� How to Play

1. **Connect**: Launch the client and connect to the server
2. **Match**: Wait for another player to join your session
3. **Strategize**: Plan your shots to eliminate opponent stones
4. **Shoot**: Use velocity controls to aim and shoot your stones
5. **Win**: Be the last player with stones remaining on the board

### Game Rules
- Each player starts with 10 stones (White vs Black)
- Take turns shooting stones across the 100x100 game board
- Stones collide realistically with physics simulation
- Stones that go out of bounds are eliminated
- Last player with stones remaining wins

### Contributors
- [@yanmoyy](https://github.com/yanmoyy)
- [@Gutssssssssss](https://github.com/Gutssssssssss)