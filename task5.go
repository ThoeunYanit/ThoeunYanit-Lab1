package main

import (
    "fmt"
    "encoding/hex"
    "encoding/base64"
)

func main() {
    var text string

    fmt.Print("Enter text: ")
    fmt.Scanln(&text)

    // Binary
    fmt.Print("Binary: ")
    for i := 0; i < len(text); i++ {
        fmt.Printf("%08b ", text[i])
    }

    // Hexadecimal
    fmt.Println("\nHexadecimal:", hex.EncodeToString([]byte(text)))

    // Base64
    fmt.Println("Base64:", base64.StdEncoding.EncodeToString([]byte(text)))
}

