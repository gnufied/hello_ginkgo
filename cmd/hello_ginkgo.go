package main

import (
	"fmt"
	"github.com/gnufied/hello_ginkgo/pkg/greeter"
)

func main() {
	message := greeter.SayHello()
	fmt.Println("Hello World")
}
