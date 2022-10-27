// Constant

package main

import (
	"fmt"
	"math"
	"math/big"
	"time"
)

func main() {
	//constant
	const five int = 5 //this is a constant variable
	fmt.Println(five)
	//variable
	var six int = 6
	fmt.Println(six)

	//diffence between constant and variable

	six = 10 //we can reassign variable
	fmt.Println(six)

	// five = 7 //can't reassign value to the constant
	fmt.Println(five)

	const pi = 3.14
	fmt.Println(math.Pi)

	//Multiple Constants
	// const (
	// 	a = 1
	// 	b = 2
	// 	c = 3
	// 	d = 4
	// )
	// fmt.Println(a, b, c, d)

	const (
		a = 1
		b //here b takes previous declaration variable value b------>1
		c = 3
		d //d----------->3
	)
	fmt.Println(a, b, c, d) //1 1 3 3

	// const (
	// 	zero int = iota //iota ---> it is constant starting from '0'
	// 	one  //1
	// 	two //2
	// 	three //3
	// )
	// fmt.Println(zero, one, two, three) //0,1,2,3

	//with 'iota' start from 5
	const (
		zero  int = iota + 5 //0+5
		one                  //6
		two                  //7
		three                //8
	)
	fmt.Println(zero, one, two, three) //5,6,7,8

	const (
		_ = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
		eb
		zb
		yb
	)
	fmt.Println(kb, mb, gb)

	//untyped Constant
	// fmt.Println(yb) //we get overflow this is much larger than integer value

	fmt.Println(yb / zb) //1024 we can perform division

	//Math Library Constants

	fmt.Println(math.Pow(2, 100))

	//constant of a large number

	// const largeNum = math.Pow(2, 100)  //here math.pow is not constant so we getting error

	//Builtin Common Constant

	fmt.Println(math.Pi)     //3.141592653589793
	fmt.Println(time.March)  //March
	fmt.Println(time.Second) //1s
	fmt.Println(time.UTC)    //UTC
	fmt.Println(big.MaxExp)  //2147483647
	fmt.Println(big.MinExp)  //-2147483648

}
