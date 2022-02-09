package main

import "fmt"

func averageTestScore() {
	var x [5]float64
	x[0] = 98
	x[1] = 93
	x[2] = 77
	x[3] = 82
	x[4] = 83
	var total float64 = 0
	for i := 0; i < len(x); i++ {
		total += x[i]
	}
	fmt.Println(total / float64(len(x)))
}

func averageTestScoreIdiomatic() {
	x := [5]float64{98, 93, 77, 82, 83}
	var total float64 = 0
	for _, element := range x { //I do not use the i so I put _ to denote nothing
		total += element
	}
	fmt.Println(total / float64(len(x)))
}

func sliceAppend() {
	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4, 5)
	fmt.Println(slice1, slice2)
}

func sliceCopy() {
	slice1 := []int{1, 2, 3}
	slice2 := make([]int, 2)
	copy(slice2, slice1)
	fmt.Println(slice1, slice2)
}

func mapDemo() {
	x := make(map[string]int)
	x["key"] = 10
	x["otroElemento"] = 15
	fmt.Println(x)
	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["Li"] = "Lithium"
	fmt.Println(elements["Li"])
	elements2 := map[string]map[string]string{
		"H": map[string]string{
			"name":  "Hydrogen",
			"state": "gas",
		},
	}
	if ele, ok := elements2["H"]; ok {
		fmt.Println(ele["name"], ele["state"])
	}
	if ele, ok := elements2["Li"]; ok {
		fmt.Println(ele["name"], ele["state"])
	} else {
		fmt.Println("Element not found")
	}

}

func main() {
	var x [5]int
	x[4] = 100
	fmt.Println(x)
	averageTestScore()
	averageTestScoreIdiomatic()
	sliceAppend()
	sliceCopy()
	mapDemo()
}
