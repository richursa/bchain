package main

import (
	"fmt"
	"time"

	"../blockchain"
)

func main() {
	peerlist := []string{"localhost"}
	var mychain blockchain.Blockchain
	fmt.Println("current blockchain is ", mychain)
	fmt.Println("requesting blocks from peers")
	mychain.RequestLatestBlock(peerlist)
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("current blockchain is ", mychain)
}
