package main
import ("fmt";"time")

func printMessage1(msg string){
	for i:=1;i<=5;i++{
	fmt.Printf("1: %s %d\n",msg,i)
	time.Sleep(2*time.Second)
	}
}

func printMessage(msg string){
	for i:=1;i<=5;i++{
	fmt.Printf("2: %s %d\n",msg,i)
	time.Sleep(3*time.Second)
	}
}

func main(){
	go printMessage("Task 1")
	go printMessage1("Task 2")

	time.Sleep(30*time.Second)
	fmt.Println("All process completed")

}