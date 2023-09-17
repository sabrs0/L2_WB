package main

import (
	"bytes"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	"testing"
)

var goCmdArgs []string = []string{
	"-f=data2.txt -A=10 Mas",
	"-f=data2.txt -B=10 Mas",
	"-f=data2.txt -C=10 Yo",
	"-f=data2.txt -c Yo",
	"-f=data2.txt -F H",
	"-f=data2.txt -v H",
	"-f=data2.txt -n H",
	"-f=data2.txt -i w",
}

func cmpRes(fname1, fname2 string) bool {
	file1, _ := os.Open("fname1")
	defer file1.Close()
	file2, _ := os.Open("fname1")
	defer file2.Close()

	content1, _ := ioutil.ReadAll(file1)
	content2, _ := ioutil.ReadAll(file1)

	return bytes.Equal(content1, content2)
}

func TestGrep(t *testing.T) {
	for i := 0; i < len(goCmdArgs); i++ {
		cmd := exec.Command("go", "run", "task.go", goCmdArgs[i])
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		err := cmd.Run()
		if err != nil {
			t.Errorf("Some error occured while running cmd : %s", err)
		}
		if !cmpRes("greppeddata0.txt", "expected/f"+strconv.Itoa(i+1)+".txt") {
			t.Errorf("Error with test %d", i+1)
		}
	}
}
