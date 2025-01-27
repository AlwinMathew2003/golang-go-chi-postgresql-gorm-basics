package main
import ("fmt";"bufio";"os")

func main(){
	fmt.Println("Enter your name: ")
	reader := bufio.NewScanner(os.Stdin)
	reader.Scan()

	fullname:= reader.Text()

	fmt.Printf("My name is %s",fullname[:])

}