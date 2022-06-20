package main

import (
  "fmt"

  "example.com/greetings"
)

func main() {
    message := greetings.Hello("kattsun")
    fmt.Println(message)
}
