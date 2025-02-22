List Selector CLI

Overview

The **List Selector CLI** is a command-line application built using **Go** and the **Cobra** package. It allows users to navigate a list of options using arrow keys and execute associated actions.

## Features

* Navigate a menu using arrow keys
* Execute various actions upon selection
* Retrieve user input
* Exit functionality

## Installation

### Prerequisites

Ensure you have **Go** installed on your system. If not, install it from .

Clone the Repository

```
git clone https://github.com/BiodigitalJaz/go-cli-terminal.git

cd list-selector-cli
```

Install Dependencies

```
go mod tidy
```

## Usage

Run the CLI

```
go run main.go
```

Options

| Option Number | Description            |
|--------------|------------------------|
| 1            | Print "Hello, world!"   |
| 2            | Display the current date |
| 3            | Print "Goodbye!"        |
| 4            | Do nothing              |
| 5            | Allow user input        |
| 6            | Exit the program        |

## Project Structure

```
go-cli-terminal/
│── main.go                # Entry point of the application
│── controllers/
│   ├── selector.go        # Handles menu navigation logic and rendering text
│── models/
│   ├── option.go          # Defines the Option struct
|── actions/
    |── actions.go         # Provides sample actions for our options to execute on
│── go.mod                 # Go module file
│── go.sum                 # Dependencies checksum