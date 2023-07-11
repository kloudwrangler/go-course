# Making a module and importing 3rd party projects

You are told to create a script to display some data. 

```shell
 go run cmd/grapher/main.go
```

```txt
[3 4 9 6 2 4 5 8 5 10 2 7 2 5 6]
```

It works great to display the numbers, however you are told that it needs to display a graph. You find a great 3rd party module that already does this at http://github.com/guptarohit/asciigraph.

## Instructions

First, run the code locally.

```go
go run cmd/grapher/main.go
```
This will simply print the numbers.


Next, move the code to a place on your machine that is not this project (this is because this project is already a module and we are trying to practice modules 🤪).

Modify the code so that it uses the `github.com/guptarohit/asciigraph` project to make a graph of your numbers.

Try to runing it again. It will complain about it now.

You will first have to create a module (`go mod init <name-of-module>`), and download the dependencies (`go get <dep>` or `go mod tidy` ), then create the program, and make sure that it runs (`go run`).

Ensure that you are using the appropriate project structure. Typically, you want command line tools to go in `/cmd/<name-of-binary>`.

Before compiling the code, run the go formater through it with `go fmt <path-to-go-files>`.

Next, compile the binary using `go build`.

Check to ensure that the binary runs (`./binary-name`)
