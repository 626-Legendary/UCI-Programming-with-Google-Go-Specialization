package main

import "fmt"

func main() {
    fmt.Print("Enter a floating point number: ")

    var floatNum float64
    _, err := fmt.Scan(&floatNum)

    if err != nil {
        fmt.Println("Error reading input:", err)
        return
    }

    // Truncate by converting float → int (drops decimal part)
    var truncated int = int(floatNum)

    fmt.Printf("Truncated integer: %d", truncated)
}
