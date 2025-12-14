# Razd Go Example

A simple Go project demonstrating [Razd](https://github.com/razd-cli/razd) workflow.

## Prerequisites

- [Razd CLI](https://github.com/razd-cli/razd)

## Quick Start

```bash
# Install tools and run the project
razd
```

## Available Commands

| Command        | Description                          |
| -------------- | ------------------------------------ |
| `razd`         | Install dependencies and run the app |
| `razd install` | Install Go dependencies              |
| `razd dev`     | Run the application                  |
| `razd build`   | Build binary to `bin/app`            |

## Project Structure

```
.
├── Razdfile.yml   # Razd configuration
├── main.go        # Application entry point
├── go.mod         # Go module definition
└── go.sum         # Go dependencies checksum
```

## Dependencies

- [github.com/fatih/color](https://github.com/fatih/color) - Color output for terminal

## License

MIT
