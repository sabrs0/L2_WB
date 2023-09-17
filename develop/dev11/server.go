package main

import (
	"fmt"
	"net/http"
)

func logMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//log.Printf("\nLOG: Method: %s, Host: %s, URL: %s\n", r.Method, r.Host, r.URL)
		next.ServeHTTP(w, r)
	})
}

type Server struct {
	//handler CalendarHandler
	server *http.Server
}

func newRouter(h CalendarHandler) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/create_event", h.CreateEvent)
	mux.HandleFunc("/update_event", h.UpdateEvent)
	mux.HandleFunc("/delete_event", h.DeleteEvent)
	mux.HandleFunc("/events_for_day", h.EventsForDay)
	mux.HandleFunc("/events_for_week", h.EventsForWeek)
	mux.HandleFunc("/events_for_month", h.EventsForMonth)
	return mux
}
func NewServer(addr string) *Server {
	handler := newCalendarHandler(NewFileEventService())
	routerWithLogs := logMiddleware(newRouter(*handler))
	tmpserver := http.Server{
		Addr:    addr,
		Handler: routerWithLogs,
	}
	return &Server{
		server: &tmpserver,
	}
}
func (s *Server) Start() {
	fmt.Println("Start server at ", s.server.Addr)
	s.server.ListenAndServe()
}
