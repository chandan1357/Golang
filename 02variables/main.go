package main

import "fmt"

const LoginToken string = "gibberish" //Public

func main() {
	var username string = "Chandan"
	fmt.Println(username)
	fmt.Printf("Variable is of type : %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type : %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type : %T \n", smallVal)

	var smallFloat float32 = 255.911246232
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type : %T \n", smallFloat)

	//Default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable) //0
	fmt.Printf("Variable is of type : %T \n", anotherVariable)

	var anotherStringVariable string
	fmt.Println(anotherStringVariable) //space
	fmt.Printf("Variable is of type : %T \n", anotherStringVariable)

	//implcit type
	var website = "lco.in"
	fmt.Println(website)
	fmt.Printf("Variable is of type : %T \n", website)

	//no var style
	numberOfUsers := 3000000
	fmt.Println(numberOfUsers)
	fmt.Printf("Variable is of type : %T \n", numberOfUsers)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type : %T \n", LoginToken)
}
