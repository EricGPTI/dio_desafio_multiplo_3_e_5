package main

import (
	"fmt"
)

func PrintPin() {
	fmt.Println("Pin")
}

func PrintPan() {
	fmt.Println("Pan")
}

func PrintPinPan() {
	fmt.Println("Pin Pan")
}

func main() {
	for i := 0; i <= 100; i++ {
		if i % 3 == 0 && i % 5 == 0 {
			PrintPinPan()
		} else if i % 3 == 0 {
			PrintPin()
		} else if i % 5 == 0 {
			PrintPan()
		} else {
			fmt.Println(i)
		}
	}
}
