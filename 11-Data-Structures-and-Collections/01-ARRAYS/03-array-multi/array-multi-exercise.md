# Exercise: Manipulating Multidimensional Arrays

Instructions:
1. Create a new Go program.
2. Declare a two-dimensional array named "scores" with dimensions 3x4 to store the scores of students in a class.
3. Initialize the array with random integer values ranging from 0 to 100.
4. Write a function called "displayScores" that accepts the "scores" array as a parameter and displays the scores in a tabular format.
   - Use formatting to align the scores neatly.
5. Invoke the "displayScores" function in the main function to display the initial scores.
6. Write a function called "updateScore" that accepts the "scores" array, row index, and column index as parameters.
   - Prompt the user to enter a new score for the specified student and assignment.
   - Update the corresponding element in the array with the new score.
7. In the main function, prompt the user to choose a student and assignment to update.
   - Accept input for the row index and column index.
   - Call the "updateScore" function with the chosen indices and the "scores" array.
8. Invoke the "displayScores" function again to display the updated scores.
9. Compile and run the program.
10. Test different input values to verify the functionality.

Example Output:
```
Initial Scores:
-----------------------
| 90 | 80 | 75 | 92 |
| 85 | 94 | 88 | 79 |
| 70 | 68 | 91 | 86 |
-----------------------

Choose a student (0-2): 1
Choose an assignment (0-3): 2
Enter a new score: 92

Updated Scores:
-----------------------
| 90 | 80 | 75 | 92 |
| 85 | 94 | 92 | 79 |
| 70 | 68 | 91 | 86 |
-----------------------
```

Note: You may utilize Go's `fmt` package for input/output and formatting purposes.

Happy coding!