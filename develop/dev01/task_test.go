package main

import (
	"fmt"
	"os"
	"testing"
)

type Tests struct {
	host      string
	stdErr    string
	isCorrect bool
}

func TestNTP(t *testing.T) {
	tests := []Tests{
		Tests{
			host:   "helloWorld",
			stdErr: "Cant get time from ntp",
		},
		Tests{
			host:   "0.beevik-ntp.pool.ntp.org",
			stdErr: "",
		},
	}
	for i, test := range tests {
		fmt.Println("TEST NUBMER ", i+1)
		host = test.host
		ExitFunc = func(int) {}
		main()
		var stdErrMsg string
		stderrLen, _ := os.Stderr.Seek(0, os.SEEK_CUR)
		if stderrLen != 0 {
			data, _ := os.ReadFile(os.Stderr.Name())

			stdErrMsg = string(data)
		}

		if len(stdErrMsg) > 0 && stdErrMsg[:len(test.stdErr)] != test.stdErr {
			t.Errorf("Error stdErr unmatched - real stdErr: %s, expected stdErr: %s\n", stdErrMsg, test.stdErr)
		}
	}
}
