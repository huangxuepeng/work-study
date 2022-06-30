package main

var z *int

func main() {
	a := 1
	func() {
		// fmt.Println(a)
		a = a + 1
	}()
	// fmt.Println(a)
}
