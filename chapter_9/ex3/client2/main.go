package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://google.com", nil)
	req.Header.Add("Accept", "text/html")
	req.Header.Add("User-Agent", "MSIE/15.0")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("SEND ERROR!", err)
		return
	}
	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)
}
