package main

import "strings"

func ListItems(input []string) string {
	switch len(input) {
	case 0:
		return ""
	case 1:
		return "You can see " + input[0] + " here."
	case 2:
		return "You can see here " + input[0] + " and " + input[1] + "."
	default:
		return "You can see here " + strings.Join(input[:len(input)-1], ", ") + ", and " + input[len(input)-1] + "."
	}
}
