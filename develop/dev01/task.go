package main

/*
=== Базовая задача ===

Создать программу печатающую точное время с использованием NTP библиотеки.Инициализировать как go module.
Использовать библиотеку https://github.com/beevik/ntp.
Написать программу печатающую текущее время / точное время с использованием этой библиотеки.

Программа должна быть оформлена с использованием как go module.
Программа должна корректно обрабатывать ошибки библиотеки: распечатывать их в STDERR и возвращать ненулевой код выхода в OS.
Программа должна проходить проверки go vet и golint.
*/

import (
	"fmt"
	"os"
	"time"

	"github.com/beevik/ntp"
)

var host string = "0.beevik-ntp.pool.ntp.org"
var ExitFunc func(int) = os.Exit

func getNTPTime(host string) (time.Time, error) {
	return ntp.Time(host)
}

func main() {
	time, err := getNTPTime(host)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cant get time from ntp: %s\n", err.Error())
		ExitFunc(1)
	}
	fmt.Println("Current time: ", time.String())
}
