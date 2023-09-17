package main

import "testing"

type TestRes struct {
	res string
	err string
}

var (
	testStrings = []string{
		"a4bc2d5e", "abcd", "45", "", `qwe\4\5`, `qwe\45`, `qwe\\5`, `a\`,
	}
	ExpectedResults = []TestRes{
		TestRes{
			res: "aaaabccddddde",
			err: "",
		},
		TestRes{
			res: "abcd",
			err: "",
		},
		TestRes{
			res: "",
			err: "Incorrect string",
		},
		TestRes{
			res: "",
			err: "",
		},
		TestRes{
			res: "qwe45",
			err: "",
		},
		TestRes{
			res: "qwe44444",
			err: "",
		},
		TestRes{
			res: `qwe\\\\\`,
			err: "",
		},
		TestRes{
			res: "",
			err: "Incorrect string",
		},
	}
)

func TestUnpack(t *testing.T) {
	for i, test := range testStrings {
		res, err := Unpack(test)
		var errMsg string
		if err != nil {
			errMsg = err.Error()
		}
		if res != ExpectedResults[i].res || errMsg != ExpectedResults[i].err {
			t.Errorf("Unexpected result: have res = %s and err = %s, want res = %s and err = %s",
				res, err.Error(), ExpectedResults[i].res, ExpectedResults[i].err)
		}
	}
}
