# Profiling example

Look through the program `profiling.go`. It is instrumented with profiling function calls to measure the performance of the program.

Then, run the program

```shell
go run profiling.go
```

After running the program, you will have two profiling files: `cpu.prof` for CPU profiling and `mem.prof` for memory profiling.

To analyze the CPU profiling file, run the following command:

```shell
go tool pprof cpu.prof
```

This command will open an interactive shell where you can analyze the CPU profiling data. You can use commands like `top`, `png`, and `web` to explore the profiling report.

To analyze the memory profiling file, run the following command:

```shell
go tool pprof mem.prof
```

This command will also open an interactive shell where you can analyze the memory profiling data. You can use commands like `top`, `png`, and `web` to explore the memory profiling report.

Please note that you need to have the Go toolchain installed and set up correctly to use the `go tool pprof` command. Additionally, ensure that the profiling files are in the same directory from where you run the commands.