package main 	//Package Name

import (	//If you need more than one import, you'll need ()
	"fmt"	
	"example.com/user/hello/morestrings"	//Importing from a local different package
	"github.com/google/go-cmp/cmp"		//Importing from a remote package
)

func main() {
	s := "Hello World"				//setting s to String "Hello World"
	fmt.Printf(s+"\n")				//outputting s NOTE: similar to java
	fmt.Println(morestrings.ReverseRunes(s))	//Using the reverse string from morestrigns package
	fmt.Println(cmp.Diff("Hello World", "Hello Go"));
}

