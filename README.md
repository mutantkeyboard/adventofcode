# Advent of Code 2024

This repository contains my solutions for the [Advent of Code 2024](https://adventofcode.com/2024) challenges, implemented in Go.

## Project Structure

```
advent-of-code-2024/
├── Makefile
├── README.md
├── 1/
│   ├── README.md
│   ├── solution.go
│   ├── solution_test.go
│   └── input.txt
├── 2/
│   ├── README.md
│   ├── solution.go
│   ├── solution_test.go
│   └── input.txt
└── ...
```

Each day's challenge is organized in its own directory, containing:
- `README.md` - Problem description and notes
- `solution.go` - Go implementation of the solution
- `solution_test.go` - Tests for the solution
- `input.txt` - Puzzle input data

## Usage

### Prerequisites
- Go 1.21 or higher installed
- Make utility installed

### Generating Files

Generate a single day:
```bash
make generate_day day=1  # Generates files for Day 1
```

Generate multiple days at once:
```bash
make generate_days start_day=1 end_day=8  # Generates files for Days 1-8
```

### Running Solutions

To run a solution for a specific day:
```bash
make run day=1  # Runs solution for Day 1
```

### Running Tests

To run tests for a specific day:
```bash
make test day=1  # Runs tests for Day 1
```

### Cleaning Build Artifacts

To clean up any build artifacts:
```bash
make clean
```

## Progress

- [ ] Day 1
- [ ] Day 2
- [ ] Day 3
- [ ] Day 4
- [ ] Day 5
- [ ] Day 6
- [ ] Day 7
- [ ] Day 8
- [ ] Day 9
- [ ] Day 10
- [ ] Day 11
- [ ] Day 12
- [ ] Day 13
- [ ] Day 14
- [ ] Day 15
- [ ] Day 16
- [ ] Day 17
- [ ] Day 18
- [ ] Day 19
- [ ] Day 20
- [ ] Day 21
- [ ] Day 22
- [ ] Day 23
- [ ] Day 24
- [ ] Day 25

## License

This project is open source and available under the MIT License.