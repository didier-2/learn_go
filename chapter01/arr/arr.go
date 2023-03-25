package main

import "fmt"

func main() {
	var age int = 35
	fmt.Println(age)
	var ages = [5]int{32, 35, 56, 89, 2}
	fmt.Println(ages)
	var age3 [5]int = [5]int{32, 56, 55, 88, 23}
	fmt.Println(age3)
	age4 := [5]int{32, 265, 59, 56, 59}
	fmt.Println(age4)
	age2 := [3]int{}
	age2[0] = 1
	age2[1] = 2
	age2[2] = 3
	fmt.Println(ages)

	for i := 0; i < len(age2); i++ {
		fmt.Println(age2[i])
	}
	for i, val := range age2 {
		fmt.Println(age2[i], i, val)

	}

}
