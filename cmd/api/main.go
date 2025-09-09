package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("🎉 congrate!")
	fmt.Println("✨ From the Homebrew")

	fmt.Printf("Current time: %s\n", time.Now().Format("2006-01-02 15:04:05"))

	showMessage()
	calculate()
}

func showMessage() {
	fmt.Println("\n🔧 environment information:")
	fmt.Printf("Go version: %s\n", goVersion())
	fmt.Printf("Current dir: %s\n", getCurrentDir())
}

func calculate() {
	a, b := 15, 27
	sum := a + b
	fmt.Printf("\n🧮 Calculate demo: %d + %d = %d\n", a, b, sum)
}

func goVersion() string {
	return "Go 1.20.3" // Placeholder for actual Go version
}

func getCurrentDir() string {
	return "/home/user/project" // Placeholder for actual current directory
}
