/*  SCOPE OF A VARIABLE
 scope is nothing but where we can access and use that variable

There are 3 Levels of scope variable
1.Block -------> between curly braces
2.Package
3.Universe */

//1. BLOCK
// package main

// import "fmt"

// func main() {

// 	var integer = 1
// 	fmt.Println(integer)

// 	{
// 		var diff = 2
// 		fmt.Println(diff)

// 	}
// 	// fmt.Println(diff) //we can't access this 'diff' variable from outside because it is in different scope
// 	//tip:- we can also put same variable name when the another variable is within the block
// 	{
// 		var integer = 5
// 		fmt.Println(integer) //here i'm using the same variable name but it is in differrent scope
// 	}
// }

//// 2. Package
//
//package main
//
//import (
//	"fmt"
//)
//
//var five = 5 //package level variable
//func main() {
//	fmt.Println(five)
//	fmt.Println(three)
//}
//
//func Nonmain() {
//	fmt.Println(five)
//}

//3.Universe
package main

import (
	"C:\\Users\\afridi\\GolandProjects\\awesomeProject\\Demo\\integers"
	"fmt"
)

func main() {
	fmt.Println(integers.Four) //Captilize the first character
}
