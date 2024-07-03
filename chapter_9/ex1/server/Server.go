package server

import (
	"fmt"
	"net"
)

func Launch(protocol, port string) {
	var message string = "Welcome back Commander!"
	listener, err := net.Listen(protocol, ":"+port)

	if err != nil {
		fmt.Println("LAUNCH ERROR!", err)
		return
	}
	defer listener.Close()
	fmt.Println("Server is listening...")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("LISTEN ERROR!", err)
			return
		}
		conn.Write([]byte(message))
		conn.Close()
	}
}
