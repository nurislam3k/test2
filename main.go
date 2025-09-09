package main

import (
	"fmt"
	"test/test"
	"test/test/nurislam"
)

func main() {
	fmt.Println("dlfkdkfldkf")
	test.Add()
	nurislam.Nurislam() 
	arip := nurislam.Aruzhan(9)
	if arip == 80 {
		fmt.Println("correct answer")
	} else {
		fmt.Println("wrong answer")
	}
	fmt.Println(arip)
}

