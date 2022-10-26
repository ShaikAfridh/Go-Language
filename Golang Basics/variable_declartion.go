/*There are 4 types of  Declaration
1. Variable
2. Type
3. Constant
4. Function
*/

package main

import "fmt"

func main() {
	// fmt.Println("Hello World")
	//Variable Declaration
	//Syntax:
	//var name type = expression
	// var integer int = 1
	// var integer int //if we haven't assign integer value by default it will assign '0' value
	/*integer------> '0'
	  boolean------> 'false'
	  string-------> 'empty string'
	  interfaces/references-------> null*/

	// fmt.Println(integer)
	//take multiple variables
	// var integer1, integer2 int
	// fmt.Println(integer1, integer2)

	//removing type and giving expression/value in variable

	// var integer1, integer2 = 1, 2
	// fmt.Println(integer1, integer2)

	//we can declare different types of integer in same line

	var integer, string = 1, "AFRIDI"
	fmt.Println(integer, string)

	//Short Declaration of variables
	// Syntax:-
	// name := expression

	boolean := false
	//Excercise:-change the above boolean value to 'true' and run again
	fmt.Println(boolean)

}
