package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

// Function to count words in a file
func countWords(filename string) (int, error) {
    file, err := os.Open(filename)
    if err != nil {
        return 0, fmt.Errorf("could not open file: %v", err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanWords) // Set scanner to split by words

    wordCount := 0
    for scanner.Scan() {
        wordCount++
    }

    if err := scanner.Err(); err != nil {
        return 0, fmt.Errorf("error reading file: %v", err)
    }

    return wordCount, nil
}

func main() {
    // Check for filename argument
    if len(os.Args) < 2 {
        fmt.Println("Usage: go run main.go <filename>")
        os.Exit(1)
    }

    filename := os.Args[1]
    wordCount, err := countWords(filename)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error: %v\n", err)
        os.Exit(1)
    }

    fmt.Printf("The file '%s' contains %d words.\n", filename, wordCount)
}
