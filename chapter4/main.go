package main

import "fmt"

func main() {
	for i := 1; i <= 10; i = i + 1 {
		fmt.Println(i)
	}
	for j := 1; j <= 10; j++ {
		fmt.Println(j)
	}
	for j := 1; j <= 10; j++ {
		if j%2 == 0 {
			fmt.Println(j, " is even")
		} else {
			fmt.Println(j, " is odd")
		}
	}
	for j := 0; j <= 10; j++ {
		switch j {
		case 0:
			fmt.Println("Zero")
		case 1:
			fmt.Println("One")
		case 2:
			fmt.Println("Two")
		default:
			fmt.Println("Unknown Number")
		}
	}
	for j := 1; j <= 100; j++ {
		if j%3 == 0 {
			fmt.Println(j, " divisible by 3")
		}
	}
	for j := 1; j <= 100; j++ {
		x := true
		switch x {
		case ((j%3 == 0) && (j%5 == 0)):
			fmt.Println(j, " FizzBuzz")
		case (j%3 == 0):
			fmt.Println(j, " Fizz")
		case (j%5 == 0):
			fmt.Println(j, " Buzz")
		}
	}
}
