# Day 1:

- Intro
- Topics to be covered
- Why Go?
- History
- Go Playgroud
- Data types 
- Varaibles 
- Functions


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


# Day 2
- Installing go
    - Quick steps
    - IDE (VS Code)

- Understanding the basic program
    - how to run the code?
    - Go CLI
        - ```go build``` 
        - ```go run``` 
        - ```go fmt``` 
        - ```go install``` 
        - ```go get``` 
        - ```go test``` 

- Packages in Go
    - Executable packages
    - Reusable packages

- Imports 
    - Standard libraries 
    - Reusable modules 

- Use Case
    - Playing Card functionality
        - New Deck
        - Print
        - Suffle
        - Deal
        - SaveToFile
        - DeckFromFile

- Variables
    - Declaration
    - All must be used
    - no private scope 
    - type conversion 

- Arrays / Slice
    - What is a slice?
    - Must be of same type
    - How to declare and loop through?
