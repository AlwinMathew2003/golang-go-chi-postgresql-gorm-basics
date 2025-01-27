package main

import ("fmt")

//Specify the interface generally that accepts the method to send notification
type Message interface{
	send(message string) string //This method accepts a string and returns a string
}

//Implementation using email
type email struct{
	email string
}

func (cc email)send(message string)string{
	return fmt.Sprintf("This message (%s)from :%s",message,cc.email)
}

//Implementation using sms
type SMS struct{
	sms string
}

func (cc SMS)send(message string)string{
	return fmt.Sprintf("This message (%s)from :%s",message,cc.sms)
}

//Implementation using push notificatoin
type pushNotification struct{
	pushNotification string
}

func (cc pushNotification)send(message string)string{
	return fmt.Sprintf("This message (%s) is from:%s",message,cc.pushNotification)
}

func SendingMethod(method Message,message string){
	result:= method.send(message)
	fmt.Println(result)
}

func main(){
	sms := SMS{sms:"23423423423"}
	SendingMethod(sms,"Hai everyone")

	push := pushNotification{pushNotification: "3282918478"}
	fmt.Println(push.send("Hello Every one"))
}

//Interface is a type the defines a set of methods but does not provide their implementation
//Its implementation will be in some other files