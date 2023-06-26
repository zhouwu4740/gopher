package network

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func Client(servAddr string) {
	conn, err := net.Dial("tcp4", servAddr)
	if err != nil {
		log.Fatal(err)
		return
	}

	defer conn.Close()

	stdin := bufio.NewReader(os.Stdin)
	for {
		line, err := stdin.ReadString('\n')
		if err != nil {
			log.Fatal(err)
			return
		}
		line = strings.Trim(line, "\r\n")

		if line == "exit" {
			fmt.Println("用户退出客户端")
			break
		}

		_, err = conn.Write([]byte("已收到\n"))
		if err != nil {
			log.Fatal(err)
		}
	}
}
