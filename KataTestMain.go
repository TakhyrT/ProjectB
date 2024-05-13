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
	panic("ошибка: неправильное римское число")
}

func ArabicToRoman(arabic int) (string, error) {
	arabicToRoman := map[int]string{
		1:   "I",
		2:   "II",
		3:   "III",
		4:   "IV",
		5:   "V",
		6:   "VI",
		7:   "VII",
		8:   "VIII",
		9:   "IX",
		10:  "X",
		11:  "XI",
		12:  "XII",
		13:  "XIII",
		14:  "XIV",
		15:  "XV",
		16:  "XVI",
		17:  "XVII",
		18:  "XVIII",
		19:  "XIX",
		20:  "XX",
		21:  "XXI",
		22:  "XXII",
		23:  "XXIII",
		24:  "XXIV",
		25:  "XXV",
		26:  "XXVI",
		27:  "XXVII",
		28:  "XXVIII",
		29:  "XXIX",
		30:  "XXX",
		31:  "XXXI",
		32:  "XXXII",
		33:  "XXXIII",
		34:  "XXXIV",
		35:  "XXXV",
		36:  "XXXVI",
		37:  "XXXVII",
		38:  "XXXVIII",
		39:  "XXXIX",
		40:  "XL",
		41:  "XLI",
		42:  "XLII",
		43:  "XLIII",
		44:  "XLIV",
		45:  "XLV",
		46:  "XLVI",
		47:  "XLVII",
		48:  "XLVIII",
		49:  "XLIX",
		50:  "L",
		51:  "LI",
		52:  "LII",
		53:  "LIII",
		54:  "LIV",
		55:  "LV",
		56:  "LVI",
		57:  "LVII",
		58:  "LVIII",
		59:  "LIX",
		60:  "LX",
		61:  "LXI",
		62:  "LXII",
		63:  "LXIII",
		64:  "LXIV",
		65:  "LXV",
		66:  "LXVI",
		67:  "LXVII",
		68:  "LXVIII",
		69:  "LXIX",
		70:  "LXX",
		71:  "LXXI",
		72:  "LXXII",
		73:  "LXXIII",
		74:  "LXXIV",
		75:  "LXXV",
		76:  "LXXVI",
		77:  "LXXVII",
		78:  "LXXVIII",
		79:  "LXXIX",
		80:  "LXXX",
		81:  "LXXXI",
		82:  "LXXXII",
		83:  "LXXXIII",
		84:  "LXXXIV",
		85:  "LXXXV",
		86:  "LXXXVI",
		87:  "LXXXVII",
		88:  "LXXXVIII",
		89:  "LXXXIX",
		90:  "XC",
		91:  "XCI",
		92:  "XCII",
		93:  "XCIII",
		94:  "XCIV",
		95:  "XCV",
		96:  "XCVI",
		97:  "XCVII",
		98:  "XCVIII",
		99:  "XCIX",
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

	operatorPos := -1

	for i, r := range input {
		if strings.ContainsRune(operators, r) {
			operator = r
			operatorPos = i
			break
		}
	}

	if operatorPos == -1 {
		panic("не найден оператор")
	}

	isRoman := strings.ContainsAny(input, "IVXLCDM")

	if isRoman {
		leftIsRoman := strings.ContainsAny(input[:operatorPos], "IVXLCDM")
		rightIsRoman := strings.ContainsAny(input[operatorPos+1:], "IVXLCDM")
		if !leftIsRoman || !rightIsRoman {
			panic("Нельзя использовать арабские и римские числа вместе")
		}
	} else {
		leftIsRoman := strings.ContainsAny(input[:operatorPos], "IVXLCDM")
		rightIsRoman := strings.ContainsAny(input[operatorPos+1:], "IVXLCDM")
		if leftIsRoman || rightIsRoman {
			panic("Нельзя использовать арабские и римские числа вместе")
		}
	}

	var err error
	num1, err = ParseNumber(input[:operatorPos])
	if err != nil {
		panic(err)
	}

	num2, err = ParseNumber(input[operatorPos+1:])
	if err != nil {
		panic(err)
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
		panic(err)
	}
	if result < 1 || result > 10 {
		panic(fmt.Sprintf("число %d вне допустимого диапазона 1-10", result))
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
		panic(err)
	}
	result, err := Calculate(num1, num2, string(operator))
	if err != nil {
		panic(err)
	}
	isRoman := strings.ContainsAny(input, "IVXLCDM")
	if isRoman {
		romanResult, err := ArabicToRoman(result)
		if err != nil {
			panic(err)
		}
		fmt.Println("Результат: ", romanResult)
	} else {
		fmt.Println("Результат: ", result)
	}
}
