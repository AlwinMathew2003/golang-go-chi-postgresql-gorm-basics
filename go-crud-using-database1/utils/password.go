package utils

import("golang.org/x/crypto/bcrypt")

//hashing password
func HashPassword(password string)(string,error){

	//for hashing bcrypt.GenerateFromPassword(password in bytes,cost) we should mention the default cost (bcrypt.DefaultCost)
	hash,err:=bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)

	return string(hash),err//we should pass the hashed password as string
}

//verifying password
func VerifyPassword(hashedPassword string, Password string)error{
	//It is used to compare the hashed password with the plain text password from the user
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword),[]byte(Password))
}

