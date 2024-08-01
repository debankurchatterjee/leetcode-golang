package main

import "fmt"

func convert(s string, numRows int) string {
	if len(s) == 1 || len(s) == numRows || numRows <= 1 {
		return s
	}
	results := make([][]rune, numRows)
	flag := false // toggle flag when we move upwards and downwards
	i := 0
	for _, ch := range s {
		results[i] = append(results[i], ch)
		if i == 0 || i == numRows-1 {
			flag = !flag
		}
		if flag {
			i++
		} else {
			i--
		}
	}
	resultString := ""
	for _, r := range results {
		fmt.Println(string(r))
		resultString = resultString + string(r)
	}
	return resultString
}

func main() {
	fmt.Println(convert("AB", 1))
}
