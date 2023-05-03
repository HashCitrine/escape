package main

import "fmt"

func NameArr() {
	fmt.Println("{")
	for i := 0; i < len(fieldArr); i++ {
		var tempArr = fieldArr[i]

		if i > 0 {
			fmt.Println(",")
		}
		fmt.Print("\t{")
		for j := 0; j < len(tempArr); j++ {
			if j > 0 {
				fmt.Print(", ")
			}

			// fmt.Print(GetName(tempArr[j]))
		}
		fmt.Print("}")
	}
	fmt.Println()
	fmt.Print("}")
}

func FieldArr() {
	for i := 0; i < len(fieldArr); i++ {
		var tempArr = fieldArr[i]

		if i > 0 {
			fmt.Println("")
		}
		// fmt.Print("\t")
		for j := 0; j < len(tempArr); j++ {
			if j > 0 {
				fmt.Print("\t")
			}

			// fmt.Print(GetField(tempArr[j]))
		}
	}
}