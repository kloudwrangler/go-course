

#### Defer
- The `defer` statement is used to schedule a function call to be executed when the surrounding function completes.
- Deferred functions are executed in a last-in-first-out (LIFO) order.
- Commonly used for tasks like closing files, releasing resources, or logging.



```go
func main() {
    defer fmt.Println("World")
    fmt.Println("Hello")
}
```

Output:
```
Hello
World
```


#### Panic
- `panic` is a built-in function used to interrupt the normal execution flow of a program and start panicking.
- When a function encounters a panic, it immediately stops execution and begins to unwind the stack, executing any deferred functions along the way.
- Often used to handle exceptional or unrecoverable situations.



```go
func main() {
    fmt.Println("Start")
    panic("Something went wrong!")
    fmt.Println("End")
}
```
Output:
```
Start
panic: Something went wrong!
```


#### Recover
- `recover` is a built-in function used to regain control of a panicking goroutine.
- It is only useful when called inside a deferred function.
- Stops the panic and returns the value passed to the `panic` call.
- Allows graceful handling of panics and resuming normal execution.



```go
func main() {
    defer func() {
        if err := recover(); err != nil {
            fmt.Println("Recovered:", err)
        }
    }()
    panic("Something went wrong!")
    fmt.Println("This line will not be executed")
}
```
Output:
```
Recovered: Something went wrong!
```


#### Defer, Panic, Recover: Best Practices
- Use `defer` to ensure important cleanup actions are executed.
- Reserve `panic` for exceptional and unrecoverable situations.
- Gracefully handle panics using `recover` to resume normal execution when possible.
- Avoid using panics for normal error handling scenarios.
