package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

/*
=== Утилита telnet ===

Реализовать примитивный telnet клиент:
Примеры вызовов:
go-telnet --timeout=10s host port go-telnet mysite.ru 8080 go-telnet --timeout=3s 1.1.1.1 123

Программа должна подключаться к указанному хосту (ip или доменное имя) и порту по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s).

При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться.
При подключении к несуществующему сервер, программа должна завершаться через timeout.
*/

type configSet struct {
	timeOut time.Duration
	ip      string
	port    string
}

var flagT = flag.String("t", "10s", "timeout")
var cfg configSet

func checkFlags() error {
	flag.Parse()
	args := flag.Args()
	if len(args) != 2 {
		return fmt.Errorf("Incorrect num of args. Needs ip and port")
	}
	cfg.ip = args[0]
	cfg.port = args[1]
	t, err := time.ParseDuration(*flagT)
	if err != nil {
		return fmt.Errorf("Cant parse timeout: %s", err)
	}
	cfg.timeOut = t
	fmt.Println(cfg.timeOut)
	return nil
}
func mustCopy(dst io.Writer, src io.Reader) {
	n, err := io.Copy(dst, src)
	if err != nil {
		panic(err)
	} else if n == 0 { //EOF
		return
	}
}

func myTelnet(eofChan chan struct{}) error {
	c, err := net.DialTimeout("tcp", cfg.ip+":"+cfg.port, cfg.timeOut)
	if err != nil {
		return err
	}
	defer c.Close()
	badServChan := make(chan struct{})
	go func() {
		defer func() {
			badServChan <- struct{}{}
		}()
		go func() {
			mustCopy(c, os.Stdout)
			eofChan <- struct{}{}
		}()
		mustCopy(os.Stdin, c)

	}()
LOOP:
	for {
		select {
		case <-badServChan:
			fmt.Println("Interrupted by bad server response. Exiting")
			break LOOP
		case <-eofChan:
			fmt.Println("Interrupted by ctrlD. Exiting")
			break LOOP
		default:
			continue
		}
	}
	return nil
}

func main() {
	if err := checkFlags(); err != nil {
		panic(err)
	}
	ch := make(chan struct{})
	err := myTelnet(ch)
	if err != nil {
		fmt.Println("ERR TELNET - ", err)
	}

}
