//ARRAYS

package main

import "fmt"

func main() {
	//var name [size]type
	var a [3]int
	fmt.Println(a)      //[0 0 0]
	fmt.Println(a[0])   //0
	fmt.Println(len(a)) //no.of elements

	//reassign element in the array
	a[0] = 1
	fmt.Println(a) //[1 0 0]
	//arrays are values instead of address in golang

	aCopy := a // in java we get address but in Golang we get value
	fmt.Println(aCopy == a)
	fmt.Println(aCopy)

	a[0] = 2
	fmt.Println(a)     //[2 0 0]
	fmt.Println(aCopy) //[1 0 0] from these example we can understand after updating a[0]=2 element .still in aCopy variable we are getting the old value bcoz in Golang array are stored as values not by addresses
	fmt.Println(aCopy == a)
	//name := [size]type{values}
	b := [4]int{1, 2, 3, 4}
	fmt.Println(b) //[1 2 3 4]

	c := [...]int{1, 2} //here we specified '...' it automatically takes the size
	fmt.Println(c)

	d := [3]int{1, 2}
	fmt.Println(d)

	// d := [4]int{1, 2} //we can't reassign the array size
	// fmt.Println(d)

	//compare ARRAYS
	fmt.Println(a == d) //false
	fmt.Println()
	//tip:- if we don't want to declare 'i' simply we can use '_'
	for _, v := range b {
		fmt.Println(v)
	}
	fmt.Println()

	array := [...]int{0: 1} //a[0]=1
	fmt.Println(array)

	array2 := [...]int{1: 1, 4: 5} //here we assigned a[1]=1 ; a[4]=5
	fmt.Println(array2)

	//String Array

	strArray := [...]string{"string1", "string2"}
	fmt.Println(strArray)

	//Quiz:- do the above excersices with the string

	//Multi-diemenstional Array

	array2d := [2][2]int{{1, 2}, {3, 4}}
	fmt.Println(array2d)

	array3d := [2][2][2]int{array2d, array2d}
	fmt.Println(array3d)

}
