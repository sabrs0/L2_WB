package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"testing"
)

var file = "data0.txt"

var goCmdArgs = [][]string{
	[]string{"run", "task.go"},
	[]string{"run", "task.go", "-k=2"},
	[]string{"run", "task.go", "-n"},
	[]string{"run", "task.go", "-r"},
	[]string{"run", "task.go", "-u"},
}

/*[]string{
	"go run task.go",
	"go run task.go -k=2",
	"go run task.go -n",
	"go run task.go -r",
	"go run task.go -u",
}*/

var osCmdArgs = [][]string{
	[]string{"unsorted/data0.txt"},
	[]string{"unsorted/data0.txt", "-k", "2"},
	[]string{"unsorted/data0.txt", "-n"},
	[]string{"unsorted/data0.txt", "-r"},
	[]string{"unsorted/data0.txt", "-u"},
} /*[]string{
	"sort ./unsorted/data0.txt > ./expected/data0.txt",
	"sort -k 2 ./unsorted/data0.txt > ./expected/data0.txt",
	"sort -n ./unsorted/data0.txt > ./expected/data0.txt",
	"sort -r ./unsorted/data0.txt > ./expected/data0.txt",
	"sort -u ./unsorted/data0.txt > ./expected/data0.txt",
}*/

func cmpRes(fname1, fname2 string) bool {
	file1, _ := os.Open("fname1")
	defer file1.Close()
	file2, _ := os.Open("fname1")
	defer file2.Close()

	content1, _ := ioutil.ReadAll(file1)
	content2, _ := ioutil.ReadAll(file1)

	return bytes.Equal(content1, content2)
}

func runCmd(cmdStr string, cmdArgs []string) error {
	cmd := exec.Command(cmdStr, cmdArgs...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	var tmpFile *os.File
	if cmdStr == "sort" {
		tmpFile, _ := os.Open("expected/data0.txt")
		cmd.Stdout = tmpFile
	}
	defer func() {
		if cmdStr == "sort" {
			tmpFile.Close()
		}
	}()
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("Error with case %s :%s ", cmdArgs, err)
	}

	return nil
}

func TestSort(t *testing.T) {
	for i := range goCmdArgs {
		err := runCmd("go", goCmdArgs[i])
		if err != nil {
			t.Errorf(err.Error())
		}
		err = runCmd("sort", osCmdArgs[i])
		if err != nil {
			t.Errorf(err.Error())
		}

		if !cmpRes("expected/"+file, "sorted/"+file) {
			t.Errorf("Unmatch with case %s", osCmdArgs[i])
		}
	}
}
