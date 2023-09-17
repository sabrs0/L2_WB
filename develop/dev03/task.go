package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

# Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

# Дополнительное

# Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

var flagF = flag.String("f", "data0.txt", "filename")
var flagK = flag.Int("k", 0, "set column number")
var flagN = flag.Bool("n", false, "sort by numeric value")
var flagR = flag.Bool("r", false, "sort reversed")
var flagU = flag.Bool("u", false, "unique strings only")

func checkNum(str string) (float64, error) {
	sourceStr := strings.TrimSpace(str)
	var strNumber string
	for i := 0; i < len(sourceStr); i++ {
		if sourceStr[i] != '.' && sourceStr[i] != '-' &&
			(sourceStr[i] < '0' || sourceStr[i] > '9') {
			break
		}
		strNumber += string(sourceStr[i])
	}
	ans, err := strconv.ParseFloat(strNumber, 64)
	if err != nil {
		return 0, err
	}
	return ans, nil
}
func strsTostr(strs []string) string {
	s := strings.Builder{}
	for _, str := range strs {
		s.WriteString(str)
	}
	return s.String()
}

func getStrsFromMap(keys []string, mapa map[string]string) []string {
	ans := []string{}
	for _, k := range keys {
		ans = append(ans, mapa[k])
	}
	return ans
}
func getStrsFromMapArr(keys []string, mapa map[string][]string) []string {
	ans := []string{}
	for _, k := range keys {
		ans = append(ans, mapa[k]...)
	}
	return ans
}
func sortedNumsByMap(numStrings map[float64][]string) []string {
	keys := make([]float64, len(numStrings))
	for k := range numStrings {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })
	ans := []string{}
	for _, k := range keys {
		ans = append(ans, numStrings[k]...)
	}
	return ans
}
func sortedStringsByMap(numStrings map[string][]string) []string {
	keys := make([]string, len(numStrings))
	for k := range numStrings {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return getStrsFromMapArr(keys, numStrings)

}

// -k
func getSplittedStr(str string) []string {
	splitted := strings.Fields(str)
	/*ans := []string{}
	for _, s := range splitted {

		if s != "\t" {
			ans = append(ans, s)
		}
	}*/
	return splitted //ans
}

func getStrByColumn(str []string, col int) (map[string][]string, []string) {
	colStrings := make(map[string][]string)
	nonColStrings := []string{}
	for _, s := range str {

		splitted := getSplittedStr(s)
		//fmt.Println(splitted)
		if len(splitted) > col {
			colStrings[splitted[col-1]] = append(colStrings[splitted[col-1]], s)
		} else {
			nonColStrings = append(nonColStrings, s)

		}
	}
	return colStrings, nonColStrings
}
func sortByColumn(strs []string, column int) []string {
	colStrs, nonColStrs := getStrByColumn(strs, column)
	res := []string{}
	sort.Strings(nonColStrs)
	sortedColStrs := sortedStringsByMap(colStrs)
	res = append(res, nonColStrs...)
	res = append(res, sortedColStrs...)
	return res

}

// -n- делим на 2 части: где есть числовой и где нет.
// Сначала в лексикографическом порядке то, где нет,
// потом в числовом порядке, где да.
func numericSort(strs []string) []string {
	numStrings := make(map[float64][]string)
	notNumStrings := []string{}
	for _, s := range strs {
		num, err := checkNum(s)
		if err != nil {
			notNumStrings = append(notNumStrings, s)
		} else {
			numStrings[num] = append(numStrings[num], s)
		}
	}
	sort.Strings(notNumStrings)
	//линуксовый sort считает что числовое значение строки, у которой нет в нормальном виде числа - 1
	numStrings[1] = append(numStrings[1], notNumStrings...)
	res := []string{}
	nums := sortedNumsByMap(numStrings)
	res = append(res, nums...)
	return res
}

// -r
func reverseStrings(str []string) {
	n := len(str)
	for i := 0; i < n/2; i++ {
		str[i], str[n-1-i] = str[n-1-i], str[i]
	}
}

// -u
func removeDuplicates(str []string) []string {
	uniques := []string{}
	uniquesMap := make(map[string]struct{})
	for _, s := range str {
		if _, isSet := uniquesMap[s]; !isSet {
			uniquesMap[s] = struct{}{}
			uniques = append(uniques, s)
		}
	}
	return uniques
}

func readFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("Cant open file: %s", err)
	}
	defer file.Close()
	ans := []string{}
	reader := bufio.NewReader(file)
	for true {
		str, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return nil, fmt.Errorf("Cant read file: %s", err)
			}
		}
		ans = append(ans, str)
	}
	return ans, nil
}
func writeFile(filename string, strs []string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("Cant open file: %s", err)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	for _, s := range strs {
		_, err := writer.WriteString(s)
		if err != nil {
			return fmt.Errorf("Cant write in  file: %s", err)
		}
	}
	writer.Flush()
	return nil
}

// Приводим исходные строки к виду, который нужен для сортировки по логике linux sort
func makeSourceStringsMap(sourceStrs []string) map[string][]string {
	mapa := make(map[string][]string, len(sourceStrs))
	for _, str := range sourceStrs {
		unSpacedString := strings.TrimSpace(str)
		mapa[strings.ToLower(unSpacedString)] = append(mapa[strings.ToLower(unSpacedString)], str)
	}
	for k := range mapa {
		sort.Slice(mapa[k], func(i, j int) bool {
			return mapa[k][i][0] > mapa[k][j][0]
		})
	}
	return mapa

}
func getStringsByKeys(mapa map[string][]string) []string {
	ans := []string{}
	for k := range mapa {
		ans = append(ans, k)
	}
	return ans
}
func mySort() error {

	flag.Parse()
	filename := *flagF
	strs, err := readFile("unsorted/" + filename)
	if err != nil {
		return err
	}
	strsMap := makeSourceStringsMap(strs)
	strsToSort := getStringsByKeys(strsMap)
	res := make([]string, len(strs))
	if *flagK > 0 {
		sortedStrs := sortByColumn(strsToSort, *flagK)
		copy(res, getStrsFromMapArr(sortedStrs, strsMap))
	} else {
		if *flagN {
			sortedStrs := numericSort(strsToSort)
			copy(res, getStrsFromMapArr(sortedStrs, strsMap))
		} else {
			sort.Strings(strsToSort)
			copy(res, getStrsFromMapArr(strsToSort, strsMap))
		}
	}

	if *flagR {
		reverseStrings(res)
	}
	if *flagU {
		res = removeDuplicates(res)
	}
	//fmt.Println(res)
	err = writeFile("sorted/"+filename, res)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	err := mySort()
	if err != nil {
		panic(err)
	}
}
