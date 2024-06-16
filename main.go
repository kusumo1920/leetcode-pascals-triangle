package main

import "fmt"

func main() {
	input := 5
	output := generate(input)
	fmt.Println(output)
}

func generate(numRows int) [][]int {
	result := make([][]int, 0)
	for i := 0; i < numRows; i++ {
		element := make([]int, i+1)
		element[0], element[i] = 1, 1
		for j := 1; j < i; j++ {
			element[j] = result[i-1][j-1] + result[i-1][j]
		}
		result = append(result, element)
	}
	return result
}
