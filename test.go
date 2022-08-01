package main

import "strings"

func main() {

}

func generateTheString(n int) string {
	if n%2 == 0 {
		a, b := strings.Repeat("a", n-1), strings.Repeat("b", 1)
		return a + b
	} else {
		return strings.Repeat("a", n)
	}
	return ""
}
