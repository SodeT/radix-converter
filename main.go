package main

import (
	"fmt"
	"strings"
	"math"
	"os"
	"strconv"
	"log"
)

var (
	chars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ+/"
	pl = fmt.Println
)

func main() {
	if len(os.Args) != 4 {
		pl(
`Usage: radix-convert <number> <from radix> <to radix>

Description:
	This command converts a number from any radix to any other radix.
		
Arguments:
	<number>            The number to be converted.
	<from radix>        The current radix (base) of the number.
	<to radix>          The target radix (base) to which the number will be converted.

Note:
	* This program only works for radices up to 64.
	* For base64, the output does not comply with RFC 4648 (RFC 4648 ordering [A-Z][a-z][0-9]+/)
	* Upper and lowercase letters are never considerd to be equal.
		`)
		return
	}

	number := os.Args[1]
	radix1, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatal("Invalid radix...")
	}

	radix2, _ := strconv.Atoi(os.Args[3])
	if err != nil {
		log.Fatal("Invalid radix...")
	}

	n := decode(number, radix1)
	pl(encode(n, radix2))
}

func decode(number string, radix int) float64 {
	strs := strings.Split(number, ".")

	totalInt := 0
	runeInts := []rune(strs[0])
	for i, j := 0, len(runeInts)-1; i < j; i, j = i+1, j-1 {
		runeInts[i], runeInts[j] = runeInts[j], runeInts[i]
	}

	for index, runeInt := range runeInts {
		value := strings.IndexRune(chars, runeInt)
		if value == -1 || value >= radix {
			log.Fatal("Number is not in specified radix...")
		}

		totalInt += int(math.Pow(float64(radix), float64(index))) * value
	}

	totalFraction := 0.0
	if len(strs) == 2 {
		runeFractions := []rune(strs[1]) 
		for i, j := 0, len(runeFractions)-1; i < j; i, j = i+1, j-1 {
			runeFractions[i], runeFractions[j] = runeFractions[j], runeFractions[i]
		}

		for index, runeFraction := range runeFractions {
			value := strings.IndexRune(chars, runeFraction)
			if value == -1 || value >= radix {
				log.Fatal("Number is not in specified radix...")
			}

			totalFraction += math.Pow(float64(radix), -float64(index + 1)) * float64(value)
		}
	}

	return float64(totalInt) + totalFraction
}


func encode(number float64, radix int) string {
	digits := []rune{}
	num := int(number)
	fractionalDigits := 0

	for number - float64(num) > 0 {
		number *= float64(radix)
		num = int(number)
		fractionalDigits++
	}
	 
	for num != 0 { 
		rest := num % radix
		digits = append(digits, rune(chars[rest]))
		num -= rest
		num /= radix

		if len(digits) == fractionalDigits {
			digits = append(digits, '.')
		}
	}

	for i, j := 0, len(digits)-1; i < j; i, j = i+1, j-1 {
		digits[i], digits[j] = digits[j], digits[i]
	}

	return string(digits)
}
