package main

import (
	"encoding/json"
	"net/http"
)

type CalendarHandler struct {
	service EventService
}

func newCalendarHandler(service EventService) *CalendarHandler {
	return &CalendarHandler{
		service: service,
	}
}

// posts
func (handler CalendarHandler) CreateEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		wrapMsgResponse(w, 500, "Incorrect method. Need POST.", "error")
	} else {
		userID := r.FormValue("user_id")
		date := r.FormValue("date")
		name := r.FormValue("event_name")
		validator := Validator{}
		err := validator.ValidateCreate(userID, date, name)
		if err != nil {
			wrapMsgResponse(w, 400, err.Error(), "error")
		} else {
			event := Event{
				UserID: userID,
				Date:   date,
				Name:   name,
			}
			err = handler.service.CreateEvent(event)
			if err != nil {
				wrapMsgResponse(w, 503, err.Error(), "error")
			} else {
				wrapMsgResponse(w, 200, "Successfully created", "result")
			}
		}

	}

}
func (handler CalendarHandler) UpdateEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		wrapMsgResponse(w, 500, "Incorrect method. Need POST.", "error")
	} else {
		userID := r.FormValue("user_id")
		eventID := r.FormValue("event_id")
		date := r.FormValue("date")
		name := r.FormValue("event_name")
		validator := Validator{}
		err := validator.ValidateUpdate(userID, eventID, date, name)
		if err != nil {
			wrapMsgResponse(w, 400, err.Error(), "error")
		} else {
			event := Event{
				UserID:  userID,
				EventID: eventID,
				Date:    date,
				Name:    name,
			}
			err = handler.service.UpdateEvent(event)
			if err != nil {
				wrapMsgResponse(w, 503, err.Error(), "error")
			} else {
				wrapMsgResponse(w, 200, "Successfully updated", "result")
			}

		}

	}

}
func (handler CalendarHandler) DeleteEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		wrapMsgResponse(w, 500, "Incorrect method. Need POST.", "error")
	} else {
		userID := r.FormValue("user_id")
		eventID := r.FormValue("event_id")
		validator := Validator{}
		err := validator.ValidateID(userID)
		if err != nil {
			wrapMsgResponse(w, 400, err.Error(), "error")
		} else {
			err = validator.ValidateID(eventID)
			if err != nil {
				wrapMsgResponse(w, 400, err.Error(), "error")
			} else {
				event := Event{
					UserID:  userID,
					EventID: eventID,
				}
				err = handler.service.DeleteEvent(event)
				if err != nil {
					wrapMsgResponse(w, 503, err.Error(), "error")
				} else {
					wrapMsgResponse(w, 200, "Successfully deleted", "result")
				}

			}

		}

	}

}

// gets
func (handler CalendarHandler) EventsForDay(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		wrapMsgResponse(w, 500, "Incorrect method. Need GET.", "error")
	} else {
		day := r.FormValue("day")
		events, err := handler.service.EventsForDay(day)
		if err != nil {
			wrapMsgResponse(w, 503, err.Error(), "error")
		} else {
			jsonResponse := map[string]interface{}{
				"result": events,
			}
			err = json.NewEncoder(w).Encode(jsonResponse)
			if err != nil {
				wrapMsgResponse(w, 500, err.Error(), "error")
			} else {
				w.WriteHeader(200)
			}

		}

	}

}
func (handler CalendarHandler) EventsForWeek(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		wrapMsgResponse(w, 500, "Incorrect method. Need GET.", "error")
	} else {
		week := r.FormValue("week")

		events, err := handler.service.EventsForWeek(week)
		if err != nil {
			wrapMsgResponse(w, 503, "Incorrect method. Need GET.", "error")
		} else {
			jsonResponse := map[string]interface{}{
				"result": events,
			}
			err = json.NewEncoder(w).Encode(jsonResponse)
			if err != nil {
				wrapMsgResponse(w, 500, "Incorrect method. Need GET.", "error")
			} else {
				w.WriteHeader(200)
			}
		}

	}

}
func (handler CalendarHandler) EventsForMonth(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		wrapMsgResponse(w, 500, "Incorrect method. Need GET.", "error")
	} else {

		month := r.FormValue("month")

		events, err := handler.service.EventsForWeek(month)
		if err != nil {
			wrapMsgResponse(w, 503, "Incorrect method. Need GET.", "error")
		} else {
			jsonResponse := map[string]interface{}{
				"result": events,
			}
			err = json.NewEncoder(w).Encode(jsonResponse)
			if err != nil {
				wrapMsgResponse(w, 500, "Incorrect method. Need GET.", "error")
			} else {
				w.WriteHeader(200)
			}
		}

	}

}
func wrapMsgResponse(w http.ResponseWriter, code int, msg string, result string) {
	jsonResponse := map[string]interface{}{
		result: msg,
	}
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(jsonResponse)
}
