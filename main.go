package main

import (
	"fmt"

	"github.com/haroonalbar/jwtgo/initializers"
)

// this is executed before main
// init is mainly used to setup the project
func init() {
	initializers.LoadEnvVariables()
}

func main() {
	fmt.Println("JWT")
}
