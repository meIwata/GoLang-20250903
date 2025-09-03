package main

import "fmt"

func main() {
	// var revenue int
	// var expenses int
	var revenue, expenses int
	var taxRate float64

	fmt.Print("請輸入營收: ")
	fmt.Scan(&revenue)
	fmt.Print("請輸入支出: ")
	fmt.Scan(&expenses)
	fmt.Print("請輸入稅率(%): ")
	fmt.Scan(&taxRate)

	earningsBeforeTax := revenue - expenses
	profit := float64(earningsBeforeTax) * (1 - taxRate/100) // 要記得轉型
	ratio := float64(earningsBeforeTax) / profit

	fmt.Println("稅前盈餘:", earningsBeforeTax)
	fmt.Println("稅後盈餘:", profit)
	fmt.Println("稅前盈餘與稅後盈餘的比率:", ratio)
}

/*
cd 01-go-essentials/02-exercise-solution
go run profit_calculator.go
*/
