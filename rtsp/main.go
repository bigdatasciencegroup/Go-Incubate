package main

import "fmt"

func main() {
	fmt.Println("Hi")

	gh := func(){
		fmt.Println("yes sucess")
	}

	gh()
	lo := tester()
	lo()

}

func tester()func(){
	ki := func(){
		fmt.Println("ji")
	}
	return ki
}