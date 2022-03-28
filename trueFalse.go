package main

import "fmt"

func compare(value int) string {
	//do not change this variable resultMessage, secretValue
	resultMessage := ""
	secretValue := 88

	//Insert your code from here
	if value > secretValue {
		resultMessage = "Bigger than secret value"
	}
	
	if value < secretValue {
		resultMessage = "Smaller than secret value"
	}
	
	if value == secretValue {
		resultMessage = "Same as secret value"
	}
	
	fmt.Println(resultMessage)

	//do not remove this line
	return resultMessge
}

func main() {
	var guess int
	fmt.Println("Enter an integer value: ")
	fmt.Scanln(&guess)
	compare(guess)
}
