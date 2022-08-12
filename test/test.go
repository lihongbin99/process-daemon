package main

import (
	"flag"
	"fmt"
	"net"
)

func init() {
	a := flag.String("a", "", "test")
	b := flag.String("b", "", "test")
	flag.Parse()
	fmt.Println("test", *a)
	fmt.Println("test", *b)
}

func main() {
	//addr, err := net.ResolveTCPAddr("tcp", ":51234")
	addr, err := net.ResolveTCPAddr("tcp", ":54321")
	if err != nil {
		panic(err)
	}

	listener, err := net.ListenTCP("tcp", addr)
	if err != nil {
		panic(err)
	}

	for {
		_, err = listener.AcceptTCP()
		if err != nil {
			panic(err)
		}
	}
}
