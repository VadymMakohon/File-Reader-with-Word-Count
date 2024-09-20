package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "unicode"
)

func countWords(filename string) (int, map[string]int, error) {
    file, err := os.Open(filename)
    if err != nil {
        return 0, nil, fmt.Errorf("could not open file: %v", err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanWords)

    wordCount := 0
    wordFrequency := make(map[string]int)

    for scanner.Scan() {
        word := strings.ToLower(scanner.Text())
        word = strings.TrimFunc(word, func(r rune) bool {
            return !unicode.IsLetter(r)
        })
        if word != "" {
            wordCount++
            wordFrequency[word]++
        }
    }

    if err := scanner.Err(); err != nil {
        return 0, nil, fmt.Errorf("error reading file: %v", err)
    }

    return wordCount, wordFrequency, nil
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
