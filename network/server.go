package network

import (
	"fmt"
	"log"
	"net"
)

func StartServer(addr string) {
	l, err := net.Listen("tcp4", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
			return
		}

		go handleRequest(conn)
	}
}

// handleRequest 处理请求
func handleRequest(conn net.Conn) {
	defer conn.Close()

	for {
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print(string(buf[:n]))
	}
}
