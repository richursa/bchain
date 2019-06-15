package main

import "../blockchain"

func main() {
	chain := blockchain.Block{}
	if len(chain) == 0 {
		chain.requestNewblock()
	}
}
