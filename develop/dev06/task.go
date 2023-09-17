package main

import (
	"bufio"
	"flag"
	"fmt"
	"strconv"
	"strings"
)

/*
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

type flagSet struct {
	f string
	d string
	s bool
}

var fSet flagSet

// В Go заданнная функция init() выделяет элемент кода,
// который запускатся до любой другой части вашего пакета.
// Этот код запускается сразу же после импорта пакета
func init() {
	flag.StringVar(&fSet.f, "f", "", "fields")
	flag.StringVar(&fSet.d, "d", "\t", "delimiter")
	flag.BoolVar(&fSet.s, "s", false, "separated")

}

// f f+s, f+d, f+d+s
// else error
func IsFlagSet(flagName string) bool {
	isSet := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == flagName {
			isSet = true
		}
	})
	return isSet
}
func checkFlagF() []int {
	var ans []int
	if fSet.f != "" {
		f := strings.FieldsFunc(fSet.f, func(r rune) bool {
			return r == ' ' || r == ','
		})
		if len(f) == 0 {
			return nil
		} else {
			for _, s := range f {
				intVal, err := strconv.ParseInt(s, 10, 64)
				if err != nil {
					return nil
				}
				ans = append(ans, int(intVal)-1)
			}
		}

	}
	return ans
}
func myCut(scanner *bufio.Scanner) (string, error) {

	var cols []int = checkFlagF()
	if cols == nil {
		return "", fmt.Errorf("Please specify flags correctly, your flags are : %v", fSet)
	}
	ans := strings.Builder{}
	for scanner.Scan() {
		line := scanner.Text()
		tmpAns := strings.Split(line, fSet.d)
		if len(tmpAns) == 1 {
			if !fSet.s {
				//ans = append(ans, line)
				//fmt.Println(line)
				ans.WriteString(line)
			} else {

				ans.WriteString(" ")
			}
		} else {
			for _, ind := range cols {
				if ind < len(tmpAns) {
					ans.WriteString(tmpAns[ind])
					ans.WriteString(fSet.d)
				} else {
					break
				}
			}
			//ans = append(ans, res)
			//fmt.Println(res[:len(res)-1])

		}
		ans.WriteString("\n")
	}
	return strings.TrimSuffix(ans.String(), "\n"), nil
}

/*
func main() {
	flag.Parse()
	scanner := bufio.NewScanner(os.Stdin)
	err := myCut(scanner)
	if err != nil {
		fmt.Println(err)
	}
}*/
