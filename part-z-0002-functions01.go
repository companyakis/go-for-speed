package main

import "fmt"

func sum_ages(age1 uint16, age2 uint16, age3 uint16) uint16 {
	
	return age1 + age2 + age3
}

func main() {

	fmt.Println(sum_ages(16, 76, 95)) // 187
}
