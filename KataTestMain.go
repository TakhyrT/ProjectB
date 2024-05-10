package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func RomanToArabic(roman string) (int, error) {
	romanNumerals := map[string]int{
		"I":    1,
		"II":   2,
		"III":  3,
		"IV":   4,
		"V":    5,
		"VI":   6,
		"VII":  7,
		"VIII": 8,
		"IX":   9,
		"X":    10,
	}

	if value, exists := romanNumerals[roman]; exists {
		return value, nil
	}
	return 0, fmt.Errorf("ошибка")
}

    func ArabicToRoman(arabic int) (string, error) {
    arabicToRoman := map[int]string{
        1:  "I",
        2:  "II",
        3:  "III",
        4:  "IV",
        5:  "V",
        6:  "VI",
        7:  "VII",
        8:  "VIII",
        9:  "IX",
        10: "X",
          11: "XI",
        12: "XII",
        13: "XIII",
        14: "XIV",
        15: "XV",
        16: "XVI",
        17: "XVII",
        18: "XVIII",
        19: "XIX",
        20: "XX",
          30: "XXX",
          40: "XL",
          50: "L",
          100: "C",
    }

    if roman, ok := arabicToRoman[arabic]; ok {
        return roman, nil
    }
    panic(fmt.Sprintf("ошибка поведения для числа %d", arabic))
}
func Calculate(num1, num2 int, operator string) (int, error) {
	switch operator {
	case "+":
		return num1 + num2, nil
	case "-":
		return num1 - num2, nil
	case "*":
		return num1 * num2, nil
	case "/":
		if num2 == 0 {
			return 0, fmt.Errorf("деление на ноль")
		}
		return num1 / num2, nil
	default:
		return 0, fmt.Errorf("неизвестный оператор")
	}
}

func ParseLine(input string) (int, int, rune, error) {
    var num1, num2 int
    var operator rune
    operators := "+-*/"
    isRoman := strings.ContainsAny(input, "IVXLCDM")

    for i, r := range input {
        if strings.ContainsRune(operators, r) {
            operator = r

            if isRoman != strings.ContainsAny(input[:i], "IVXLCDM") {
                panic("Нельзя использовать арабские и римские числа вместе")
            }

            var err error
            num1, err = ParseNumber(input[:i])
            if err != nil {
                return 0, 0, 0, err
            }

            num2, err = ParseNumber(input[i+1:])
            if err != nil {
                return 0, 0, 0, err
            }
            break
        }
    }

    if operator == 0 {
        panic("не найден оператор")
    }

    return num1, num2, operator, nil
    
	if operator == 0 {
		return 0, 0, 0, fmt.Errorf("не найден оператор")
	}

	return num1, num2, operator, nil
}
func ParseNumber(numberString string) (int, error) {
	if strings.ContainsAny(numberString, "IVXLCDM") {
		return RomanToArabic(numberString)
	}
	return StringToArabic(numberString)
}

func StringToArabic(arabic string) (int, error) {
	result, err := strconv.Atoi(arabic)
	if err != nil {
		return 0, err
	}
	if result < 1 || result > 10 {
		return 0, fmt.Errorf("число %d вне допустимого диапазона 1-10", result)
	}

	return result, nil 
}
func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите выражение:")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	num1, num2, operator, err := ParseLine(input)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	result, err := Calculate(num1, num2, string(operator))
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	isRoman := strings.ContainsAny(input, "IVXLCDM")
	if isRoman {
		romanResult, err := ArabicToRoman(result)
		if err != nil {
    panic(err)
			return
		}
		fmt.Println("Результат: ", romanResult)
	} else {
		fmt.Println("Результат: ", result)
	}
}
