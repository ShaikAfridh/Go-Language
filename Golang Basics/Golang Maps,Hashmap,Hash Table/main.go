package main

import "fmt"

func main() {
	var m map[string]string
	fmt.Println(m == nil)                              //true
	m = map[string]string{}                            //we instantiate map object with curly braces{}
	fmt.Println(m == nil)                              //false
	fmt.Println(m)                                     //map[]
	fmt.Println(len(m))                                //0
	m = make(map[string]string, 5 /*map size*/)        //we instantiate map uing 'make' method
	fmt.Println(len(m))                                //0
	m = map[string]string{"Afridi": "DevOps Engineer"} //instantiate map with some predefined elements
	fmt.Println(m)                                     //map[Afridi:DevOps Engineer]
	fmt.Println(len(m))                                //1
	m["Wade"] = "Not a DevOps Engineer"                //here we given key:value in map
	fmt.Println(m)                                     //map[Afridi:DevOps Engineer Wade:Not a DevOps Engineer]

	//By default in map is not Ordered
	m["Wade"] = "Wade is also a DevOps Enginner" //reassigning map value
	fmt.Println(m)                               //map[Afridi:DevOps Engineer Wade:Wade is also a DevOps Enginner]
	//print specific map value
	fmt.Println(m["Wade"]) //   Wade is also a DevOps Enginner

	//delete value from the map
	delete(m, "Wade")
	fmt.Println(m) //map[Afridi:DevOps Engineer]

	//modify content as long as we have key
	m["Wade"] += "He is also an Actor"
	fmt.Println(m) //map[Afridi:DevOps Engineer Wade:He is also an Actor]

	//Maps are irretable  data structures
	//we can use maps foreach loops

	for name, title := range m {
		fmt.Println(name, title)
		//Output:-
		// Afridi DevOps Engineer
		// Wade He is also an Actor
	}

	fmt.Println(m["Wade"]) //He is also an Actor

	m["Wade"] = "Programmer also "
	title, ok := m["Wade"] //here we give 'Wade' map value to 'title'

	//we can use maps in if-else statements
	if ok {
		fmt.Println(title) //printing the title value
		//Programmer also
	} else {
		fmt.Println("We Didn't find Wade")
	}

}
