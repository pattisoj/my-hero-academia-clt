package main

import (
	"io/ioutil"
	"fmt"
	"net/http"
	//"math/rand"
	//"time"
	"strings"
)

func main() {
	/*min := 1
  	max := 133

	// NOTE: I had to set a seed here to ensure a random number is generated each time
	rand.Seed(time.Now().UnixNano())

	id := rand.Intn(max - min) + min */
	id := 92
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
	
	fmt.Print(stringManipulation(stringBody, id))
 }

 func stringManipulation(stringBody string, id int)string{
	if id>45 && id<134 {
		if id>91 {
			replacer := strings.NewReplacer(
				"{", "",
				"}", "",
				"\"", "",
				":", "",
				"name", "",
				"hero_", "",
				"quirk", "",
			)
			cleanString := replacer.Replace(stringBody)
			splitString := strings.Split(cleanString, ",")
			output := fmt.Sprintf("You've randomly drawn a hero! %s, you may also know them as %s!", splitString[1], splitString[3])
			
			return output

		} else {

		}
	} else {
		replacer := strings.NewReplacer(
			"{", "",
			"}", "",
			"\"", "",
			":", "",
			"[", "",
			"]", "",
			"name", "",
			"other_names", "",
            "quirk", "",
		)
		cleanString := replacer.Replace(stringBody)
		splitString := strings.Split(cleanString, ",")
		output := fmt.Sprintf("You've randomly drawn a student! %s, you may also know them as %s!", splitString[1], splitString[3])
		
		return output
	}

	return ""
 }