package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("アドレスとポートを指定してください。")
		return
	}

	ip := os.Args[1]
	port := os.Args[2]

	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%s", ip, port))
	if err != nil {
		fmt.Printf("TCP port %s is closed.", port)
		return
	}
	defer conn.Close()
	fmt.Printf("TCP port %s is opened.", port)
}
