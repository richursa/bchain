package main

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"strings"
	"time"
)

type Block struct {
	Time  int64  // to store unix time
	Data  string // transactions/data which is to be stored in a block
	Prev  string //	hash of the previous block
	Hash  string // hash of the current block
	Nonce int64  // the nonce which is added to block header to produce a valid hash
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
func (b Block) BlockToString() string {
	return strconv.FormatInt(b.Time, 10) + "," + b.Data + "," + b.Hash + "," + b.Prev + "," + strconv.FormatInt(b.Nonce, 10)
}
func stringToBlock(s string) (Block, error) {
	splittedBlock := strings.Split(s, ",")
	Time, err := strconv.ParseInt(splittedBlock[0], 10, 64)
	if err != nil {
		return Block{}, err
	}
	Nonce, err := strconv.ParseInt(splittedBlock[4], 10, 64)
	if err != nil {
		return Block{}, err
	}
	return Block{Time, splittedBlock[1], splittedBlock[2], splittedBlock[3], Nonce}, nil
}
