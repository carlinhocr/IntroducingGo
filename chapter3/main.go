package main

import "fmt"

var globalName string = "que nombre"

func doubleNumber() {
	fmt.Println("enter a number: ")
	var inputNumber float64
	fmt.Scanf("%f", &inputNumber)
	result := inputNumber * 2
	fmt.Println("resultado loco: ", result)
}

func fahrenheitToCelcius() {
	fmt.Println("enter temperature in farenheit: ")
	var inputNumber float64
	fmt.Scanf("%f", &inputNumber)
	result := (inputNumber - 32) * 5 / 9
	fmt.Println("resultado en celcius: ", result)
}

func feetToMeters() {
	fmt.Println("enter distance in feet: ")
	var inputNumber float64
	fmt.Scanf("%f", &inputNumber)
	result := inputNumber * 0.3048
	fmt.Println("resultado in meters: ", result)
}

func main() {
	var x string = "Hola, mundoo"
	fmt.Println(x)
	x = "re hola mundooo"
	fmt.Println(x)
	x += " pepin"
	fmt.Println(x)
	var y string = "pedro"
	fmt.Println(x == y)
	z := "jose"
	fmt.Println(z)
	fmt.Println(globalName)
	carlos()
	const w string = "constante no se cambia"
	//w = "hola" this gives an error because w is a constant
	var (
		a = 1
		b = 3
		c = 7
	)
	fmt.Println(a, b, c)
	doubleNumber()
	fahrenheitToCelcius()
	feetToMeters()

}

func carlos() {
	fmt.Println(globalName)
}
