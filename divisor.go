// Write a program that asks for an initial integer. The program then repeatedly prompts for integers. If the integer i is a divisor of the initial integer, print 'i is a divisor'.
// If the integer i is not a divisor of the initial integer, print 'i is not a divisor'. The program should continue until it is given the value 0 as input.

package main

import (
	"fmt" //,
)

func main() {
	var integralInput int
	var followingNum int
	var numbers []int

	fmt.Print("Enter a base number: ") // initial integer
	fmt.Scanln(&integralInput)         // scan previous line's input and append user input to integralInput

	for {
		fmt.Print("Enter a divisor: ")
		fmt.Scanln(&followingNum) // same thing as previous scanln, but with followingNum.

		if followingNum == 0 {
			fmt.Println()
			break
		}
		numbers = append(numbers, followingNum)
	}

	for _, n := range numbers {
		if integralInput%n == 0 {
			fmt.Println(n, "is a divisor")
		} else {
			fmt.Println(n, "is not a divisor")
		}
	}
}
