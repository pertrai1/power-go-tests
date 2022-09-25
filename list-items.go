package main

import "strings"

func ListItems(input []string) string {
	result := "You can see here "
	result += strings.Join(input[:len(input)-1], ", ")
	result += ", and "
	result += input[len(input)-1]
	result += "."
	return result
}
