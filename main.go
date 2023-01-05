package main

import (
	"io/ioutil"
	"fmt"
	"net/http"
	"math/rand"
	"time"
	"strings"
)

func main() {
	min := 1
  	max := 133

	// NOTE: I had to set a seed here to ensure a random number is generated each time
	rand.Seed(time.Now().UnixNano())

	id := rand.Intn(max - min) + min
	url := fmt.Sprintf("https://myheroacademia-api.onrender.com/characters/%d", id)

	resp, err := http.Get(url)
	if err != nil {
	   fmt.Print(err)
	}
 
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err)
	}
 
	stringBody := string(body)
	returnString := stringManipulation(stringBody)

	fmt.Print(returnString)
 }

 func stringManipulation(stringBody string)[]string{
	replacer := strings.NewReplacer(
		"{", "",
		"}", "",
		"\"", "",
	)
	cleanString := replacer.Replace(stringBody)
	
	splitString := strings.Split(cleanString, ",")

	return splitString
 }