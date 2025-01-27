package main

import("fmt")

type Animal interface{
	Sound() string
}

type Dog struct{}
//If the Dog behaviour is there then it barks so we need to write a function that satisfies this
func (d Dog)Sound()string{
	return "Bow!"
}


type Cat struct{}
func (c Cat)Sound()string{
	return "Meow!"
}

// func main(){

// 	Animal := Cat{}

// 	fmt.Println(Animal.Sound())

// 	Animal := Dog{}  In this case we cannot redeclare the variable if use use (:=), It automatically 
					//identifies the interface type whether it is implemented or not
					//So we can use var Animal Animal

// 	fmt.Println(Animal.Sound())

	
// }