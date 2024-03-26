package main

import (
	"fmt"
	"strings"
	"math"
)

var (
	chars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ+/"
	pl = fmt.Println
)

func main() {
	pl(decode("fff", 16))
	pl(encode(23, 16))
}

func decode(number string, radix int) int {
	totalValue := 0
	digits := []rune(number)

	for i, j := 0, len(digits)-1; i < j; i, j = i+1, j-1 {
		digits[i], digits[j] = digits[j], digits[i]
	}

	for index, digit := range digits {
		value := strings.IndexRune(chars, digit)
		totalValue += int(math.Pow(float64(radix), float64(index))) * value
	}
	return totalValue
}

func encode(number int, radix int) string {
	digits := []rune{}
	for number != 0 { 
		rest := number % radix
		digits = append(digits, rune(chars[rest]))
		number -= rest
		number /= radix 
	}

	for i, j := 0, len(digits)-1; i < j; i, j = i+1, j-1 {
		digits[i], digits[j] = digits[j], digits[i]
	}

	return string(digits)
}