package main

import (
	"fmt"
	"math"
)

func main() {
	// var x int = 2
	var x = 1 // int32 int64
	var x32 int32 = 1
	var x64 int64 = 1
	fmt.Println(x)
	fmt.Println(x32)
	fmt.Println(x64)

	fmt.Printf("%T,%T,%T", x, x32, x64) //here we can observe different types of integers in output
	fmt.Println()
	// x= x32 //we can't use 'x32' bcoz it is not type int
	// x= x64

	x = int(x32)
	fmt.Println(x)
	x = int(x64)
	fmt.Println(x)

	var unsigned uint = 1 //this is unsigned integer
	fmt.Println(unsigned)

	y := 2
	//Mathematical Operations
	fmt.Println(x + y)  //3
	fmt.Println(x % y)  //1
	fmt.Println(x == y) //false
	fmt.Println(x < y)  //true

	//floating point numbers

	pi := 3.141
	fmt.Println(pi)
	fmt.Printf("%T", pi) //this is will give the type of variable {By default it takes 'float64' }

	var pi32 float32 = 3.141
	fmt.Printf("%T", pi32) //this is 'float32'

	pi = float64(3.141) //this is called type conversion
	fmt.Println(pi)

	power := 1e100
	fmt.Println(power)

	//Conversion of integer &floating point

	a := 1
	b := 3.141
	fmt.Println(a)
	fmt.Println(b)
	a = int(b) //float(3.141) --------> integer(3)
	fmt.Println(a)

	// b = float64(a) //integer(1) --------> float(1)

	//Math Library

	fmt.Println(math.Ceil(3.99999)) //it will gives us '4'
	fmt.Println(math.Min(1, 0))     //0
	fmt.Println(math.Max(1, 0))     //1
	fmt.Println(math.Abs(-1))       //1
	fmt.Println(math.Sqrt(100))     //10
	fmt.Println(math.Pow(2, 4))     //16

	//Complex Numbers ----> using builtin library
	fmt.Println(complex(1, 2)) //1+2i

}
