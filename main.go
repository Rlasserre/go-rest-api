package main

import (
	"fmt"
	"go-rest-api/routes"
)

func main() {
	fmt.Println("Initializing Go server...")
	routes.HandleRequest()
}
