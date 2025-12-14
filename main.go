package main

import (
	"fmt"

	"github.com/fatih/color"
)

func main() {
	fmt.Println("=== Go + Razd Example ===")
	
	// Используем библиотеку color для цветного вывода
	green := color.New(color.FgGreen).SprintFunc()
	cyan := color.New(color.FgCyan).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()
	
	fmt.Println(cyan("Привет из Go!"))
	fmt.Println(green("Проект успешно запущен"))
	fmt.Println(yellow("Это пример использования зависимости github.com/fatih/color"))
	
	fmt.Println()
	fmt.Println("=== Готово! ===")
}
