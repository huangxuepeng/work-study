package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("/bin/sh", "./test.sh")
	bytes, err := cmd.Output()
	if err != nil {
		fmt.Println("cmd.Output:", err)
		return
	}
	fmt.Println(string(bytes))
}

// func solveEquation(equation string) string {
// 	arr := strings.Split(equation, "=")
// 	a, b := []byte(arr[0]), []byte(arr[1])

// 	arr1, arr2 := []string{string(a[0])}, []string{string(b[0])}

// 	for i := 1; i < len(a); i += 2 {
// 		arr1 = append(arr1, string(a[i:i+2]))
// 	}
// 	for i := 1; i < len(b); i += 2 {
// 		arr2 = append(arr2, string(b[i:i+2]))
// 	}
// 	d := map[string]int{}
// 	for _, v := range arr1 {
// 		if len(v) == 1 {
// 			if v == "x" {
// 				d["x"]++
// 			} else {

// 			}
// 		}
// 	}
// }
