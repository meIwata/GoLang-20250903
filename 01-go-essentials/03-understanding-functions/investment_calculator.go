package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
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

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

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
