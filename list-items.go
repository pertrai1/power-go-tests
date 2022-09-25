package main

import "strings"

func ListItems(input []string) string {
	result := "You can see here "
	if len(input) == 0 {
		return ""
	}
	if len(input) == 1 {
		return "You can see " + input[0] + " here."
	}
	if len(input) < 3 {
		return result + input[0] + " and " + input[1] + "."
	}
	result += strings.Join(input[:len(input)-1], ", ")
	result += ", and "
	result += input[len(input)-1]
	result += "."
	return result
}
