package main

import (
	Client "NetworkProgramming/client"
	Server "NetworkProgramming/server"
	"fmt"
)

func main() {
	var (
		choice   int
		port     string = "4545"
		ip       string = "127.0.0.1"
		protocol string = "tcp"
	)
	fmt.Print("Please enter mode\n", "1: Send request from client;\n", "2: Launch server;\n", "Mode: ")
	fmt.Scan(&choice) //fmt.Fscanln(os.Stdin, &choice)
	switch choice {
	case 1:
		{
			Client.SendRequest(protocol, ip, port)
		}
	case 2:
		{
			Server.Launch(protocol, port)
		}
	}
}
