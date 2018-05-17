package main

import "fmt"

func main(){
	i := -2
	switch {
	case i < 0:
		fmt.Println(i, "Minus")

	case i == 0:
		fmt.Println(i, "Zero")

	case i > 0:
		fmt.Println(i, "Plus")

	}
}
