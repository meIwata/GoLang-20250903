package main

import (
	"fmt"
	"math"
)

func main() {
	// fmt.Println("Hello, World!")
	// fmt.Print(`Hello, World!`)

	var investmentAmount = 1000
	var expectedReturnRate = 5.5
	var years = 10

	var futureValue = float64(investmentAmount) * math.Pow(1+expectedReturnRate/100, float64(years))

	fmt.Println("未來價值:", futureValue)
	// 要執行的時候可以在 terminal 輸入 go run investment_calculator.go 或是 investment_calculator 或是 go run .
}

/*
go mod init（只需做一次），像是 go mod init example.com/first-app
go build
之後會產生一個exe檔案
回到terminal執行 ./ + exe檔案名稱，像 ./first-app
*/

// GO library: https://pkg.go.dev/std
