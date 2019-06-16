package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net"
	"strconv"
	"strings"
	"time"
)

// Block : the basic data structure of a blockchain
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
		}
		nonce++
	}
}

// NewBlock : returns a new block with the hash of required difficulty
func NewBlock(data string, prev string) Block {
	t := time.Now().Unix()
	difficulty := "00000"
	nonce, hash := computeHashWithProofOfWork(intToStr(t)+prev+data, difficulty)
	return Block{t, data, prev, hash, nonce}
}

// BlockToString : converts the block to a string with values separeted by a coma
func (b Block) blockToString() string {
	return strconv.FormatInt(b.Time, 10) + "," + b.Data + "," + b.Hash + "," + b.Prev + "," + strconv.FormatInt(b.Nonce, 10)
}

// StringToBlock : converts a string, with values separeted by coma to Block
func StringToBlock(s string) (Block, error) {
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

// Blockchain :  a slice of blocks
type Blockchain []Block

// RequestLatestBlock : request a new block from a list of peers
func (b Blockchain) RequestLatestBlock(peerlist []string) Block {
	length := len(b)
	c := make(chan string)
	go func() {
		for _, peer := range peerlist {
			go func() {
				conn, err := net.Dial("tcp", peer+":8888")
				if err != nil {
					fmt.Println(err)
					return
				}
				data := make([]byte, 100000)
				data = []byte("requestLatestBlock" + "," + intToStr(int64(length)))
				conn.Write(data)
				data, err = ioutil.ReadAll(conn)
				fmt.Println("received data = ", string(data))
				if err != nil {
					fmt.Println(err)
					return
				}
				c <- string(data)
			}()
		}
	}()
	select {
	case recievedBlock := <-c:
		{
			latestBlock, err := StringToBlock(recievedBlock)
			if err != nil {
				fmt.Println(err)
			}
			return latestBlock
		}
	}
	// TO-DO :: implement default and timeout
}
func (b Blockchain) ServeBlock() {
	ln, err := net.Listen("tcp", "localhost:8888")
	if err != nil {
		fmt.Println(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
		}
		go func() {
			tmp := make([]byte, 10000)
			conn.Read(tmp)
			blockID, err := strconv.ParseInt(strings.Split(string(tmp), ",")[1], 10, 64)
			fmt.Println("blockid =", blockID)
			if err != nil {
				//fmt.Println(err)
			}
			if (int64(len(b))) >= blockID {
				s := b[blockID].blockToString()
				conn.Write([]byte(s))
				conn.Close()
			} else {
				conn.Close()
			}
		}()
	}

}
