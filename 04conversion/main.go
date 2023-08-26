package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("Welcome to our pizza app")
	fmt.Println("Please rate our pizza between 1 to 5")

	reader := bufio.NewReader(os.Stdin)

	rating, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating us,", rating)
	fmt.Printf("The type of rating is %T \n", rating)

	//Getting the type using reflect.TypeOf()

	fmt.Println("The type of rating is", reflect.TypeOf(rating))

	// numRating, err := strconv.ParseFloat(rating, 64)

	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println("Adding 1 to your given rating:", numRating+1)
	// }

	//Throwing an error ==>strconv.ParseFloat: parsing "4\n": invalid syntax
	//When we are entering a rating (a number as a rating),after that we are entering '\n' character by
	//pressing Enter on keyboard.so this whole this is coming to us as input.In this case we need to trim
	//the next line character(/n) by using strings.TrimSpace(rating)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(rating), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Adding 1 to your given rating:", numRating+1)
	}
}
