package main

import (
	"fmt"
	"net"
)

var dict = map[string]string{
	"red":    "красный",
	"yellow": "жёлтый",
	"green":  "зелёный",
	"white":  "белый",
	"black":  "чёрный",
}

func main() {
	listener, err := net.Listen("tcp", ":4545")
	if err != nil {
		fmt.Println("LAUNCH ERROR!", err)
	}
	defer listener.Close()
	fmt.Println("Server is listening")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("CONNECT ERROR!", err)
			conn.Close()
			continue
		}
		go hangleConnection(conn)
	}
}

func hangleConnection(conn net.Conn) {
	defer conn.Close()
	for {
		input := make([]byte, (1024 * 4))
		n, err := conn.Read(input)
		if n == 0 || err != nil {
			fmt.Println("READ ERROR!", err)
			break
		}
		source := string(input[0:n])
		target, ok := dict[source]
		if ok == false {
			target = "undefined"
		}
		fmt.Println(source, " - ", target)
		conn.Write([]byte(target))
	}
}
