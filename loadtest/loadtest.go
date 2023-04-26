package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	endpoint := "http://127.0.0.1:3000/check"
	interval := time.Second

	for {
		res, err := http.Get(endpoint)
		if err != nil {
			fmt.Println("Error:", err)
		}
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println("Error:", err)
		}
		fmt.Println("Response Body:", string(body))
		time.Sleep(interval)
	}
}
