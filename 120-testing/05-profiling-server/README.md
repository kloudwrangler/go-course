# HTTP Server with Profiling

Look through the program `profiling.go`. It is instrumented with profiling which attaches to the server that it starts.

Run the program

```shell
go run profiling.go
```

A server is listening on localhost:8080. It is now exposing debug information on path `/debug/pprof`. Go to your browser to see what it is automatically exposing at http://localhost:8080/debug/pprof/. This view is legacy text mode. 

If you would like to download the binaries instead, you can download them via curl e.g.


```shell
# Download the heap's profile
curl -s  http://localhost:8080/debug/pprof/heap > heap.out
# Parse it with pprof
go tool pprof heap.out
```
You can also use it directly

```shell
go tool pprof  http://localhost:6060/debug/pprof/heap
```

You can use the pprof tool just like normal. Try `top`, `web`, `png`, etc.