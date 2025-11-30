package main

import (
	"fmt"
	"strings"
)

func main() {
    fmt.Print("Enter a string: ")

    var userInput string
	
    fmt.Scanf("%[^\n]", &userInput)

    // Convert to lower-case so comparisons are case-insensitive
    s := strings.ToLower(userInput)

    // Trim spaces at the end before checking suffix
    trimmed := strings.TrimSpace(s)

    // Check the three required conditions
    startsWithI := strings.HasPrefix(trimmed, "i")
    endsWithN := strings.HasSuffix(trimmed, "n")
    containsA := strings.Contains(trimmed, "a")

    if startsWithI && endsWithN && containsA {
        fmt.Println("Found!")
    } else {
        fmt.Println("Not Found!")
    }
}
