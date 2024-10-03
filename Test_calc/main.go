package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Введите значение:")
	reader := bufio.NewReader(os.Stdin)
	for {
		console, _ := reader.ReadString('\n')
		s := strings.ReplaceAll(console, " ", "")
		other(strings.ToUpper(strings.TrimSpace(s)))
	}
}

var roman = map[string]int{
	"C":    100,
	"XC":   90,
	"L":    50,
	"XL":   40,
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
var convIntToRoman = [14]int{
	100,
	90,
	50,
	40,
	10,
	9,
	8,
	7,
	6,
	5,
	4,
	3,
	2,
	1,
}

var a, b *int

var operation = map[string]func() int{
	"+": func() int { return *a + *b },
	"-": func() int { return *a - *b },
	"/": func() int { return *a / *b },
	"*": func() int { return *a * *b },
}
var data []string

func other(s string) {
	var oper string
	var stringsFound int
	numbers := make([]int, 0)
	romans := make([]string, 0)
	romansToInt := make([]int, 0)
	for idx := range operation {
		for _, val := range s {
			if idx == string(val) {
				oper += idx
				data = strings.Split(s, oper)
			}
		}
	}
	switch {
	case len(oper) > 1:
		panic("Вывод ошибки, так как формат математической операции " +
			"не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
	case len(oper) < 1:
		panic("Вывод ошибки, так как строка " +
			"не является математической операцией.")
	}
	for _, elem := range data {
		num, err := strconv.Atoi(elem)
		if err != nil {
			stringsFound++
			romans = append(romans, elem)
		} else {
			numbers = append(numbers, num)
		}
	}

	switch stringsFound {
	case 1:
		panic("Вывод ошибки, так как используются " +
			"одновременно разные системы счисления.")
	case 0:
		errCheck := numbers[0] > 0 && numbers[0] < 11 &&
			numbers[1] > 0 && numbers[1] < 11
		if val, ok := operation[oper]; ok && errCheck == true {
			a, b = &numbers[0], &numbers[1]
			fmt.Println(val())
		} else {
			panic("Калькулятор умеет работать только с арабскими целыми " +
				"числами или римскими цифрами от 1 до 10 включительно")
		}
	case 2:
		for _, elem := range romans {
			if val, ok := roman[elem]; ok && val > 0 && val < 11 {
				romansToInt = append(romansToInt, val)
			} else {
				panic("Калькулятор умеет работать только с арабскими целыми " +
					"числами или римскими цифрами от 1 до 10 включительно")
			}
		}
		if val, ok := operation[oper]; ok {
			a, b = &romansToInt[0], &romansToInt[1]
			intToRoman(val())
		}
	}
}
func intToRoman(romanResult int) {
	var romanNum string
	if romanResult == 0 {
		panic("Вывод ошибки, так как в римской системе нет числа 0.")
	} else if romanResult < 0 {
		panic("Вывод ошибки, так как в римской системе " +
			"нет отрицательных чисел.")
	}
	for romanResult > 0 {
		for _, elem := range convIntToRoman {
			for i := elem; i <= romanResult; {
				for index, value := range roman {
					if value == elem {
						romanNum += index
						romanResult -= elem
					}
				}
			}
		}
	}
	fmt.Println(romanNum)
}
