package main

import (
	"errors"
	"fmt"
	"io"
)

// 定义一个 Ustr 类型（以 string 为基类型）
type Ustr string

// 实现 Ustr 类型的 Read 方法
func (s Ustr) Read(p []byte) (n int, err error) {
	i, ls, lp := 0, len(s), len(p)
	for ; i < ls && i < lp; i++ {
		// 将小写字母转换为大写字母，然后写入 p 中
		if s[i] >= 'a' && s[i] <= 'z' {
			p[i] = s[i] + 'A' - 'a'
		} else {
			p[i] = s[i]
		}
	}
	// 根据读取的字节数设置返回值
	switch i {
	case lp:
		return i, nil
	case ls:
		return i, io.EOF
	default:
		return i, errors.New("Read Fail")
	}
}
func main() {
	us := Ustr("hello world!")
	buf := make([]byte, 32)
	n, err := io.ReadFull(us, buf)
	fmt.Println(string(buf))
	fmt.Println(n, err)
}
