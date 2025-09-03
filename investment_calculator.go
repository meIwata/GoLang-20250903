package main

import (
	"fmt"
	"math"
)

func main() {
	// fmt.Println("Hello, World!")
	// fmt.Print(`Hello, World!`)

	// var investmentAmount float64 = 1000
	// var investmentAmount, years float64 = 1000, 10
	// investmentAmount, years := 1000.0, 10.0

	const inflationRate float64 = 2.5
	var investmentAmount float64 = 1000
	var years float64 = 10

	// var expectedReturnRate float64 = 5.5
	expectedReturnRate := 5.5 // 也可以寫成這樣

	// var years float64 = 10

	// var futureValue = float64(investmentAmount) * math.Pow(1+expectedReturnRate/100, float64(years))
	// var futureValue = investmentAmount * math.Pow(1+expectedReturnRate/100, float64(years))
	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println("未來預估價值:", futureValue)
	fmt.Println("未來實際價值:", futureRealValue)
	// 要執行的時候可以在 terminal 輸入 go run investment_calculator.go 或是 investment_calculator 或是 go run .
}

/*
go mod init（只需做一次），像是 go mod init example.com/first-app
go build
之後會產生一個exe檔案
回到terminal執行 ./ + exe檔案名稱，像 ./first-app
*/

// GO library: https://pkg.go.dev/std
