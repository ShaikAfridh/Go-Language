package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	s := "Hello World" //this is a string

	fmt.Println(s)
	fmt.Printf("%T", s) //string
	fmt.Println()
	//length of the string
	fmt.Println(len(s)) //11

	//Read individual characters of string
	fmt.Println(s[0])      //it shows ASCII value of character s[0]------> H {it's ASCII value is 72}
	fmt.Printf("%c", s[0]) //H
	fmt.Println()
	fmt.Println(s[0:6]) //Hello
	fmt.Println(s[:6])  //Hello {by default it takes '0' as starting index}
	fmt.Println(s[6:])  //World {by default it takes 'n-1' character here n is length of the character}

	//Concatenation of String

	s = s + "again"
	s += "again"
	fmt.Println(s) //Hello Worldagainagain

	//String Literals
	fmt.Println("Hello \n World") // '\n'-------> new line
	fmt.Println("Hello \t World") // '\t'-------> new tab
	fmt.Println("Hello \b World") // '\b'-------> back space

	//slice ------>it is a Data Structure  and it is similar to "ARRAY LIST" in java
	abc := "abc"
	b := []byte(abc)
	fmt.Printf("%s,%s", abc, b) //it will print abc twice
	fmt.Println()
	abc2 := string(b)
	fmt.Println(abc2)

	//String Library

	fmt.Println(strings.ToUpper("Hello World"))            //HELLO WORLD
	fmt.Println(strings.ToLower("Hello World"))            //hello world
	fmt.Println(strings.HasPrefix("Hello World", "Hello")) //true
	fmt.Println(strings.HasSuffix("Hello World", "Hello")) //false
	fmt.Println(strings.HasSuffix("Hello World", "World")) //true
	fmt.Println(strings.Replace("Hello World World", "World" /*old*/, "afridi" /*new*/, 1 /*no.of place to replace string*/))
	fmt.Println(strings.Replace("Hello World World", "World" /*old*/, "afridi" /*new*/, 2 /*no.of place to replace string*/))

	//String Builders -----> it is a tool that we can utilize to build our strings

	var sb strings.Builder
	fmt.Println("This is a String BUilder: ", sb.String())
	fmt.Println(sb.Cap()) //it gives the capactiy of the string------>'0'
	fmt.Println(sb.Len()) //it gives the length of the string----->'0'
	sb.WriteString("Hello ")
	fmt.Println("This is a String Builder:", sb.String())
	// fmt.Println(sb.Cap()) //0 2 4 8 {keep on doubleing}------>'8'
	// fmt.Println(sb.Len()) //it gives the length of the string----->'6'

	//increase capacity manually by using
	sb.Grow(100)
	fmt.Println(sb.Cap()) //100+ 8{cap} * 2 {it's doubling up }
	fmt.Println(sb.Len()) //6
	fmt.Println(sb.String())
	sb.Reset()
	fmt.Println(sb.Cap())        //it gives again the capactiy of the string------>'0'
	fmt.Println(sb.Len())        //it gives again the length of the string----->'0'
	sb.Write([]byte{65, 65, 65}) //appending A three times using byte{ASCII value[65]}
	fmt.Println(sb.String())

	//Conversion of String and Other data types

	x := 123
	fmt.Printf("%T", x) //integer
	fmt.Println()
	y := strconv.Itoa(x) //converting int ------->string
	fmt.Println(y)
	fmt.Printf("%T", y) //string
	fmt.Println()
	z, _ := strconv.Atoi(y) //converting string------> int
	//in above return type consists of two things (int, error(_)) we don't care about error it's a way of error handling
	fmt.Println(z)
	fmt.Printf("%T", z) //integer
}
