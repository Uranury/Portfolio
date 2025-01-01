package main

import (
	"fmt"
	"os"
	"strings"
	"sync"
)

var Wg sync.WaitGroup
var mu sync.Mutex

func wordCount(text string) map[string]int {
	wordFreq := make(map[string]int)
	words := strings.Fields(strings.ToLower(text))
	for _, word := range words {
		wordFreq[word]++
	}
	return wordFreq
}

func main() {
	files := [...]string{"a.txt", "b.txt", "c.txt"}

	/*
		texts := []string{
			"In the heart of the bustling city, amidst the noise and rush, there lies a park where tranquility reigns supreme. It is a place where time slows down, and the mind can take a breath from the constant activity of the urban jungle.",
			"On the edge of the park, the trees sway gently with the breeze, their leaves dancing in the wind as if they were performing a silent ballet for the few who stop to listen. The scent of fresh flowers mingles with the earthy aroma of damp soil, creating a calming atmosphere.",
			"At the center of the park, there is a small pond, its surface reflecting the sky above and the green foliage surrounding it. Ducks glide gracefully across the water, their quacks echoing softly in the quietude. People gather on the benches, some reading, others simply enjoying the serenity.",
			"As the sun begins to set, painting the sky in hues of orange and pink, the park transforms into a golden haven. The quiet chatter of friends and families fills the air, but there is still a sense of peace that envelops the park, even as the day turns to night.",
			"For those who seek solace in nature, the park is an oasis. It serves as a reminder that even in the busiest of cities, there is always a place to find peace and reconnect with the world around us.",
		}

		// Start goroutines for writing files
		for idx, file := range files {
			Wg.Add(1)
			go func(filename, text string) {
				defer Wg.Done() // Defer should be inside the goroutine
				err := WriteFile(filename, text)
				if err != nil {
					fmt.Printf("Error writing file %s: %v\n", filename, err)
				}
			}(file, texts[idx])
		}
	*/
	// Start goroutines for reading files and counting word frequencies
	wordFrequencies := make(map[string]map[string]int)
	for _, file := range files {
		Wg.Add(1)
		go func(filename string) {
			defer Wg.Done() // Defer should be inside the goroutine
			content, err := os.ReadFile(filename)
			if err != nil {
				fmt.Printf("Error reading file %s: %v\n", filename, err)
				return
			}
			wordFreq := wordCount(string(content))

			// Lock before modifying the shared map
			mu.Lock()
			wordFrequencies[filename] = wordFreq
			mu.Unlock()
		}(file)
	}

	// Wait for all goroutines to finish
	Wg.Wait()

	// Print the word frequencies
	for file, freq := range wordFrequencies {
		fmt.Printf("Word frequencies in %s:\n", file)
		for word, count := range freq {
			fmt.Printf("  %s: %d\n", word, count)
		}
	}
	fmt.Print("Successfully executed.")
}

func WriteFile(filename string, data string) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return fmt.Errorf("failed to open a file: %w", err)
	}
	defer file.Close()
	_, err = file.WriteString(data)
	if err != nil {
		return fmt.Errorf("failed to write into a file: %w", err)
	}
	return nil
}
