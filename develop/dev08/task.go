package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	gops "github.com/shirou/gopsutil/process"
)

/*
=== Взаимодействие с ОС ===

Необходимо реализовать собственный шелл

встроенные команды: cd/pwd/echo/kill/ps
поддержать fork/exec команды
конвеер на пайпах

Реализовать утилиту netcat (nc) клиент
принимать данные из stdin и отправлять в соединение (tcp/udp)
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func pwd() {
	ans, err := os.Getwd()
	if err != nil {
		fmt.Fprintf(os.Stderr, "pwd error: %s\n", err)
	} else {
		fmt.Println(ans)
	}
}
func echo(args []string) {
	fmt.Println(strings.Join(args, " "))
}
func cd(args []string) {
	if len(args) == 0 {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error cd: %s\n", err)
		}
		err = os.Chdir(homeDir)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error cd: %s\n", err)
		}
	} else if len(args) == 1 {
		err := os.Chdir(args[0])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error cd: %s\n", err)
		}
	} else {
		fmt.Fprintf(os.Stderr, "Incorrect amount of args for cd. Need 0 or 1\n")
	}
}
func ps() {
	processes, err := gops.Processes()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error ps: \n")
	} else {
		for _, process := range processes {
			pid := process.Pid
			name, _ := process.Name()
			status, _ := process.Status()
			fmt.Printf("%-10d %-20s %-5s\n", pid, name, status)
		}
	}
}
func kill(args []string) {
	if len(args) != 1 {
		fmt.Fprintf(os.Stderr, "Incorrect amount of args for kill. Need 1\n")
		return
	}
	pid, err := strconv.ParseInt(args[0], 10, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error kill: %s\n", err)
		return
	}
	proc, err := os.FindProcess(int(pid))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error kill: %s\n", err)
		return
	}
	err = proc.Kill()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error kill: %s\n", err)
		return
	}

}

// StartProcess is a low-level interface. The os/exec package provides higher-level interfaces.
func forkExec(args []string) {
	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "Incorrect amount of args for fork-exec. Need > 0\n")
		return
	}
	forkedArgs := []string{}
	if len(args) > 1 {
		forkedArgs = args[1:]
	}
	cmd := exec.Command(args[0], forkedArgs...)
	go func() {
		cmd.Run()
	}()

}
func myUnixShell() {

}
func main() {
	myUnixShell()
}
