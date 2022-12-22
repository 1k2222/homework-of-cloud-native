package main

import "fmt"

func main() {
	input := []string{"I", "am", "stupid", "and", "weak"}
	for i := range input {
		if i == 2 {
			input[i] = "smart"
		} else if i == 4 {
			input[i] = "strong"
		}
	}
	fmt.Println(input)
}
