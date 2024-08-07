package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:4545")
	if err != nil {
		fmt.Println("CONNECTION ERROR!", err)
		return
	}
	defer conn.Close()
	for {
		var source string
		fmt.Print("\nEnter word: ")
		_, err := fmt.Scan(&source)
		if err != nil {
			fmt.Println("ENTER ERROR!", err)
			continue
		}
		if n, err := conn.Write([]byte(source)); n == 0 || err != nil {
			fmt.Println("SEND ERROR!", err)
			return
		}
		fmt.Print("\nTranslate: ")
		conn.SetReadDeadline(time.Now().Add(time.Second * 5))
		for {
			buff := make([]byte, 1024)
			n, err := conn.Read(buff)
			if err != nil {
				break
			}
			fmt.Print(string(buff[0:n]))
			conn.SetReadDeadline(time.Now().Add(time.Millisecond * 700))
		}
		fmt.Println()
		/*
			buff := make([]byte, (1024 * 4))
			n, err := conn.Read(buff)
			if err != nil {
				break
			}
			fmt.Print(string(buff[0:n]))
			fmt.Println()
		*/
	}
}
