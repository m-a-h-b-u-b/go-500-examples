package main

import (
    "bufio"
    "fmt"
    "net"
    "os"
)

func main() {
    // Define server address
    serverAddr := "127.0.0.1:8080" // change to your server IP:Port

    // Connect to server
    conn, err := net.Dial("tcp", serverAddr)
    if err != nil {
        fmt.Println("Error connecting to server:", err)
        return
    }
    defer conn.Close()
    fmt.Println("Connected to server:", serverAddr)

    // Read input from user
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter message to send: ")
    msg, _ := reader.ReadString('\n')

    // Send message to server
    _, err = conn.Write([]byte(msg))
    if err != nil {
        fmt.Println("Error sending message:", err)
        return
    }

    // Receive response from server
    response, err := bufio.NewReader(conn).ReadString('\n')
    if err != nil {
        fmt.Println("Error reading response:", err)
        return
    }
    fmt.Println("Server response:", response)
}
