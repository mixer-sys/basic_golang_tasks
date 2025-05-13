package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Пожалуйста, введите число в качестве аргумента.")
		return
	}
	celcius_temp, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fmt.Println("Ошибка: аргумент должен быть числом.")
		return
	}
	farienheit_temp := float64(celcius_temp*9/5) + 32
	fmt.Printf("%.0f°C = %.2f°F \n", celcius_temp, farienheit_temp)
}
