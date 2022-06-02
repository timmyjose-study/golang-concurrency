package main

import "fmt"

func main() {
	defer fmt.Println("f")

	{
		defer fmt.Println("e")

		{
			defer fmt.Println("d")
			{
				defer fmt.Println("c")
			}
		}
	}

	fmt.Println("a")
	defer fmt.Println("b")
}