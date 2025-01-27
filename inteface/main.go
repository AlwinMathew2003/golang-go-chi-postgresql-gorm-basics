package main

import("fmt")

//It is basically the interface
type PaymentMethod interface{
	Pay(amount float64) string
}

//Credit card implementation
//This is used to store the creditCard details in the from of struct
type creditCard struct{
	cardNumber string
}

//This function means that the this Pay method belong to creditCard object
func (cc creditCard)Pay(amount float64) string{
	return fmt.Sprintf("Paid %.2f using credit card(%s)",amount,cc.cardNumber)
}

type payPal struct{
	email string
}

//This is
func (cc payPal)Pay(amount float64)string{
	return fmt.Sprintf("Paid %.2f using paypal email(%s)",amount,cc.email)
}

//This function receives the paymentMethod which is a type created using interface
func ProcessPayment(method PaymentMethod,amount float64){
	result:=method.Pay(amount)
	fmt.Println(result)
}

func main(){

	//An instance is created with this object type and stored in the memory
	creditCard1 := creditCard{cardNumber:"12234231432"}
	ProcessPayment(creditCard1,100.50)//On callin this it invokes that function


	payPal1 := payPal{email:"hai@gmail.com"}
	ProcessPayment(payPal1,200.50)


}