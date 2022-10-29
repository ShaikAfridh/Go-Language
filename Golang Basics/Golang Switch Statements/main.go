package main

import "fmt"

func main() {

	venues := []string{"Home", "Work", "School", "Restarunt", "Gym"}

	//Depending on different venues we have different greeting

	for _, venue := range venues {
		if venue == "Home" {
			fmt.Println("Hey Mom, I'm Hungry")
		} else if venue == "Work" {
			fmt.Println("Let's try some Code")
		} else if venue == "School" {
			fmt.Println("It's kinda Boring to goto School")
		} else if venue == "Restarunt" {
			fmt.Println("Let's Order today Special item")
		} else if venue == "Gym" {
			fmt.Println("Uhhh.... It's Hard to maintain Healthy Body")
		} else {
			fmt.Println("Hello")
		}
	}

	//There is another Smarter way to do this
	for _, venue := range venues {
		switch venue {
		case "Home":
			fmt.Println("Hey Mom, I'm Hungry")

		case "Work":
			fmt.Println("Let's try some Code")
		case "School":
			fmt.Println("It's kinda Boring to goto School")
		case "Restarunt":
			fmt.Println("Let's Order today Special item")
		case "Gym":
			fmt.Println("Uhhh.... It's Hard to maintain Healthy Body")
		default:
			fmt.Println("Hello")
		}
	}
	//tip:- In Other Programming Languages like c,c++,java we add break staement at end of the case in Golang we don't need to add any break statements By default it will automatically break the switch statement

}
