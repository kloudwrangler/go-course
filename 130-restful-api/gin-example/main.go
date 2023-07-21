package main

import (
	"gin-example/routes"
)

func main() {
	// Our server will live in the routes package
	routes.Run()
}
