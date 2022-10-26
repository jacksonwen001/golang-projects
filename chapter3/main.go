package main

import (
	"fmt"
	"regexp"
)

// map
func maps() {
	// 直接声明
	m := map[string]int{
		"john": 20,
	}

	fmt.Println(m)

	m1 := make(map[string]string)
	m1["name"] = "john"
	fmt.Println(m1)
	for k, v := range m1 {
		fmt.Printf("%s : %s \n", k, v)
	}
	// map 不要直接赋值 nil  否则就会崩溃，而是赋值一个空的 map
	m3 := make(map[string]string)
	fmt.Println(m3 == nil)
	// 崩溃
	// m3 = nil
	// m3["test"] = "sss"
}

// struct
// 通常字段都是大写开头， 大写开头表示public ，小写表示 private，只在包的范围内使用
// regexp

func reg() {
	re := regexp.MustCompile(`^[A-Z].*?,.*?T`)
	s := re.Match([]byte("HELLO"))
	fmt.Println(s)
	ss := re.FindString("HELLO, ,This is GOLANG")
	fmt.Println(ss)
}

func main() {
	// maps()
	reg()
}
