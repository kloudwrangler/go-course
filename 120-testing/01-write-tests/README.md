# Testing Exercises

Here's an example program that demonstrates the basics of testing, assertions using the stretch/testify package, and profiling in Go. It's a simple calculator program with addition and multiplication operations.

```go
// calculator.go

package calculator

func Add(a, b int) int {
	return a + b
}

func Multiply(a, b int) int {
	return a * b
}
```

`calculator_test.go` has corresponding unit tests using the testing package and assertions with stretch/testify.

```go
// calculator_test.go

package calculator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	result := Add(2, 3)
	assert.Equal(t, 5, result, "Addition failed")
}

func TestMultiply(t *testing.T) {
	result := Multiply(2, 3)
	assert.Equal(t, 6, result, "Multiplication failed")
}

func TestMultWithoutAssert(t *testing.T) {
	got := Multiply(2, 3)
	if got != 6 {
		t.Errorf("Multiply(2,3) = %d; want 6", got)
	}
}
```
## Part 1. Run tests
1. Run the tests using the `go test` command:
   
   ```
   go test
   ```

   You should see the test results, indicating whether the tests passed or failed.

## Part2. See a failure for both stretchr/testify and testing package
1. Make `TestMultiply` and `TestMultWithoutAssert` fail by changing the expected number.
2. Run the tests
3. Notice the difference in the output. Also note how expressive the tests are writtten with

# Part 3. Writing Tests with stretchr/testify

1. In `calculator.go` add a function for subtraction
1. In the `calculator_test.go` add a function to test the new function
2. Run the tests using the `go test` command

# Part 3. Writing Tests with testing package

1. In the `calculator_test.go` add a function to test the new subtraction function but only using whats available in the standard library
2. Run the tests using the `go test` command


## Part 4. Profiling
With the program and tests in place, let's walk through the steps to run the tests and perform profiling.


1. Now, let's perform CPU profiling. Run the tests with the `-cpuprofile` flag to generate a CPU profiling file:
   ```
   go test -cpuprofile=cpu.prof
   ```

2. Next, let's analyze the CPU profiling data using the `go tool pprof` command:
   ```
   go tool pprof cpu.prof
   ```

   This opens an interactive shell. You can type `top` to see the top functions consuming CPU time.

3. To perform memory profiling, run the tests with the `-memprofile` flag to generate a memory profiling file:
   ```
   go test -memprofile=mem.prof
   ```

4. Analyze the memory profiling data using the `go tool pprof` command:
   ```
   go tool pprof mem.prof
   ```

   Similar to CPU profiling, this opens an interactive shell where you can explore memory allocations and usage.

That's it! You have now walked through a working Go program with tests and learned how to run the tests and perform profiling. Feel free to explore further and experiment with different scenarios and profiling options.