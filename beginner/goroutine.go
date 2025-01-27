package main
import ("fmt";"time")

func hello(){
	fmt.Println("Say Hello")
}

func main(){
	go hello()
	fmt.Println("Hai I am inside main function")
	//It does not give time for go routine to execute just execute the contents in the main funciton
	time.Sleep(time.Second)
}