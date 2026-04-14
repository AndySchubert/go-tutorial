package main

import (
	"fmt"
	"strings"
)

func main() {
	var myString1 = []rune("résumé")
	var myString2 = "résumé"

	fmt.Printf("myString1 is %v of type %T\n", myString1, myString1)
	fmt.Printf("myString2 is %v of type %T\n", myString2, myString2)

	var indexed = myString1[1]
	var indexed1 = myString2[1]
	fmt.Printf("myString1[1] is %v of type %T\n", indexed, indexed)
	fmt.Printf("myString2[1] is %v of type %T\n", indexed1, indexed1)

	fmt.Printf("%v, %T\n", indexed, indexed)
	fmt.Printf("%v, %T\n", indexed1, indexed1)

	for i, v := range myString1 {
		fmt.Println(i, v)
	}

	for i, v := range myString2 {
		fmt.Println(i, v)
	}
	fmt.Printf("\nThe length of the string is %v\n", len(myString1))
	fmt.Printf("\nThe length of the string is %v\n", len(myString2))

	var myRune = 'a'
	fmt.Printf("\nmyRune is %v of type %T\n", myRune, myRune)

	var strSlice = []string{"Hello", "World", "!"}
	var strBuilder strings.Builder
	for i := range strSlice {
		strBuilder.WriteString(strSlice[i])
	}
	var catStr = strBuilder.String()
	fmt.Printf("\n%v\n", catStr)

	// Inefficient way to concatenate strings
	// var catStr = ""
	// for i := range strSlice {
	// 	catStr += strSlice[i]
	// }
	// fmt.Printf("\n%v\n", catStr)
	// // catStr[0] = 'a' // This will cause an error because strings are immutable
}
