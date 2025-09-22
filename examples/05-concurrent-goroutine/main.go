package main

import (
    "fmt"
    "time"
)

func sayHello(name string) {
    for i := 0; i < 3; i++ {
        fmt.Println("Hello", name, "iteration", i)
        time.Sleep(500 * time.Millisecond)
    }
}

func main() {
    go sayHello("Alice")
    go sayHello("Bob")

    // Wait for goroutines to finish
    time.Sleep(2 * time.Second)
    fmt.Println("Done!")
}
