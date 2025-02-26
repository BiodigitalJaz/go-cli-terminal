# Interactive CLI Application with Cobra

A structured Golang CLI application using the Cobra framework, implementing an MVC-like pattern with interactive menus.

## Features

- **Structured Architecture**: Organized using models, controllers, and actions
- **Command-line Interface**: Traditional CLI commands using Cobra
- **Interactive Mode**: TUI-based menu system for easy navigation
- **Extensible Design**: Easy to add new commands and options

## Project Structure

```
.
├── actions/
│   ├── project/
│   │   └── project.go
│   └── user/
│       └── user.go
├── cmd/
│   ├── interactive.go
│   └── root.go
├── controllers/
│   ├── project.go
│   └── user.go
├── models/
│   └── option.go
├── go.mod
├── go.sum
├── main.go
└── README.md
```

## Installation

Clone the repository and build the application:

```bash
git clone https://github.com/biodigitalJaz/go-cli-terminal.git
cd cli-app
go mod tidy
go build -o cli-app
```

## Usage

### Command Mode

Run specific commands directly:

```bash
# User commands
./cli-app user list
./cli-app user create
./cli-app user delete username

# Project commands
./cli-app project list
./cli-app project create
./cli-app project delete projectname
```

### Interactive Mode

Launch the interactive menu for a guided experience:

```bash
./cli-app interactive
```

This starts a TUI-based menu system where you can navigate through options using arrow keys.

## Architecture

### Models

The core model is the `Option` type which represents menu items:

```go
type Option struct {
    Name        string
    Description string
    Action      func() error
}
```

These are collected in `OptionList` to build interactive menus.

### Controllers

Controllers are implemented as Cobra commands, organizing the application's functionality:

- `UserCmd`: Handles user management operations
- `ProjectCmd`: Handles project management operations
- `InteractiveCmd`: Launches the interactive menu system

### Actions

Business logic is implemented in action functions, separated by domain:

- `user.ListUsers()`, `user.CreateUser()`, etc.
- `project.ListProjects()`, `project.CreateProject()`, etc.

## Dependencies

- [github.com/spf13/cobra](https://github.com/spf13/cobra): Command line interface framework
- [github.com/manifoldco/promptui](https://github.com/manifoldco/promptui): Interactive prompt library

## License

MIT

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request
