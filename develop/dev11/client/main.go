package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Event struct {
	UserID  string `json:"user_id"`
	EventID string `json:"event_id"`
	Date    string `json:"date"`
	Name    string `json:"event_name"`
}

func createMethod() {
	createVals := url.Values{}
	createVals.Add("user_id", "1")
	createVals.Add("date", "2023-01-01")
	createVals.Add("event_name", "new yearr")
	resp, err := http.PostForm("http://localhost:8080/create_event", createVals)
	if err != nil {
		fmt.Println(err)
	} else {

		defer resp.Body.Close()
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(data))
	}
}
func getByDayMethod() {
	valUrl := url.URL{}
	valUrl.Scheme = "http"
	valUrl.Host = "localhost:8080"
	valUrl.Query().Add("day", "1")
	resp, err := http.Get("http://localhost:8080/events_for_day?day=1")
	if err != nil {
		fmt.Println(err)
	} else {

		defer resp.Body.Close()
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(data))
	}
}
func main() {
	//createMethod()
	getByDayMethod()
}
