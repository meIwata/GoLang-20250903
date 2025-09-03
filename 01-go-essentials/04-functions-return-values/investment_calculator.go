package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {
	var investmentAmount float64
	var years float64
	expectedReturnRate := 5.5

	// fmt.Print("投資金額: ")
	outputText("投資金額: ")
	fmt.Scan(&investmentAmount)

	// fmt.Print("預期年化報酬率(%): ")
	outputText("預期年化報酬率(%): ")
	fmt.Scan(&expectedReturnRate)

	// fmt.Print("投資年限: ")
	outputText("投資年限: ")
	fmt.Scan(&years)

	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)

	//futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	//futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	formattedFV := fmt.Sprintf("未來預估價值: %.1f\n", futureValue)
	formattedRFV := fmt.Sprintf("未來實際價值: %.1f\n", futureRealValue)
	// outputs information
	// fmt.Println("未來預估價值:", futureValue)
	// fmt.Printf(`未來預估價值: %.1f\n

	// Future Value (adjusted for Inflation): %.1f`, futureValue, futureRealValue)
	// fmt.Println("未來預估價值 (adjusted for Inflation):", futureRealValue)
	fmt.Print(formattedFV, formattedRFV)
}

func outputText(text string) {
	fmt.Print(text)
}

// func outputText(text string, text2 string) {
// 	fmt.Print(text)
// 	fmt.Print(text2)
// }

func calculateFutureValues(investmentAmount float64, expectedReturnRate float64, years float64) (float64, float64) { // 這裡的(float64, float64)代表回傳兩個float64 的futureValue, realFutureValue值
	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	realFutureValue := futureValue / math.Pow(1+inflationRate/100, years)

	return futureValue, realFutureValue
}
