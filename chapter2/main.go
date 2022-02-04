package main

import "fmt"

func main() {
	fmt.Println("1 +1 =", 1+1)
	fmt.Println("1 +1 =", 1.0+1.0)
	fmt.Println(len("Hola mundo"))
	fmt.Println("Hola mundo"[1])
	fmt.Println("Hola " + "mundo")
	fmt.Println("true && true", true && true)
	fmt.Println("true && false", true && false)
	fmt.Println("true || true", true || true)
	fmt.Println("true || false", true || false)
	fmt.Println("!true", !true)
	fmt.Println("!false", !false)
	fmt.Println("32132 * 42452 =", 32132*42452)
	fmt.Println("(true && false) || (false && true) || !(false && false)", (true && false) || (false && true) || !(false && false))
}
