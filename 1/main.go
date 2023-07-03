package main

import (
	"flag"
	"fmt"
)

func main() {
	addr := flag.String("addr", "localhost:5051", "the address to connect to")
	fmt.Println(*addr)

	fmt.Println("Hello World!")
	f := foo{}
	f.can()

	// ctx, _ := context.WithTimeout()
}

type foo struct {
	v   int
	str string
}

func (f foo) can() {
	fmt.Println("hahah")

}
IIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIII()