package main

import (
	"flag"
	"fmt"
	"log"
	"log/syslog"
)

// simple func for fmt
func print() {
	fmt.Print("不空一行")
	fmt.Printf("格式化 Hello,  %s", "jackson")
	fmt.Println("空一行")
}

// controlling program flow

func programFlow() {
	fmt.Println("Go support: if-else / switch ")
	answer := "Ok"
	if answer == "Ok" {
		fmt.Println("the answer is OK")
	} else {
		fmt.Println("the answer is No")
	}

	fruit := "apple"
	switch fruit {
	case "apple":
		fmt.Println("the fruit is apple")
		// fallthrough 接下去的条件都会运行 基本不使用
	case "banana":
		fmt.Println("the fruit is banana")
	default:
		fmt.Println("there are no fruit!")
	}
}

func loop() {
	fmt.Println("for / range")
	for i := 0; i < 10; i++ {
		fmt.Printf("the number i is: %d \n", i)
	}
	fmt.Println("Go 不支持 while 循环，所以我们也使用 for")

	// 无限循环
	i := 1
	for {
		if i == 10 {
			fmt.Printf("the num is : %d break\n", i)
			break
		}
		i++
	}

	slices := []int{1, 2, 3, 4, 5}
	for _, v := range slices {
		fmt.Printf("the number is %d in slices \n", v)
	}
}

// accept user input
func userInput() {
	var name string
	fmt.Scanln(&name)
	fmt.Printf("Hello %s \n", name)
}

// flag
func flagParse() {
	var host string
	var port int
	// 变量， 名字， 默认值，使用方法
	flag.StringVar(&host, "h", "localhost", "Please input your host, default is localhost")
	flag.IntVar(&port, "p", 8080, "-p=8080")
	flag.Parse()

	fmt.Printf("The host is %s\n", host)
	fmt.Printf("the port is %d\n", port)
}
func logInfo() {
	sys, err := syslog.New(syslog.LOG_SYSLOG, "systemLog.go")
	if err != nil {
		log.Println(err)
		return
	} else {
		log.SetOutput(sys)
		log.Print("Everything is fine")
	}

	// log.Fatal() 发生错误，不排除堆栈, 退出程序
	// log.Panic() 发生错误，抛出错误，退出程序

}

// main func
func main() {
	fmt.Println("Hello Golang")
	// print()
	// programFlow()
	// loop()
	// userInput()
	// flagParse()
	// logInfo()
}
