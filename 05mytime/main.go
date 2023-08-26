package main

import (
	"fmt"
	"reflect"
	"time"
)

func main() {

	fmt.Println("Welcome to time study of Golang")

	//Current Time

	presentTime := time.Now()
	fmt.Println("Present Time:", presentTime)

	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))
	// 01-02-2006 15:04:05 Monday ==> Standard for formating

	fmt.Println("Type of presentTime =>", reflect.TypeOf(presentTime))

	createdDate := time.Date(2023, time.August, 26, 17, 30, 0, 0, time.UTC)
	fmt.Println("created Date =>", createdDate)

	//Formating
	fmt.Println("Formated date =>", createdDate.Format("01-02-2006 Monday"))

}
