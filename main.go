package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "unicode"
)

// Function to count words in a file
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

// Function to count lines in a file
func countLines(filename string) (int, error) {
    file, err := os.Open(filename)
    if err != nil {
        return 0, fmt.Errorf("could not open file: %v", err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    lineCount := 0
    for scanner.Scan() {
        lineCount++
    }

    if err := scanner.Err(); err != nil {
        return 0, fmt.Errorf("error reading file: %v", err)
    }

    return lineCount, nil
}

// Function to count characters in a file
func countCharacters(filename string) (int, error) {
    file, err := os.Open(filename)
    if err != nil {
        return 0, fmt.Errorf("could not open file: %v", err)
    }
    defer file.Close()

    reader := bufio.NewReader(file)
    charCount := 0
    for {
        char, _, err := reader.ReadRune()
        if err != nil {
            break
        }
        charCount++
    }

    return charCount, nil
}

func main() {
    // Check for filename argument
    if len(os.Args) < 2 {
        fmt.Println("Usage: go run main.go <filename>")
        os.Exit(1)
    }

    filename := os.Args[1]

    // Count words and unique word frequency
    wordCount, wordFrequency, err := countWords(filename)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error: %v\n", err)
        os.Exit(1)
    }

    // Count lines
    lineCount, err := countLines(filename)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error: %v\n", err)
        os.Exit(1)
    }

    // Count characters
    charCount, err := countCharacters(filename)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error: %v\n", err)
        os.Exit(1)
    }

    // Display statistics
    fmt.Printf("File Statistics for '%s':\n", filename)
    fmt.Printf("Lines: %d\n", lineCount)
    fmt.Printf("Words: %d\n", wordCount)
    fmt.Printf("Characters: %d\n", charCount)
    fmt.Println("\nUnique Word Frequency:")

    for word, frequency := range wordFrequency {
        fmt.Printf("%s: %d\n", word, frequency)
    }
}
