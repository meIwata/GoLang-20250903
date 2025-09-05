package main

import "fmt"

func main() {

	var accountBalance = 1000.0

	for i := 0; i < 200; i++ {
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
		} else if choice == 2 {
			fmt.Print("請輸入存款金額: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("存款金額必須大於0")
				//return // 結束程式
				continue // 回到迴圈開頭
			}

			accountBalance += depositAmount
			fmt.Println("存款成功! 目前餘額: ", accountBalance)
		} else if choice == 3 {
			fmt.Print("請輸入提款金額: ")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)

			if withdrawAmount <= 0 {
				fmt.Println("提款金額必須大於0")
				return // 結束程式
			}
			if withdrawAmount > accountBalance {
				fmt.Println("餘額不足!")
				return // 結束程式
			} else {
				accountBalance -= withdrawAmount
				fmt.Println("提款成功! 目前餘額: ", accountBalance)
			}
		} else if choice == 4 {
			fmt.Println("感謝使用GO bank, 再見!")
			//return // 結束程式
			break // 跳出迴圈
		}

		//fmt.Println("選擇的服務為: ", choice)
	}

	fmt.Println("謝謝惠顧GO bank") //使用break後會執行到這裡

}
