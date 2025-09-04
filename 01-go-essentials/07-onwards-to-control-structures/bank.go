package main

import "fmt"

func main() {
	fmt.Println("Welcome to GO bank")
	fmt.Println("選擇你要的服務: ")
	fmt.Println("1. 查詢餘額")
	fmt.Println("2. 存款")
	fmt.Println("3. 提款")
	fmt.Println("4. 退出")

	var choice int
	fmt.Print("請選擇服務: ")
	fmt.Scan(&choice)
	fmt.Println("選擇的服務為: ", choice)
}
