package main

import "fmt"

func xorEncrypt(text string, key string) string {
    result := ""

    for i := 0; i < len(text); i++ {
        result += string(text[i] ^ key[i%len(key)])
    }

    return result
}

func main() {
    var text string
    var key string

    fmt.Print("Enter plaintext: ")
    fmt.Scanln(&text)

    fmt.Print("Enter key: ")
    fmt.Scanln(&key)

    // Encrypt
    ciphertext := xorEncrypt(text, key)
    fmt.Printf("Ciphertext: %x\n", ciphertext)

    // Decrypt using the same function and key
    plaintext := xorEncrypt(ciphertext, key)
    fmt.Println("Decrypted text:", plaintext)
}

