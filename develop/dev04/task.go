package main

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
)

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func isAnagram(str, isAnagramStr []rune) bool {
	sort.Slice(str, func(i, j int) bool {
		return str[i] < str[j]
	})
	sort.Slice(isAnagramStr, func(i, j int) bool {
		return isAnagramStr[i] < isAnagramStr[j]
	})
	return reflect.DeepEqual(str, isAnagramStr)
}

func deleteElem(s []string, ind int) {
	copy(s[ind:], s[ind+1:])
	s[len(s)-1] = ""
	s = s[:len(s)-1]
}

func checkAnagrams(toCheck string, resSlice []string) ([]string, []string) {
	res := []string{}
	count := 0
	for i, s := range resSlice {
		if isAnagram([]rune(s), []rune(toCheck)) {
			res = append(res, s)
			resSlice[i], resSlice[count] = resSlice[count], resSlice[i]
			count++
		}
	}
	return res, resSlice[count:]
}

func anagramSearch(data []string) map[string][]string {
	for i := range data {
		data[i] = strings.ToLower(data[i])
	}
	res := make(map[string][]string)
	for {
		if len(data) == 0 {
			break
		}
		toCheck := data[0]
		data = data[1:]
		var ans []string
		ans, data = checkAnagrams(toCheck, data)
		if len(ans) != 0 {
			sort.Strings(ans)
			res[toCheck] = ans
		}
	}
	return res
}
func main() {
	strs := []string{
		"пятак", "пятка", "тяпка",
		"листок", "слиток", "столик",
	}
	fmt.Println(anagramSearch(strs))
}
