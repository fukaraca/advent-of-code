package main

import (
	"fmt"
	"os"
	"strconv"
)

var ln int

func main() {
	file, err := os.ReadFile("2023/day1/d1data")
	if err != nil {
		panic(err)
	}
	sumDigits := calculateByDigitsOnly(file)
	sumDigitsAndLetters := calculateByDigitsAndLetters(file)
	fmt.Println("Sum of all calibration values")
	fmt.Println("Calculated by digits only:", sumDigits)
	fmt.Println("Calculated by digits and letters:", sumDigitsAndLetters)

}

func calculateByDigitsOnly(file []byte) int {
	var sum int
	var first, last string
	for _, b := range file {
		if b == 10 {
			sum += getCalibrationValue(first, last)
			first, last = "", ""
		}
		if b > 47 && b < 58 {
			if first == "" {
				first = string(b)
				continue
			}
			last = string(b)
		}
	}
	sum += getCalibrationValue(first, last)
	return sum
}

func calculateByDigitsAndLetters(file []byte) int {
	var sum int
	var first, last, text string
	for _, b := range file {
		if b == 10 { // linebreak
			sum += getCalibrationValue(l2d(first), l2d(last))
			first, last, text = "", "", ""
		} else if b > 47 && b < 58 { // digit
			text = ""
			if first == "" {
				first = string(b)
				continue
			}
			last = string(b)
		} else { // letters
			text += string(b)
			switch l := len(text); {
			case l < 3:
			case l > 5:
				text = ""
			default:
				if i := l2d(text); len(i) > 0 {
					if first == "" {
						first = i
					} else {
						last = i
					}
					text = text[len(text)-1:]
					continue
				}
				if !digitMap[text] {
					text = text[1:]
				}
			}
		}
	}
	sum += getCalibrationValue(l2d(first), l2d(last))
	return sum
}

func getCalibrationValue(first, last string) int {
	ln++
	if len(last) == 0 {
		first += first
	} else {
		first += last
	}
	calibrationValue, err := strconv.Atoi(first)
	if err != nil {
		panic(err)
	}
	return calibrationValue
}

var m = map[string]string{"one": "1", "two": "2", "three": "3", "four": "4", "five": "5", "six": "6", "seven": "7", "eight": "8", "nine": "9"}

func l2d(s string) string {
	if len(s) <= 1 {
		return s
	}
	return m[s]
}

var digitMap = map[string]bool{
	"thr":  true,
	"thre": true,
	"fou":  true,
	"fiv":  true,
	"sev":  true,
	"seve": true,
	"eig":  true,
	"eigh": true,
	"nin":  true,
}
