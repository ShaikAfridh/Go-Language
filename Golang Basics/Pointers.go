package main

import "fmt"

func main() {

	//***********POINTERS***********
	var x int = 1
	var p *int = &x // &---------> address of variable {we stored address of x in p}

	fmt.Println(p)
	fmt.Println(*p) //it will denote the address of value of x

	/*Output:-
	0xc000016078 -------> memory address of variable 'x'
	1 -------> it shows value of address stored in 'p'
	*/

}
