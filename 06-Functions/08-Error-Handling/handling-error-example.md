

Code Example:
```go
package main

import (
	"fmt"
	"os"
)

func main() {
	filePath := "path/to/nonexistent/file.txt"

	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		// Handle the error
		fmt.Printf("Error: Failed to open the file: %s\n", err.Error())
		return
	}
	defer file.Close()

	// Read the file content
	content := make([]byte, 1024)
	n, err := file.Read(content)
	if err != nil {
		// Handle the error
		fmt.Printf("Error: Failed to read the file: %s\n", err.Error())
		return
	}

	fmt.Printf("File Content (%d bytes):\n%s\n", n, content[:n])
}
```

Explanation:
1. In the `main` function, we define a file path `filePath` that points to a nonexistent file.
2. We use `os.Open` to open the file. If an error occurs during the file opening, the `err` variable will hold the error value.
3. We check if the `err` variable is not `nil` to determine if an error occurred. If so, we print an error message and return.
4. If the file is successfully opened, we defer the file's closure using `defer file.Close()`. This ensures that the file is properly closed before exiting the function.
5. We create a byte slice `content` to store the file's content and use the `file.Read` function to read the content into it. The number of bytes read and any error encountered are returned.
6. We again check if an error occurred during the file reading. If so, we print an error message and return.
7. Finally, we print the file content on the console.

In this example, we catch and handle errors by checking the `err` variable after each potentially error-prone operation. If an error occurs, we provide meaningful error messages and take appropriate actions.