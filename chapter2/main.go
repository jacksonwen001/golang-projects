package main

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func goError() error {
	fmt.Println("go 把  error 当做一个 值 来对待")
	val, err := strconv.Atoi("99")
	if err != nil {
		fmt.Printf("err: %s\n", err)
	} else {
		fmt.Println(val)
	}

	// 自定义 error
	return errors.New("this is error")
}

// go 的数值运算不准确 使用 go get github.com/shopspring/decimal
// https://www.imhanjm.com/2017/08/27/go%E5%A6%82%E4%BD%95%E7%B2%BE%E7%A1%AE%E8%AE%A1%E7%AE%97%E5%B0%8F%E6%95%B0-decimal%E7%A0%94%E7%A9%B6/
func bigdecimal() {

}

// 字符和字符串的相互转换
func stringAndByte() {
	bs := []byte{'h', 'e', 'l', 'l', 'o'}
	hello := string(bs)
	fmt.Println(hello)         // hello
	fmt.Println([]byte(hello)) // [104 101 108 108 111]
	b2 := []byte(hello)
	for _, v := range b2 {
		value := string(v)
		fmt.Println(value)
	}
}

// 数字和字符串的相互转换
func stringAndNumber() {
	number, _ := strconv.Atoi("19")
	fmt.Println(number)
	stringNumber := strconv.Itoa(19)
	fmt.Println(stringNumber)
	numberFloat, _ := strconv.ParseFloat("19.2324", 64) // 64, 32
	fmt.Println(numberFloat)

}

// time and date
func timeAndDate() {
	fmt.Println("比较特殊 从2006年开始算起 ")
	fmt.Println("01 表示 月 ")
	fmt.Println("02 表示 日")
	fmt.Println("06 表示 两位数的年， 2006 表示4位数的年")
	fmt.Println("03 表示12小时， 15 表示 24 小时的时")
	fmt.Println("04 表示分")
	fmt.Println("05 表示秒")
	// fmt.Println(time.Now()) //2022-10-26 11:12:44.277316 +0800 CST m=+0.000189834

	parseTime, _ := time.Parse("2006-01-02 15:04:05", "2022-10-26 05:12:44")
	fmt.Println(parseTime)
	fmt.Println(time.Now().Format("02 Jan 06 Monday 03:04:05"))

	now := time.Now()
	loc, _ := time.LoadLocation("America/New_York")
	fmt.Printf("New York Time: %s\n", now.In(loc))

	fmt.Println("add hour minutes seconds: ", time.Now().Add(1*time.Hour))

	fmt.Println("add date: ", time.Now().AddDate(0, 0, 1))
}

// array
func arrays() {
	arr := make([]int, 10)
	fmt.Println(arr)
	arr2 := []int{1, 2, 3, 4}
	arr3 := arr2[0:1:2] // 取一个元素，但是 cap 是 1
	fmt.Printf("arr3: %p\n", arr3)
	arr3 = append(arr3, 10)
	fmt.Printf("new arr3: %p\n", arr3)

	// 切片的 len 和 cap 的区别
	// 如果 容量 > 长度， 那么地址还是同一个地址;
	// 如果 容量 < 长度， 那么就会扩容，变成一个新地址;
}

// pointer
// 基础数据不使用， 但是对象数据使用， 改变源对象的值

// random 随机 是 伪随机
func randNum() {
	fmt.Println(time.Now().UnixMilli())
	rand.Seed(time.Now().UnixMilli())
	fmt.Println(rand.Intn(10))
}

func main() {
	fmt.Println("Basic Go Data Types, Cover: ")
	fmt.Println()
	fmt.Println("error")
	fmt.Println("numeric")
	fmt.Println("non-numeric")
	fmt.Println("go constants")
	fmt.Println("grouping similar data")
	fmt.Println("pointer")
	fmt.Println("random numbers")
	fmt.Println()

	// err := goError()
	// fmt.Println(err)
	// bigdecimal()
	// stringAndByte()
	// stringAndNumber()
	// timeAndDate()
	// arrays()
	randNum()
}
