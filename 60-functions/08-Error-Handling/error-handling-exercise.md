Exercise: Error Handling in File Operations

Scenario:
You are building a command-line utility that reads a file path from the user and displays its content on the console. However, you need to handle potential errors that may occur during file operations.

Instructions:
1. Write a Go program that prompts the user to enter a file path.
2. Implement a function called `ReadFile` that takes a file path as a parameter and returns the file content as a string and an error if any.
3. In the `ReadFile` function, use the `os.ReadFile` function from the standard library to read the file content.
4. If an error occurs during file reading, return an error with a meaningful error message.
5. In the main function, call the `ReadFile` function with the user-provided file path and handle any errors that occur.
6. If an error occurs, print the error message and exit the program.
7. If the file is read successfully, print the file content on the console.

Example Output:

```
Enter the file path: /path/to/nonexistent/file.txt
Error: Failed to read the file: open /path/to/nonexistent/file.txt: no such file or directory

Enter the file path: /path/to/existing/file.txt
File Content:
This is the content of the file.

```

Note: The exercise focuses on error handling in file operations, and you can encourage participants to discuss the error scenarios and suggest improvements in error messages or error handling techniques.