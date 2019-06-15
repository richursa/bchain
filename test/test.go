package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:80")
	if err != nil {
		fmt.Println(err)
	}
	tmp := make([]byte, 10000)
	conn.Read(tmp)
	fmt.Println(string(tmp))
}
