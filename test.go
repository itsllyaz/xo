package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"xo/internal/httphandler"
)

func main() {
	url := "http://cs.bell-labs.co/who/ken/"
	httpRequest := httphandler.CreateGetRequest(url)
	res, err := http.DefaultClient.Do(httpRequest)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	fmt.Println(string(body))

}
