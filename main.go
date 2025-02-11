package main

import (
	"fmt"
	"practice-go/newpackage"
)

/*go does not support classes and objects
do struct is a miniature replacememt oops */
//declaration
type student struct {
	name  string
	class int
	roll  int
}

func main() {
	fmt.Println("Hello !!!")
	sayHi("Abhi")
	//use func from different package
	fmt.Println("Lets add two numbers")
	result := newpackage.Myaddition(177, 285)
	fmt.Printf("Result is %v \n\n", result)

	//multiple variable declaration
	a, b, c, age := "Abhi", "kumar", "singh", 36
	fmt.Printf("type of a is %T, type of b is %T, type of c is %T, type of age is %T \n", a, b, c, age)
	fmt.Println(a, b, c)

	// Getting user input
	var name string
	fmt.Println("Enter your name")
	fmt.Scan(&name)
	fmt.Printf("Welcome %v \n", name)

	//slices or list
	var names []string
	names = append(names, "Mukesh")
	fmt.Println(names)
	morenames := []string{"Ramesh", "Suresh"}
	allnames := append(names, morenames...)

	//fmt.Println(allnames)
	//loops - there is only one loop, for loop
	for index, name := range allnames {
		fmt.Println(index, name)
	}
	// switch statement
	var city string
	fmt.Println("Enter your city")
	fmt.Scan(&city)

	switch city {
	case "Ranchi":
		fmt.Println("you are in Ranchi")
	case "Patna":
		fmt.Println("You are in Patna")
	default:
		fmt.Println("Invalid city")
	}

	// maps in go lang
	newmap := map[string]int{}
	newmap["abhi"] = 36
	newmap["shyam"] = 34
	for k, v := range newmap {
		fmt.Println(k, v)
	}

	//basic for loop in Go

	for i := 1; i <= 5; i++ {
		fmt.Printf("iteration %v \n", i)
	}
	var student1 student
	var student2 student
	var studentlist []student
	student1.name = "Ravi"
	student1.class = 1
	student1.roll = 11

	student2.name = "Mukesh"
	student2.class = 1
	student2.roll = 12

	studentlist = append(studentlist, student1)
	studentlist = append(studentlist, student2)
	for _, student := range studentlist {
		fmt.Println(student)
	}
}

func sayHi(name string) {
	fmt.Printf("Hi There mr %v, how are you today\n", name)
}
