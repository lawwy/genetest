package utils

var current int = 0

func AutoId() int {
	current = current +1
	return current
}