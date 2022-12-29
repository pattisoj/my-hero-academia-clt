package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("https://myheroacademia-api.onrender.com/characters/5")
	if err != nil {
	   log.Fatalln(err)
	}
 
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
	   log.Fatalln(err)
	}
 
	sb := string(body)
	log.Printf(sb)
 }
