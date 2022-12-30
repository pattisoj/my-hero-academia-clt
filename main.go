package main

import (
	"io/ioutil"
	"fmt"
	"net/http"
	"math/rand"
	"time"
)

func main() {
	min := 1
  	max := 133

	// NOTE: I had to set a seed here to ensure a random number is generated each time
	rand.Seed(time.Now().UnixNano())

	fmt.Println(rand.Intn(max - min) + min)

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
