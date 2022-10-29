package main

import "fmt"

func main() {
	fixed := [...]int{0, 1, 2} //fixed Array
	fmt.Println(fixed)
	fmt.Printf("%T", fixed) //[3]int
	fmt.Println()
	// fixed=[...]int{0,1} //we can't reassign the size of array

	//so we use slices
	a := []int{0, 1}    //slices
	fmt.Println(a)      //[0 1]
	fmt.Printf("%T", a) //[]int
	fmt.Println()
	fmt.Println(len(a)) //2 //len in slice/map ------>number of elements in slice_variable, if slice_variable is nil then len(slic_variable) is zero

	//reassign size differnt size
	a = []int{0, 1, 2}
	fmt.Println(a)      //[0 1 2 ]
	fmt.Printf("%T", a) //[]int
	fmt.Println()
	fmt.Println(len(a)) //3

	//append more elements

	a = append(a, 3, 4, 5, 6)
	fmt.Println(a)
	fmt.Println(cap(a), len(a)) //len---->size of the how many elements there are
	//slice---->the maximum length the slice can reach when resliced;

	//Declare Slices using the make function
	//Make function is Builtin function it is used to declare and intilize an object
	// Slice: The size specifies the length. The capacity of the slice is
	// equal to its length. A second integer argument may be provided to
	// specify a different capacity; it must be no smaller than the
	// length. For example, make([]int, 0, 10) allocates an underlying array
	// of size 10 and returns a slice of length 0 and capacity 10 that is
	// backed by this underlying array.
	b := make([]int, 5)         //we have not provided any values so by default it assign values in'0'
	fmt.Println(b)              //[0 0 0 0 0]
	fmt.Println(len(b), cap(b)) //5 5

	//Slice of a Slice

	fmt.Println(a) //[0 1 2 3 4 5 6] ---> 7 elelments in slice
	// a = a
	a = a[0:7]
	fmt.Println(a)      //[0 1 2 3 4 5 6]
	a = a[0:8]          //here we have 7 elements only this actually append one extra element
	fmt.Println(a)      //[0 1 2 3 4 5 6 0]
	fmt.Println(len(a)) //8
	fmt.Println(cap(a)) //8
	a = a[0:7]          //
	fmt.Println(a)      //[0 1 2 3 4 5 6] here it actually sliced one end  element
	a = a[0:]
	fmt.Println(a) //[0 1 2 3 4 5 6]
	a = a[:7]
	fmt.Println(a) //[0 1 2 3 4 5 6]

	//Comparision Between Slices and nil

	if a == nil {

	}
	var c []int
	fmt.Println(c == nil)

}
