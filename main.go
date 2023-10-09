package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	romanValBack := false
	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	expression := strings.Split(in.Text(), " ")
	if len(expression) != 3 {
		log.Fatal("Введено некорректное выражение")
	}
	firstNum, err := strconv.Atoi(expression[0])
	secNum, err2 := strconv.Atoi(expression[2])
	if err != nil && err2 == nil || err == nil && err2 != nil {
		fmt.Println("Несовпадение типов")
		return
	}
	if err != nil && err2 != nil {
		if expression[1] == "-" {
			log.Fatal("Римские цифры не могут быть отрицательными")
		}
		romanValBack = true
		firstNum = romanToArabic(expression[0])
		secNum = romanToArabic(expression[2])
	} else {
		if 0 >= firstNum || 10 < firstNum || 0 >= secNum || 10 < secNum {
			log.Fatal("Введеные числа должны быть в диапазоте от 1 до 10")
		}
	}
	var result int
	switch expression[1] {
	case "+":
		result = firstNum + secNum
	case "-":
		result = firstNum - secNum
	case "*":
		result = firstNum * secNum
	case "/":
		result = firstNum / secNum
	default:
		fmt.Println("Некорректный математический оператор:", expression[1])
		return
	}
	if romanValBack {
		fmt.Println(arabicToRoman(result))
	} else {
		fmt.Println(result)
	}
}

func romanToArabic(roman string) int {
	r2a := map[string]int{
		"X":    10,
		"IX":   9,
		"VIII": 8,
		"VII":  7,
		"VI":   6,
		"V":    5,
		"IV":   4,
		"III":  3,
		"II":   2,
		"I":    1,
	}
	arabic := r2a[strings.ToUpper(roman)]
	if arabic == 0 {
		log.Fatal(fmt.Sprintf("Введено некорректное чило: %s", roman))
	}
	return arabic
}

func arabicToRoman(num int) string {

	var romanNumeral string
	romanValues := []int{100, 90, 50, 40, 10, 9, 5, 4, 1}
	romanSymbols := []string{"C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	for i := 0; i < len(romanValues); i++ {
		for num >= romanValues[i] {
			romanNumeral += romanSymbols[i]
			num -= romanValues[i]
		}
	}

	return romanNumeral
}
