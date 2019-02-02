package main

import "fmt"

func main() {
	a := 1
	fmt.Printf(HelloWorld("appleboy"))
	fmt.Println("一天就學會 Go 語言")

	if a >= 1 {
		fmt.Println("a >= 1")
	}
}

// HelloWorld ..
func HelloWorld(userName string) string {
	return fmt.Sprintf("Hi, %s ", userName)
}

// formatting
// gofmt -d filepath : display diff
// gofmt -w filepath : write result to file

// linting
// gotlint <package/directory>
