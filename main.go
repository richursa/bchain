package main

func main() {
	b0 := NewBlock("Hello Crypto World Here is my first block", "0000000000000000000000000000000000000000000000000000000000000000")
	blockChain := []Block{b0}
	s, _ := stringToBlock(blockChain[0].BlockToString())
}
