package main

import "fmt"

func main() {

	var accountBalance = 1000.0

	fmt.Println("Welcome to GO bank")
	fmt.Println("選擇你要的服務: ")
	fmt.Println("1. 查詢餘額")
	fmt.Println("2. 存款")
	fmt.Println("3. 提款")
	fmt.Println("4. 退出")

	var choice int
	fmt.Print("請選擇服務: ")
	fmt.Scan(&choice)

	//wantsCheckBalance := choice == 1
	//if wantsCheckBalance{
	//	fmt.Println("你選擇了查詢餘額")
	//}

	if choice == 1 {
		//fmt.Println("你選擇了查詢餘額")
		fmt.Println("餘額: ", accountBalance)
	}

	fmt.Println("選擇的服務為: ", choice)
}
