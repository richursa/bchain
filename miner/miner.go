package main

import (
	"fmt"

	"../blockchain"
)

func main() {
	peerlist := []string{"localhost"}
	var mychain blockchain.Blockchain
	fmt.Println("current blockchain is ", mychain)
	fmt.Println("requesting blocks from peers")
	mychain = append(mychain, mychain.RequestLatestBlock(peerlist))
	fmt.Println("current blockchain is ", mychain)
}
