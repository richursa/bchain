package main

import (
	"fmt"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", "localhost:80")
	if err != nil {
		fmt.Println(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
		}
		go func() {
			conn.Write([]byte("server aayit connect aayitund ketta"))
		}()
	}
}
