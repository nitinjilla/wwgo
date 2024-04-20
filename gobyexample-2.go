package main

import (
	"fmt"
)

func main() {

	fmt.Printf("Printing from main for singleReturnValue(): %s\n", singleReturnValue())

	aa, ab := multipleReturnValues()
	fmt.Printf("Printing from main for multipleReturnValues(): %v, %v\n", aa, ab)

	ac, ad := differentReturnValues()
	fmt.Printf("Printing from main for multipleDifferentReturnValues(): %v, %v\n", ac, ad)

	fmt.Printf("Printing from main for iTakeArguments(): %v\n", (iTakeArguments(10, 10)))

	fmt.Printf("Printing from main for variadicFunction(): %v\n", (variadicFunction(10, 10, 10)))

	message := closureFunc()
	fmt.Println(message())

	song := "dosed_rhcp_dwnld.mp4" // working with pointers
	fmt.Printf("/Music/%s\n", song)
	renameFile(&song)
	fmt.Println("File successfully renamed to", song)

	a := 10
	normal(&a)
	fmt.Printf("How did a become %v?", a)
}

func normal(i *int) {
	*i = 0 // changes value of a from 10 to 0
}
func singleReturnValue() string {
	m := "hello"
	return m
}

func multipleReturnValues() (aa, ab int) {
	num := 10
	numtoo := 20
	return num, numtoo
}

func differentReturnValues() (ac string, ad int) {
	str := "hello"
	num := 20
	return str, num
}

func iTakeArguments(a, b int) int {
	return a + b
}

func variadicFunction(ns ...int) int {
	sum := 0

	for _, i := range ns { // 	for i := range ns {
		sum = sum + i // 		sum = sum + ns[i]
	} // 						}

	return sum

}

func closureFunc() func() string {

	return func() string {
		return "hello"
	}
}

func renameFile(sn *string) string {
	*sn = "dosed.mp4"
	return *sn

}
