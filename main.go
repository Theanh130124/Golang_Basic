package main

import (
	"fmt"
	"os"
)

func main() {

	fmt.Println("Hello World" + os.Getenv("USERNAME"))

}
