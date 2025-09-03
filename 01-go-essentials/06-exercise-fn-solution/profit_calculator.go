package main

import "fmt"

func main() {
	revenue := getUserInput("收入: ")
	expenses := getUserInput("支出: ")
	taxRate := getUserInput("稅率: ")

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	//fmt.Println(ebt)
	//fmt.Println(profit)
	//fmt.Println(ratio)
	fmt.Printf("稅前淨利: %v\n", ebt)
	fmt.Printf("稅後淨利: %v\n", profit)
	fmt.Printf("稅前淨利與稅後淨利比率: %.2f\n", ratio)
}

func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}

func getUserInput(infoText string) float64 {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)

	return userInput
}
