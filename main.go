package main

import (
	"io/ioutil"
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get("https://myheroacademia-api.onrender.com/characters/5")
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
