package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Welcome to my quiz game!")
	var name string

	fmt.Printf("Enter your name :")
	fmt.Scan(&name)
	fmt.Printf("Hello %v,Welcome to the game!", name)

	fmt.Printf("\nEnter your age:")
	var age uint
	fmt.Scan(&age)

	if age >= 15 {
		fmt.Printf("You can play\n")
	} else {
		fmt.Printf("You can't play the game\n")
		return
	}

	fmt.Printf("Continue\n")
	score := 0
	num_questions := 2
	fmt.Printf("What is better, the RTX 3080 or RTX 3090\n")
	var answer string
	var answer2 string
	fmt.Scan(&answer, &answer2)

	if strings.ToUpper(answer)+" "+strings.ToUpper(answer2) == "RTX 3090" {
		fmt.Printf("Correct\n")
		score++
	} else {
		fmt.Printf("Incorrect!\n")
	}

	fmt.Printf("How many cores in Ryzen 9 3900x have ?\n")
	var cores uint
	fmt.Scan(&cores)

	if cores == 12 {
		fmt.Printf("Correct\n")
		score++
	} else {
		fmt.Printf("Incorrect!\n")
	}

	fmt.Printf("Your score is : %v", score/num_questions*100)

}
