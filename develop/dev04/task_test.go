package main

import (
	"reflect"
	"testing"
)

var strs = []string{
	"пятак", "тяпка", "пятка",
	"листок", "столик", "слиток",
	"abc",
}
var expectedMap = make(map[string][]string, 2)

func TestMain(t *testing.T) {
	expectedMap["пятак"] = []string{"пятка", "тяпка"}
	expectedMap["листок"] = []string{"слиток", "столик"}
	gotMap := anagramSearch(strs)
	if !reflect.DeepEqual(expectedMap, gotMap) {
		t.Errorf("Incorrect test. Maps are not equal.\nExpected : %v\nRecieved: %v", expectedMap, gotMap)
	}
}
