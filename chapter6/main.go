package main

import "fmt"

func average(xs []float64) float64 {
	total := 0.0
	for _, element := range xs { //I do not use the i so I put _ to denote nothing
		total += element
	}
	return total / float64(len(xs))
}

func averageVariadic(xs ...float64) float64 {
	total := 0.0
	for _, element := range xs { //I do not use the i so I put _ to denote nothing
		total += element
	}
	return total / float64(len(xs))
}
func makeEvenGenerator() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

func panicTest() {
	defer func() {
		str := recover()
		fmt.Println("Result from recover", str)
	}()
	panic("PANICO LOCOOO")
	fmt.Println("This never executes")
}

func zero(xPtr *int) { // a pointer to an int
	*xPtr = 0
}

func factorial(x uint) uint {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}

func main() {
	xs := []float64{98, 93, 77, 82, 83}
	fmt.Println(average(xs))            //print
	fmt.Println(averageVariadic(xs...)) //unpacking the slice
	add := func(x, y int) int {
		return x + y
	}
	fmt.Println(add(1, 1))
	nextEven := makeEvenGenerator()
	fmt.Println("Next Even Number", nextEven())
	fmt.Println("Next Even Number", nextEven())
	fmt.Println("Next Even Number", nextEven())
	xf := uint(5)
	fmt.Println("Factorial of", xf, "is ", factorial(xf))
	panicTest()
	xpointer := 8
	fmt.Println("X pointer before modifying", xpointer)
	zero(&xpointer)
	fmt.Println("X pointer after modifying", xpointer)
}
