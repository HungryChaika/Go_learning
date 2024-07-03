package client

import (
	"fmt"
	"io"
	"net"
	"os"
)

func SendRequest(protocol, ip, port string) {
	conn, err := net.Dial(protocol, ip+":"+port) //"tcp", "127.0.0.1:4545"
	if err != nil {
		fmt.Println("CONNECT ERROR!", err)
		return
	}
	defer conn.Close()

	io.Copy(os.Stdout, conn)
	fmt.Print("\nGOOD JOB!")
}
