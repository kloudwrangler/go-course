Exercise: Working with Maps in Go

Instructions:
1. Create a new Go program on your machine.
2. Implement the following exercise by following the instructions.
3. Test your program to verify its correctness.
4. If you get stuck or need assistance, refer to the provided solution and explanation.

Exercise Description:
Write a program that keeps track of the number of occurrences of each word in a given sentence using a map. The program should prompt the user to enter a sentence and then display the count of each word in the sentence.

Example:
Input: "I love programming in Go. Go is a powerful language."
Output:
- I: 1
- love: 1
- programming: 1
- in: 1
- Go: 2
- is: 1
- a: 1
- powerful: 1
- language: 1

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
	// Prompt the user to enter a sentence
	fmt.Print("Enter a sentence: ")

	// Read the sentence from standard input
	reader := bufio.NewReader(os.Stdin)
	sentence, _ := reader.ReadString('\n')

	// Remove leading/trailing whitespaces and convert to lowercase
	sentence = strings.TrimSpace(strings.ToLower(sentence))

	// Split the sentence into words
	words := strings.Split(sentence, " ")

	// Create a map to store word counts
	wordCount := make(map[string]int)

	// Iterate over the words and update the word counts
	for _, word := range words {
		wordCount[word]++
	}

	// Display the count of each word
	fmt.Println("Word Counts:")
	for word, count := range wordCount {
		fmt.Printf("- %s: %d\n", word, count)
	}
}
```

Explanation:
1. We prompt the user to enter a sentence using `fmt.Print` and read the input using `bufio.NewReader` and `ReadString`.
2. The input sentence is stored in the `sentence` variable, and we use `strings.TrimSpace` to remove any leading/trailing whitespaces.
3. The sentence is converted to lowercase using `strings.ToLower`.
4. We split the sentence into individual words using `strings.Split` and store them in the `words` slice.
5. A map named `wordCount` is created using `make` to store the word counts.
6. We iterate over each word in the `words` slice using a `for` loop. For each word, we increment its count in the `wordCount` map using `wordCount[word]++`.
7. Finally, we display the count of each word by iterating over the `wordCount` map using a `for` loop and printing the word and its count.

Testing:
1. Build the program using `go build` and run it.
2. Enter a sentence, such as "I love programming in Go. Go is a powerful language."
3. Verify that the program displays the correct count of each word as shown in the example output.

Note: Encourage students to experiment with different sentences to test the program's functionality and explore the behavior of maps in Go.