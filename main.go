package main

import (
	"fmt"
	"masala/packets"
	"masala/state_manager"
	"net"
)

func main() {
	// Initialize state
	state_manager.GetState()

	tcpAddr := net.TCPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 3001}
	lis, err := net.ListenTCP("tcp", &tcpAddr)
	fmt.Println("Masala listening at 127.0.0.1:3001...")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := lis.AcceptTCP()

		if err != nil {
			panic(err)
		}
		buffer := make([]byte, 256)
		n, err := conn.Read(buffer)
		if err != nil {
			panic(err)
		}

		fmt.Println("Received:", n)

		linkMess, linkCaps := packets.ParseSpiceLinkMess(buffer)
		packets.NewSpiceLinkReply()

		fmt.Println(linkMess, linkCaps)

	}
}
