# Slides for daily sessions

## Day 1

- Intro
- Topics to be covered
- Why Go?
- History
- Go Playgroud
- Data types
- Varaibles

Hello World

```go
package main

import "fmt"

func main() {
 fmt.Println("Hello Wordl!")
}

```

Hello Web

```go
package main

import (
 "log"
 "net/http"
)

func helloHandler(writer http.ResponseWriter, request *http.Request) {
 writer.Write([]byte("<h1>Hello, web!</h1>"))
}

func main() {
 http.HandleFunc("/hello", helloHandler)
 err := http.ListenAndServe("localhost:8080", nil)
 log.Fatal(err)
}

```

## Day 2

### Environment Setup

- Install Go
- Install Visual Studio Code and Analysis Tools

### Understanding the basic program

- how to run the code?
- Go CLI
  - `go build`
  - `go run`
  - `go fmt`
  - `go install`
  - `go get`
  - `go test`

### Packages in Go

- Executable packages
- Reusable packages

### Imports

- Standard libraries
- Reusable modules

### Use Case

- Playing Card functionality
  - New Deck
  - Print
  - Suffle
  - Deal
  - SaveToFile
  - DeckFromFile

### Variables

- Declaration
- All must be used
- no private scope
- type conversion

### Arrays / Slice

- What is a slice?
- Must be of same type
- How to declare and loop through?

## Day 3

### Flow controll statements

- `if / else`
- `for`
- `switch`
- `defer`

### Zero values in Go

| Type          | Zero Value  |
| -----------   | ----------- |
| Integer       | 0           |
| Floating point| 0.0         |
| Boolean       | false       |
| String        | ""          |
| Pointer       | nil         |
| Interface     | nil         |
| Slice         | nil         |
| Map           | nil         |
| Channel       | nil         |
| Function      | nil         |

### Playing Cards

- creating new deck
- printing
- deal
- toString

## Day 4

### Functions

- A language that supports first class functions
- example: [functions](./examples/functions/main.go)
- example: [methods](./examples/methods/main.go)
- example: [more on functions](./examples/functions_variables/main.go)

### Byte slice

- Converting a string to byte slice
- example: [byte slice](./examples/byte_slice/main.go)
