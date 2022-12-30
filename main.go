package main

import (
	"io/ioutil"
	"fmt"
	"net/http"
)

func main() {
	id := 5
	url := fmt.Sprintf("https://myheroacademia-api.onrender.com/characters/%d", id)

	resp, err := http.Get(url)
	if err != nil {
	   fmt.Print(err)
	}
 
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err)
	}
 
	sb := string(body)
	fmt.Print(sb)
 }
