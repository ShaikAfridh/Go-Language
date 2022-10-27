//Loops

package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	fmt.Println("Hello World")
	fmt.Println("Hello World")

	// i := 0
	//infinite loop
	// for {
	// 	fmt.Println(i)
	// 	i = i + 1 //this loop dosen't break any where
	// }
	// for i < 10 { //here loop will break i<10 dosen't satisify condition {0-9}
	// 	fmt.Println(i)
	// 	i += 1
	// }

	// //we can also declare variable 'i' value in for loop
	// for i := 0; i < 10; {
	// 	fmt.Println(i)
	// 	i += 1
	// }

	//we can also move incrementing value in for loop
	// for i := 0; i < 10; i += 1 {
	// 	fmt.Println(i)
	// }

	//declaring multiple variables

	// for i, j := 0, 1; i < 10 && j < 10; i, j = i+1, j+1 {
	// 	fmt.Println(i, j) //0, j=0+1 =1 ,1+1=2 -----> 0 1
	// 	                  //1, j=1+1 =2 , 2+1=3 ----->1 2
	// 					  //2, j=
	// }

	//for each loop

	s := "Hello World"
	for i := 0; i < len(s); i++ {
		fmt.Printf("%d %c", i, s[i])
		fmt.Println()
		//i=0, 0<len(s)--->11 ; i=0+1
		// i=0 , s[0] =H
		//i=1, 1<11;i=1+1
		//i=1, s[1] =e
		//-------------------------------------
	}
	fmt.Print()
	for i, c := range s { //here we have created for each loop in c var we stored range of s
		fmt.Printf("%d %c", i, c) //c= 0 to 10
		fmt.Println()
	}
	fmt.Println()
	//keywords in loops

	//break
	for _ /*here we don't care about we focus on content so i put '_'*/, c := range s {
		if c == ' ' {
			break //Hello
		}
		fmt.Printf("%c", c)
		fmt.Println()
	}
	fmt.Println()
	//Continue
	for _, c := range s {
		if c == ' ' { //it will skip the value here it skiped ' ' character
			continue
		}
		fmt.Printf("%c", c)
		fmt.Println()
	}
	fmt.Println()
	//Label
iForLoop:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if j == 3 {
				break iForLoop
			}
			fmt.Printf("%d%d ", i, j)
		}
		fmt.Println()
	}
}
