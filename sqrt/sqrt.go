package main	  //Start of main

import (
	"fmt"
	"math"	  //importing math to use Abs "absolute value"
	"bufio"   //import for Reader
	"os"	  //import to ready user input
	"strconv" //import to convert user input to usable value
	"strings"
)

//Function to find "square root" of passed value using Newton's Method
//Pre: float value
//Post: estimated value of sqrt root
func Sqrt(x float64) float64 { 		
	z := x/4				//Choosing arbritary starting point
	for i, count:= x-z, 0;  i > .001; {	//Start of for loop
	
	//fmt.Println("Current z value: ", z, "count #",count)	//Displaying the current values of z and count
	if power := math.Pow(z, 2); power == x { //Checking if the current value of z squared is = x
		return z 
	} 
		
	if count == 10{		//If we've reached 10 iterations we end and return z
		return z
	}

	temp := z		//set temp to z (to check if we're making any progress or stagnating)
	z -= (z*z - x) / (2*z) 	//Updating z following Newton's method
	i = math.Abs(z-temp)	//Updating i based on the new values of previous z and current z
	count++			//Incrementing count
	}

	return z		//if we reach stagnantion on the number, exit for loop and return z
}

func main() {
	reader := bufio.NewReader(os.Stdin)
        
	fmt.Println("Enter Value (can include decimals) to find approx square root")
        text, _ := reader.ReadString('\n')
        text = strings.Replace(text, "\n","",-1)
	validInput, err := strconv.ParseFloat(text, 64);
        
	for err != nil{
                fmt.Println("Ever a valid number (can include decimals)")
	      	text, _ = reader.ReadString('\n')
	        text = strings.Replace(text, "\n","",-1)
		validInput, err = strconv.ParseFloat(text, 64);
        }
	
	fmt.Println("Approx Square Root of ",validInput," is ",Sqrt(validInput));
	//fmt.Println(Sqrt(25)) Used for Testing
}
