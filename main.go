package main

import (
	"fmt"
	"time"

	"./blockchain"
)

func main() {
	var mychain blockchain.Blockchain
	mychain = append(mychain, blockchain.NewBlock("Hello Crypto World Here is my first block", "0000000000000000000000000000000000000000000000000000000000000000"))
	mychain = append(mychain, blockchain.NewBlock("this is the second block", mychain[0].Hash))
	fmt.Println("current blockchain is ", mychain)
	go mychain.ServeBlock()
	time.Sleep(100000 * time.Millisecond)
}
