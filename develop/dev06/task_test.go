package main

import (
	"bufio"
	"fmt"
	"strings"
	"testing"
)

var flags []flagSet = []flagSet{
	{"1,2,3", " ", false},
	{"1,2,3", " ", true},
}

var toCheckStrs []string = []string{
	"hello i am",
	"hello,i,am",
	"hello, i, am, boy",
}
var expectedStrsFt []string = []string{
	"hello i am ",
	"hello,i,am",
	"hello, i, am, ",
}
var expectedStrsSc []string = []string{
	"hello i am ",
	" ",
	"hello, i, am, ",
}

func makeCmd(curFlags flagSet, strToCheck, expected string) error {
	fSet = curFlags
	strOutPut, err := myCut(bufio.NewScanner(strings.NewReader(strToCheck)))
	if err != nil {
		return err
	}

	if strOutPut != expected {
		return fmt.Errorf("Unmatch:\nExpected - %s(%d)\nRecieved - %s(%d)", expected, len(expected), strOutPut, len(strOutPut))
	}
	return nil

}

func TestCut(t *testing.T) {
	for i, curFlags := range flags {
		var expectedArr []string
		if i == 0 {
			expectedArr = expectedStrsFt
		} else {
			expectedArr = expectedStrsSc
		}
		for i, s := range toCheckStrs {
			err := makeCmd(curFlags, s, expectedArr[i])
			if err != nil {
				t.Errorf(err.Error())
			}
		}
	}

}
