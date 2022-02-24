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

func fibonacci(x int) int {
	if x == 0 {
		return 0
	}
	if x == 1 {
		return 1
	}
	return fibonacci(x-1) + fibonacci(x-2)
}

func makeOddGenerator() func() uint {
	i := uint(1)
	return func() (ret uint) {
		ret = i
		i += 1
		if i%2 == 0 {
			i += 1
		}
		return

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
	fmt.Println("zero xPtr address", &xPtr)
}

func factorial(x uint) uint {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}

func suma(args []int) (total int) {
	total = 0
	for _, value := range args {
		total += value
	}
	return
}

func half(x int) (int, bool) {
	halved := x / 2
	fmt.Println(halved % 2)
	even := false
	if halved%2 == 0 {
		even = true
	} else {
		even = false
	}
	return halved, even
}

func max(args ...int) int {
	maximum := 0
	for _, v := range args {
		if v > maximum {
			maximum = v
		}
	}
	return maximum
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
	fmt.Println(("total of suma"), suma([]int{1, 2, 3}))
	result, even := half(5)
	fmt.Println("the half of 5, is:", result, "is it Even?", even)
	maxlist := []int{1, 4, 56, 77, 99, 5, 44}
	fmt.Println("Maximum of 1,4,56,77,99,5,44 is", max(maxlist...))
	nextEven2 := makeOddGenerator()
	fmt.Println("Next Odd Number", nextEven2())
	fmt.Println("Next Odd Number", nextEven2())
	fmt.Println("Next Odd Number", nextEven2())
	fmt.Println("Finocacci 0", fibonacci(0))
	fmt.Println("Finocacci 1", fibonacci(1))
	fmt.Println("Finocacci 2", fibonacci(2))
	fmt.Println("Finocacci 5", fibonacci(5))
	fmt.Println("Finocacci 10", fibonacci(10))
}
