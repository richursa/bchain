package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Block struct {
	Time  int64
	Data  string
	Prev  string
	Hash  string
	Nonce int64
}

func binaryToString(bytes []byte) string {
	return hex.EncodeToString(bytes)
}

func intToStr(num int64) string {
	return strconv.FormatInt(num, 10)
}
func calcHash(data string) string {
	hashed := sha256.Sum256([]byte(data))
	return binaryToString(hashed[:])
}
func computeHashWithProofOfWork(data string, difficulty string) (int64, string) {
	nonce := int64(0)
	for {
		hash := calcHash(intToStr(nonce) + data)
		if strings.HasPrefix(hash, difficulty) {
			return nonce, hash
		} else {
			nonce++
		}
	}
}
func NewBlock(data string, prev string) Block {
	t := time.Now().Unix()
	difficulty := "000000"
	nonce, hash := computeHashWithProofOfWork(intToStr(t)+prev+data, difficulty)
	return Block{t, data, prev, hash, nonce}
}
func stringToBlock(s string) Block {

}
func main() {
	b0 := NewBlock("Hello Crypto World Here is my first coin", "0000000000000000000000000000000000000000000000000000000000000000")
	b1 := NewBlock("Hello Cryptos ", b0.Hash)
	fmt.Println(b0)
	fmt.Println(b1)
	fmt.Println("\n\nBlockchian\n\n")
	blockChain := []Block{b0, b1}
	blockChain = append(blockChain, NewBlock("sukhamnu doode", b1.Hash))
	fmt.Println(blockChain)
}
