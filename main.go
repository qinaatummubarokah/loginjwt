package main

import (
    "fmt"

)

func main() {
    e := NewRouter()

    fmt.Println("starting web server at http://localhost:3636/")
    e.Start(":3636")
}