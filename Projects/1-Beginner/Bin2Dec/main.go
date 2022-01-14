package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var binaryNumber string
	for {
		fmt.Print("Enter your binary number (up to 8 digits): ")
		_, err := fmt.Scanf("%s", &binaryNumber)
		if err != nil {
			fmt.Println("Error: ", err)
			continue
		}
		isValid, msgReturn := isValidNumber(binaryNumber)
		if !isValid {
			fmt.Println(msgReturn)
			continue
		}
		fmt.Println("Equivalent decimal number:", convertBinaryToDecimal(binaryNumber))
	}
}

func isValidNumber(binaryNumber string) (bool, string) {
	isValid := true
	var msgReturn []string

	if len(binaryNumber) > 8 {
		isValid = false
		msgReturn = append(msgReturn, "Invalid size number. Must be up to 8 digits.")
	}
	if !isValidBinaryNumber(binaryNumber) {
		isValid = false
		msgReturn = append(msgReturn, "Invalid binary number. Must contain only 0's and 1's.")
	}
	return isValid, strings.Join(msgReturn, "\n")
}

func isValidBinaryNumber(s string) bool {
	for _, i := range s {
		if i == '1' || i == '0' {
			continue
		}
		return false
	}
	return true
}

func convertBinaryToDecimal(s string) uint8 {
	var subTotal uint8 = 0
	for _, i := range s {
		j, err := strconv.ParseUint(string(i), 10, 8)
		if err != nil {
			fmt.Println("Error parsing digit: ", err)
		}
		subTotal = subTotal*2 + uint8(j)
	}
	return subTotal
}
