Exercise: Word Frequency Counter

Objective: The objective of this exercise is to reinforce the concepts of maps in Go by creating a word frequency counter. Participants will write a program that reads a text file, counts the frequency of each word, and prints the result as a map.

Instructions:
1. Ask the participants to create a new Go file, e.g., `word_frequency.go`, on their machines.
2. Instruct them to implement a program that performs the following steps:
   - Read a text file (e.g., `sample.txt`) containing multiple words.
   - Split the text into individual words.
   - Create a map to store the word frequency.
   - Iterate over the words and increment the count for each word in the map.
   - Print the word frequency map.

Solution:

```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Open the text file
	file, err := os.Open("sample.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file
	scanner := bufio.NewScanner(file)

	// Create a map to store word frequency
	wordFreq := make(map[string]int)

	// Process each line of the file
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line)
		for _, word := range words {
			// Increment word frequency
			wordFreq[word]++
		}
	}

	if scanner.Err() != nil {
		fmt.Println("Error scanning file:", scanner.Err())
		return
	}

	// Print word frequency map
	for word, freq := range wordFreq {
		fmt.Printf("%s: %d\n", word, freq)
	}
}
```

Explanation:
1. The program starts by opening the text file using `os.Open` and handling any potential errors.
2. A scanner is created to read the file line by line.
3. A map named `wordFreq` is initialized to store the word frequencies. Each word will be a key, and the corresponding value will represent its frequency.
4. The program loops through each line of the file using `scanner.Scan()`.
5. Inside the loop, the line is split into individual words using `strings.Fields`.
6. For each word, its frequency is incremented in the `wordFreq` map.
7. After processing the file, the program checks for any scanning errors using `scanner.Err()`.
8. Finally, the program iterates over the `wordFreq` map and prints each word along with its frequency using `fmt.Printf`.

Note: Participants can create a `sample.txt` file with some text to test the program and see the word frequencies in action.

This exercise allows participants to practice working with maps, file I/O, string manipulation, and looping. They will gain hands-on experience in creating a word frequency counter, which demonstrates the practical application of maps in Go programming.