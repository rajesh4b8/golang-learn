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

## Day 5

### Byte slice

- Converting a string to byte slice
- example: [byte slice](./examples/byte_slice/main.go)

### Completed the usecase for playing cards

- SaveToFile
- DeckFromFile
- Shuffle
- Discussed some standard packages

### Basics of go testing

- created tests for deck.go

## Day 6

### Struts

- Structs definition
- Declaring the structs
- Updating the structs
- Embedded structs
- Receiver functions / methods
- Pass by value
- Pass by pointer

### Pinters

- Use `&` to access the pointer (address)
- Use `*` to access the value from pointer
- `*` in front for a type incates we are passing pointer
- Pointer indirection when using pointer recievers
- You don't need to pass pointers for reference types, as the reference types are a reference to value types

### Value Types / Reference types

| Value Type       | Referece Type |
| ---------------- | ------------- |
| int              | slices        |
| float            | maps          |
| string           | channels      |
| bool             | pointers      |
| structs          | functions     |

### Assignment 1

- Print numbers from 1 to 10 and print whether it's even or odd
- Sample output

``` text
1 - odd
2 - even
3 - odd
4 - even
...
```

## Day 7

### Maps

- All the keys should be of same type
- All the values should be of same type
- Maps vs Structs

| Map                 | Struct               |
|-------------------- |----------------------|
| All keys - same type| different types      |
| related properties  | represent something    |
| All values - same type|can be different |
| keys - can be added | Define them at compile time |
| keys are indexed   | keys not indexed |
| Reference Type     | Value Type |

### Interfaces

- Example

## Day 8

### GOPATH / GOROOT

- GOROOT is a variable that defines where your Go SDK is located. You do not need to change this variable, unless you plan to use different Go versions.
- GOPATH is a variable that defines the root of your workspace. By default, the workspace directory is a directory that is named go within your user home directory (~/go for Linux and MacOS, %USERPROFILE%/go for Windows). GOPATH stores your code base and all the files that are necessary for your development. You can use another directory as your workspace by configuring GOPATH for different scopes. GOPATH is the root of your workspace and contains the following folders:
  - src/: location of Go source code (for example, .go, .c, .g, .s).
  - pkg/: location of compiled package code (for example, .a).
  - bin/: location of compiled executable programs built by Go.

### GoLang developer workspace

- Define your GOPATH
- Setup your project from https://github.com/rajesh4b8/users-api-batch-2

### MVC Architecture

- MVC is abbreviated as Model View Controller is a design pattern created for developing applications specifically web applications. As the name suggests, it has three major parts. The traditional software design pattern works in an "Input - Process - Output" pattern whereas MVC works as "Controller -Model - View" approach

### Use Case - Users API

- Create User
- Get User
- Update User
- Delete User

## Day 9

- Use case development
  - User controller with get user
  - created User model with type struct
  - created users service
  - error handling with errors util

## Day 10

- Git basics
  - checkout a branch -> get the code from a branch
