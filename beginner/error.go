package main
import ("fmt";"errors")

func div(a int,b int)(int,error){
	if(b==0){
		return 0,errors.New("Cannot Divide by Zero")
	}
	return a/b,nil
}
func main(){
result,err:=div(10,2)
if err!=nil{
	fmt.Printf("Error:%s",err)
}else{
	fmt.Printf("Result:%d",result)
}
}