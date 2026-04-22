
# MarsGo: Marseille Chess Logic Library
**Version 1.0.0-alpha** | Author: [itsVentie](https://github.com/itsVentie)

A lightweight, robust Go library for implementing **Marseille Chess** (double-move variant) rules. Built on top of `notnil/chess`, MarsGo handles the complex sub-turn logic and check-validation specific to the Marseille variant.

---

## ♟️ What is Marseille Chess?

Marseille Chess is a tactical variant where players move **twice** per turn.
* **Balanced Start:** White moves only once on the first turn.
* **The Check Constraint:** If a player gives Check on their 1st sub-move, their turn ends immediately.
* **Tempo:** Games are significantly faster and more volatile than standard chess.

---

## 📂 Library Structure

```text
MarsGo/
├── api/                # API definitions (gRPC/Proto and OpenAPI)
├── cmd/
│   ├── cli/            # Interactive Terminal Interface
│   └── server/         # Backend server implementation (Engine API)
├── frontend/           # Web-based UI (Next.js/React)
├── internal/           # High-performance core (Bitboards, Search, Eval)
│   ├── bitboard/       # Bitboard representation of the board
│   ├── search/         # Search algorithms (Minimax/Alpha-Beta)
│   └── evaluation/     # Board scoring heuristics
├── pkg/                # Publicly accessible libraries
│   ├── engine/         # Marseille Protocol & Game State Machine
│   ├── notation/       # FEN & PGN parsing
│   └── uci/            # UCI protocol handler for GUI compatibility
├── tests/              # Integration and Benchmark suites
├── go.mod              # Dependency management
└── README.md           # Project Documentation
```

---

## 🛠 Installation

```bash
go get [github.com/itsVentie/MarsGo](https://github.com/itsVentie/MarsGo)
```

---

## 🚀 Running the Project

### 1. CLI Mode (Terminal)
Play directly in your terminal using the built-in CLI tool:
```bash
go run cmd/cli/main.go
```

### 2. GUI Mode (Fullstack)
To run the web interface, you need to start both the backend server and the frontend development server.

**Start the Engine Server:**
```bash
go run cmd/server/main.go
```

**Start the Frontend:**
```bash
cd frontend
pnpm install
pnpm dev
```

---

## ⚙️ Core Logic Implementation

### 🛡 Deep Scan Validation
MarsGo implements a **Geometric Attack Scanner**. Since standard FEN manipulation can sometimes "blind" standard engines to checks in non-standard positions, MarsGo manually verifies king safety after every sub-move to ensure compliance with the Marseille "Check-End" rule.

### 🔄 State Reconstruction
The library uses a **History-Based Reconstruction** for the `Undo()` function. This ensures that even in complex double-move sequences, the internal board state remains 100% consistent with the move history.

---

## 🏗 Roadmap (Alpha v1)

* [ ] **v1.1.0:** Minimax AI integration specifically tuned for $40^2$ branch factors.
* [ ] **v1.2.0:** Full PGN support for Marseille-style move notation.
* [ ] **v2.0.0:** Plug-and-play move validation for web-based multiplayers.

---

## Contributing
Issues and PRs are welcome. Help us make the Marseille variant more accessible to the Go community.

## 📄 License
MIT License. Created by [itsVentie](https://github.com/itsVentie).
```
