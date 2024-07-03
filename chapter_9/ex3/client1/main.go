package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	client := http.Client{
		Timeout: 6 * time.Second,
	}
	resp, err := client.Get("https://google.com")
	if err != nil {
		fmt.Println("CONNECTION ERROR!", err)
		return
	}
	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)
}
