package main

func main() {

}
func convert(s string, numRows int) string {
	if len(s) <= numRows {
		return s
	}
	n := len(s)
	t := 2*numRows - 2
	arr := []byte{}
	for i := 0; i < n; i++ {
		for j := 0; j+i < n; j++ {
			arr = append(arr, s[i+j])
			if i > 0 && i < n-1 && j+t-i < n {
				arr = append(arr, s[j+t-i])
			}
		}
	}
	return string(arr)
}
