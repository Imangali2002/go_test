package main

import "fmt"

func add(n int) (int, error) {
	if n < 0 {
		return -1, fmt.Errorf("SUBZERO_ERROR")
	}
	return n + 10, nil
}

func main() {
	add(10)
}
