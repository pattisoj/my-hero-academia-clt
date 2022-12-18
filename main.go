package main

import (
	"fmt"
)

func main() {
	i := 7
	inc(&i)
	fmt.Println(i)
}

func inc(x *int) {
	*x++
}

/* This example is looking at pointers. Without the * on line 13 (before the int)
and without the & on line 9 this function on line thirteen receives a copy so running the file
would return 7 - as it doesn't edit the original due to lack of return interaction. 
By pointing to it's place in memory (with the * and &) we are able to make 'i' = 8. HOWEVER!
It is important to note the additional * on line 14. Without this we increment the memory address 
it is important that we de reference the pointer in the func */