package main

import "./blockchain"

func main() {
	b0 := blockchain.NewBlock("Hello Crypto World Here is my first block", "0000000000000000000000000000000000000000000000000000000000000000")
	blockChain := []blockchain.Block{b0}
}
