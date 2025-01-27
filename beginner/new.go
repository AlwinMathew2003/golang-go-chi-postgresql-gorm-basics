package main
import ("fmt")

func main(){
	var str string = "Hello World"
	for i:=0;i<len(str);i++{
		fmt.Println(string(str[i]))
	}
}