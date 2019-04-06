package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type customWriter struct{}

func main() {

	resp, err := http.Get("http://www.wykop.pl")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// bs := make([]byte, 999999)
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))

	cw := customWriter{}

	// io.Copy(os.Stdout, resp.Body)
	io.Copy(cw, resp.Body)
}

func (customWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Read:", len(bs))
	return len(bs), nil
}
