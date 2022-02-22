package main

import "fmt"

func main() {
	var numbers [5]int
	for i := 0; i < 5; i++ {
		numbers[i] = i
	}

	stooges := [...]string{"Moe", "Larry", "Curly"}

	fmt.Println(stooges)
	// stooges = append(stooges)  . Will give err as stooges is an array , not a slice

	slice := []string{"Moe", "Larry", "Curly"}

	slice = append(slice, "Test")
	fmt.Println(slice)

	arr := [4]int{3, 2, 5, 4}
	longerArr := [4]int{5, 7, 1, 2}
	fmt.Println(longerArr)
	// Only if two arrays size is same, we can assign one array to other and compare them
	if longerArr == arr {
		fmt.Println("Equals")
	}
	longerArr = arr
	fmt.Println(longerArr)

	// It follows then that slices of different lengths can be assigned each other’s values. Their type is the same, with the pointer, length, and capacity changing:

	slice1 := []int{6, 1, 2}
	slice2 := []int{9, 3}

	// slices of any length can be assigned to other slice types
	slice1 = slice2
	fmt.Println(slice1)
	//fmt.Println(slice1 == slice2) // Cannot compare slices using ==
}
