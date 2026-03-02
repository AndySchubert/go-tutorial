package main

import "fmt"

func temp() {
	var p *int32 = new(int32)
	var i int32

	fmt.Printf("The value of p points to is: %v", *p)
	fmt.Printf("\nThe value of i is: %v\n", i)

	p = &i
	*p = 10

	fmt.Printf("The value of p points to is: %v", *p)
	fmt.Printf("\nThe value of i is: %v\n", i)

	var slice = []int32{1, 2, 3}
	var sliceCopy = slice
	sliceCopy[2] = 4
	fmt.Println(slice)
	fmt.Println(sliceCopy)

}
