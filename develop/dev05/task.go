package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/
var flagf = flag.String("f", "data.txt", "filename")
var flagA = flag.Int("A", 0, "after +n strs")
var flagB = flag.Int("B", 0, "befpre +n strs")
var flagC = flag.Int("C", 0, "A+B")
var flagc = flag.Bool("c", false, "count strs")
var flagi = flag.Bool("i", false, "ignore case")
var flagv = flag.Bool("v", false, "invert")

var flagF = flag.Bool("F", false, "fixed, not pattern")
var flagn = flag.Bool("n", false, "line num")

type signFunc func(pattern string, str string) bool

func readFile(fname string) ([]string, error) {
	file, err := os.Open(fname)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	ans := []string{}

	for {
		str, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return nil, err
			}

		}
		ans = append(ans, str)
	}
	return ans, nil
}

func writeFile(fname string, strs []string) error {
	file, err := os.Create(fname)
	if err != nil {
		return err
	}
	defer file.Close()
	for _, s := range strs {
		_, err = file.WriteString(s)
		if err != nil {
			return err
		}
	}
	return nil
}

func recurAfterN(pattern string, strs []string, cur *int, end, curN, n int, curAns []string) []string {
	//fmt.Printf("Rec: curStr = %s, curN = %d, cur Cur = %d, cur ans = %v\n", strs[*cur], curN, *cur, curAns)
	if *cur >= end-1 {
		return curAns
	} else {
		if curN >= n+1 {
			*cur = *cur + 1
			return recurAfterN(pattern, strs, cur, end, 0, n, curAns)
		}
		curNN := curN
		if match, _ := regexp.MatchString(pattern, strs[*cur]); match {
			curNN = 1

		} else {
			if curNN == 0 {
				*cur = *cur + 1
				return recurAfterN(pattern, strs, cur, end, curNN, n, curAns)
			} else {
				curNN++
			}
		}
		*cur = *cur + 1
		return recurAfterN(pattern, strs, cur, end, curNN, n, append(curAns, strs[*cur-1]))
	}

}
func afterN(pattern string, strs []string, n int) []string {
	start := 0
	end := len(strs)
	return recurAfterN(pattern, strs, &start, end, 0, n, []string{})
}
func bCore(pattern string, strs []string, curAns []string, n int, curInd, lastMarked int) ([]string, int) {

	if match, _ := regexp.MatchString(pattern, strs[curInd]); match {
		//fmt.Println(lastMarked, curInd)
		if n > (curInd - lastMarked) {
			return append(curAns, strs[lastMarked:curInd+1]...), curInd + 1
		} else {
			return append(curAns, strs[curInd-n:curInd+1]...), curInd + 1
		}
	}
	return curAns, lastMarked
}
func beforeN(pattern string, strs []string, n int) []string {
	ans := []string{}
	start := 0
	end := len(strs)
	lastMarked := 0
	for start < end {
		ans, lastMarked = bCore(pattern, strs, ans, n, start, lastMarked)
		start++
	}
	return ans
}
func checkContextAfterN(pattern string, strs []string, curStart int) ([]string, int) {
	for i, str := range strs {
		if match, _ := regexp.MatchString(pattern, str); match {
			return []string{}, curStart + i + 1
		}
	}
	return strs, curStart + len(strs)
}
func contextN(pattern string, strs []string, n int) []string {
	ans := []string{}
	start := 0
	end := len(strs)
	lastMarked := 0
	for start < end {
		tmpLastMarked := lastMarked
		ans, lastMarked = bCore(pattern, strs, ans, n, start, lastMarked)
		if tmpLastMarked != lastMarked {
			strsToCheck := []string{}
			if start+n+1 > end-1 {
				strsToCheck = strs[start+1 : end]
			} else {
				strsToCheck = strs[start+1 : start+n+1]
			}
			var tmpAns []string
			tmpAns, start = checkContextAfterN(pattern, strsToCheck, start)
			fmt.Println(tmpAns)
			ans = append(ans, tmpAns...)
		} else {
			start++
		}

	}
	return ans

}
func countStrs(pattern string, strs []string) int {
	return len(grepping(pattern, strs, defaultSign))
}
func ignoreCaseSign(pattern string, str string) bool {
	match, _ := regexp.MatchString(strings.ToLower(pattern), strings.ToLower(str))
	return match
}
func invertedSign(pattern string, str string) bool {
	match, _ := regexp.MatchString(pattern, str)

	return !match
}
func fixedSign(pattern string, str string) bool {
	match := strings.Contains(str, pattern)

	return match
}
func defaultSign(pattern string, str string) bool {
	match, _ := regexp.MatchString(pattern, str)
	return match
}
func writeLines(pattern string, strs []string) []string {
	for i, str := range strs {
		strs[i] = strconv.Itoa(i) + "." + str
	}
	return strs
}
func grepping(pattern string, strs []string, f signFunc) []string {
	ans := []string{}
	for _, str := range strs {
		if f(pattern, str) {
			ans = append(ans, str)
		}
	}
	return ans
}

func myGrep() error {
	flag.Parse()
	fmt.Println("Usage: go run task.go [-<flag name>=<flag value>...] <pattern>")
	filename := flagf
	if len(os.Args) < 2 {
		return fmt.Errorf("At least 2 args. Your args = %d", len(os.Args))
	}

	pattern := os.Args[len(os.Args)-1]
	strs, err := readFile(*filename)
	if err != nil {
		return err
	}
	var ans []string
	var funcToGrep signFunc = defaultSign
	if *flagA > 0 {
		ans = afterN(pattern, strs, *flagA)
	} else if *flagB > 0 {
		ans = beforeN(pattern, strs, *flagB)
	} else if *flagC > 0 {
		ans = contextN(pattern, strs, *flagC)
	} else if *flagF {
		funcToGrep = fixedSign
	} else if *flagn {
		ans = writeLines(pattern, strs)
	} else if *flagi {
		funcToGrep = ignoreCaseSign
	} else if *flagv {
		funcToGrep = invertedSign
	} else if *flagc {
		ans = append(ans, strconv.Itoa(countStrs(pattern, strs)))
	}
	if len(ans) == 0 {
		ans = grepping(pattern, strs, funcToGrep)
	}
	return writeFile("grepped"+*filename, ans)
}

func main() {
	err := myGrep()
	if err != nil {
		fmt.Println(err)
	}

}
