package main

import "fmt"

func main() {
	type ninja struct {
		name    string //for using this name variable outside we can captilize first letter
		weapons []string
		level   int
	}

	//asigning values in struct
	// afridi := ninja{name: "Afridi"}
	// fmt.Println(afridi)  //{Afridi [] 0} ----->{name [weapons] level}

	// Assigning  All Struct values
	afridi := ninja{
		name:    "Afridi",
		weapons: []string{"Ninja Star"},
		level:   1,
	}
	fmt.Println(afridi)

	//Printing struct individual values
	fmt.Println(afridi.name)
	fmt.Println(afridi.weapons)
	fmt.Println(afridi.level)

	afridi.level++                                         //incrementing struct integer value
	afridi.weapons = append(afridi.weapons, "Ninja Sword") //appending string in strut

	fmt.Println(afridi.level) //2
	fmt.Println(afridi.weapons)
	//assigning struct  into struct

	type dojo struct {
		name  string
		ninja ninja //name type

	}

	golangDojo := dojo{
		name:  "Golang Dojo",
		ninja: afridi,
	}
	fmt.Println(golangDojo) //{Golang Dojo {Afridi [Ninja Star Ninja Sword] 2}}
	//{name         { name [weapons] level} ----> this is old struct value}

	fmt.Println(golangDojo.ninja.level) // struct_var.ninja{dojo struct value}.level{ninja struct value}

	golangDojo.ninja.level = 3
	fmt.Println(golangDojo.ninja.level) //3
	fmt.Println(afridi.level)           //2 {bcoz in 'golangDojo' variable we declared 'ninja' value as a copy of 'afridi' variable}
	fmt.Println(afridi.level)

	//Pointers and Addresses
	type AddressedDojo struct {
		name  string
		ninja *ninja
	}
	addressed := AddressedDojo{"Addressed Golang Dojo", &afridi} //here we have given 'afridi' struct variable address to 'addressed' variable
	fmt.Println(addressed)
	fmt.Println(*addressed.ninja)
	addressed.ninja.level = 4
	fmt.Println(addressed.ninja.level) //4
	fmt.Println(afridi.level)          //4 here we can observe the values are updated {bcoz here we defined 'addressed' variable with address of 'afridi' variable}

	wadePointer := new(ninja) //here 'ninja' is a pointer and we assigning new 'ninja' pointer variable to 'wadePointer'
	fmt.Println(wadePointer)  //&{ [] 0} ----->  address
	fmt.Println(*wadePointer) //{ [] 0} ------> value

	wadePointer.name = "Wade"
	wadePointer.weapons = []string{"Ninja star"}

	fmt.Println(*wadePointer)

	intern := ninjaIntern{name: "Intern", level: 1} //assigning value to the 'intern' with 'ninjaIntern' struct
	intern.sayHello()                               //here 'intern' variable has a  function[sayHello] which is given in 'ninjaIntern' struct in line 81
}

type ninjaIntern struct {
	name  string
	level int
}

func (ninjaIntern) sayHello() { //here we specified 'ninjaIntern' struct in function
	fmt.Println("Hello")
}
