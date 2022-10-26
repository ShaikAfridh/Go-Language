package main

import "fmt"

func main() {
	//***********TYPE*************

	// fahrenheit & celsius

	type fahrenheit int
	type celsius int

	var f fahrenheit = 32
	var c celsius = 0
	// c=f //here we can't define value like this for this we have declared farhenheit,Celsius ------> Empty and we can assign the value

	c = celsius((f - 32) * 5 / 9) //Conversion farhen heit to celsius

	fmt.Println(f, c)

}
