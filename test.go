package main

import "fmt"

type Block struct {
	a int
	b int
	c string
}

func main() {
	var b Block
	b = bin()
	fmt.Println(b)
}
func bin() Block {
	return Block{}
}
