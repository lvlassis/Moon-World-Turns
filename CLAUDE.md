# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

Moon-World-Turns is a turn-based battle game built in Go using the g3n 3D engine. The game features Pokemon-style battles with 3D graphics, HUD elements, and an event-driven architecture.

## Build and Run Commands

```bash
# Development - run directly
make dev
# or
go run main.go

# Production build
make build
# Output: build/main

# Direct build (alternative)
go build .
chmod +x moon-world-turns
./moon-world-turns
```

## Architecture Overview

### High-Level Structure

The codebase follows a clean separation between game logic (game), presentation (view), and application control (app):

- **internal/app**: Application entry point that orchestrates initialization
- **internal/game**: Core game logic - battle system, characters, actions, stages
- **internal/view**: Rendering and UI layer using g3n engine
- **internal/events**: Event bus for decoupled communication between layers
- **data/**: YAML configuration files for characters and stages

### Key Architectural Patterns

**1. Game-View Separation**
- Game layer (`internal/game`) contains pure game logic with no rendering concerns
- View layer (`internal/view`) handles all g3n rendering, cameras, and GUI
- Communication happens through the event bus (`internal/events`)

**2. Data-Driven Design**
- Characters, stages, and game entities are defined in YAML files (`data/`)
- Game layer loads YAML into structs (e.g., `CharacterData`, `StageData`)
- Separation between immutable data (e.g., `CharacterData`) and mutable state (e.g., `Character.Life`)

**3. Scene-Based Rendering**
- `view.Scene` interface defines contract for renderable scenes
- `BattleScene` implements dual-camera rendering: 3D world + 2D GUI overlay
- Window manages the render loop and scene lifecycle

### Critical Components

**Battle System (internal/game/battle.go)**
- Manages turn-based combat between two characters
- Tracks current turn and which character is choosing an action
- `DoAction()` validates turn order before executing actions
- Speed attribute determines who goes first

**Action System (internal/game/action.go, internal/game/actions/)**
- Actions are defined as structs with `Effect` functions
- Effects take `actor` and `target` parameters to modify character state
- Actions are organized in `internal/game/actions/` package (e.g., `fireball.go`)

**Event Bus (internal/events/bus.go)**
- Simple pub/sub system for decoupled communication
- Subscribe to events by name with callback functions
- Used to communicate battle events from game layer to view layer

**BattleScene Rendering (internal/view/battle_scene.go)**
- Maintains two cameras: `cam` (3D world) and `guiCam` (2D orthographic GUI)
- `Render()` calls renderer twice: once for 3D scene, once for GUI
- `Update()` synchronizes HUD (health bars) with engine state
- `OnResize()` handles window resizing for both cameras

**HUD Components (internal/view/hud/)**
- `HealthBar`: Displays character name and life percentage
- `ActionButton`: Clickable buttons for character actions
- All HUD elements are g3n GUI components added to `guiRoot` panel

### Data Flow

1. **Initialization** (internal/app/run.go):
   - Load YAML data → Create game models → Create view models → Assemble scene
   - Event bus is created and passed to both game layer (Battle) and view (BattleScene)

2. **Battle Flow**:
   - Battle.Start() determines first character based on speed
   - User clicks action button → Battle.DoAction() validates and executes
   - Game layer updates character state (e.g., Life)
   - View.Update() reads updated state and updates HUD

3. **Render Loop** (internal/view/window.go):
   - Clear buffers → Scene.Update() → Scene.Render() → Repeat

## Important Conventions

- YAML files in `data/` define all game content (characters, stages)
- Assets referenced in YAML (images) are loaded from paths like `assets/charmander.png`
- Character stats: `MaxLife` (immutable from YAML) vs `Life` (mutable game state)
- Battle turn logic: Always check `battle.choosing` before allowing actions
- GUI positioning: Health bars anchored to screen edges, adjusted in `OnResize()`

## Dependencies

- **g3n/engine v0.2.0**: 3D rendering engine (OpenGL wrapper for Go)
- **gopkg.in/yaml.v3**: YAML parsing for data files
- **glfw**: Window and input handling (indirect dependency via g3n)

## Current Development State

Based on recent commits:
- Action system is being implemented (see `internal/game/actions/`)
- HUD components (health bars, action buttons) recently abstracted into `internal/view/hud/`
- Battle rendering architecture recently refactored to merge logic into BattleScene
- MVP features complete: attack buttons, health bars, stage rendering, base architecture
