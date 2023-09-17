package main

type Event struct {
	UserID  string `json:"user_id"`
	EventID string `json:"event_id"`
	Date    string `json:"date"`
	Name    string `json:"event_name"`
}
