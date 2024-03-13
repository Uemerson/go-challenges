package main

import "fmt"

func intToRoman(num int) string {
	r := ""
	tmp := num

	for tmp >= 1000 {
		tmp -= 1000
		r += "M"
	}

	for tmp >= 500 {
		if tmp >= 900 {
			tmp -= 900
			r += "CM"
		} else {
			tmp -= 500
			r += "D"
		}
	}

	for tmp >= 100 {
		if tmp >= 400 {
			tmp -= 400
			r += "CD"
		} else {
			tmp -= 100
			r += "C"
		}
	}

	for tmp >= 50 {
		if tmp >= 90 {
			tmp -= 90
			r += "XC"
		} else {
			tmp -= 50
			r += "L"
		}
	}

	for tmp >= 10 {
		if tmp >= 40 {
			tmp -= 40
			r += "XL"
		} else if tmp == 9 {
			tmp -= 9
			r += "IX"
		} else {
			tmp -= 10
			r += "X"
		}
	}

	for tmp > 0 {
		if tmp >= 9 {
			tmp -= 9
			r += "IX"
		} else if tmp >= 5 {
			tmp -= 5
			r += "V"
		} else if tmp >= 4 {
			tmp -= 4
			r += "IV"
		} else {
			tmp -= 1
			r += "I"
		}
	}

	return r
}

func main() {
	fmt.Println(intToRoman(3) == "III")
	fmt.Println(intToRoman(58) == "LVIII")
	fmt.Println(intToRoman(1994) == "MCMXCIV")
	fmt.Println(intToRoman(400) == "CD")
	fmt.Println(intToRoman(40) == "XL")
}
