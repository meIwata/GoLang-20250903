package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	fmt.Print(`Hello, World!`)
}

/*
go mod init（只需做一次），像是 go mod init example.com/first-app
go build
之後會產生一個exe檔案
回到terminal執行 ./ + exe檔案名稱，像 ./first-app
*/

// GO library: https://pkg.go.dev/std
