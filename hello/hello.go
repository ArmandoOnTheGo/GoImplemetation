package main 	//Package Name

import (
	"fmt"	//If you need more than one import, you'll need ()
	"example.com/user/hello/morestrings"	
)
func main() {
	s := "Hello World"
	fmt.Printf(s+"\n")
	//Using the reverse string from morestrigns package
	fmt.Println(morestrings.ReverseRunes(s))
}
