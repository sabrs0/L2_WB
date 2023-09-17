package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func getNumber(numStr string) (int64, error) {
	ans, err := strconv.ParseInt(numStr, 10, 32)
	return ans, err
}
func Unpack(str string) (string, error) {
	runes := []rune(str)

	ans := strings.Builder{}
	n := len(runes)
	cur := 0
	num := strings.Builder{}
	var runeTimes rune
	for cur < n {
		switch {
		//Если символ \ - то записываем то, что после него, при условии, что последний символ исходной строки не \
		case runes[cur] == '\\':
			if cur+1 < n {
				ans.WriteRune(runes[cur+1])
				cur += 2
			} else {
				return "", fmt.Errorf("Incorrect string")
			}
			continue
		//Если символ не цифра и не \ - добавляем в строку
		case (runes[cur] > '9' || runes[cur] < '0'):
			ans.WriteRune(runes[cur])
			cur++
			continue
		//Если символ не цифра и не \ - добавляем в строку
		case runes[cur] >= '0' && runes[cur] <= '9':
			if cur == 0 {
				return "", fmt.Errorf("Incorrect string")
			}
			num = strings.Builder{}
			runeTimes = runes[cur-1]
			for cur < n && runes[cur] >= '0' && runes[cur] <= '9' {
				num.WriteRune(runes[cur])
				cur++
			}
			times, err := getNumber(num.String())
			times -= 1
			if err != nil {
				return "", err
			}
			for i := 0; i < int(times); i++ {
				ans.WriteRune(runeTimes)
			}
			continue
		}

	}
	return ans.String(), nil
}

func main() {
	fmt.Println(Unpack("a4bc2d5e"))
}
