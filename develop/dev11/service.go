package main

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
)

var initID = "0"
var byDay = "day"
var byWeek = "week"
var byMonth = "month"

type EventService interface {
	CreateEvent(event Event) (err error)
	UpdateEvent(event Event) (err error)
	DeleteEvent(event Event) (err error)
	EventsForDay(day string) (events []Event, err error)
	EventsForWeek(week string) (events []Event, err error)
	EventsForMonth(month string) (events []Event, err error)
}

func NewFileEventService() *FileEventService {
	return &FileEventService{
		storage: NewJsonEventStorage("events"),
	}
}

type FileEventService struct {
	storage *JsonEventStorage
}

func (service *FileEventService) CreateEvent(event Event) (err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("Cant create event: %s", err.Error())
		}
	}()
	userDir := filepath.Join(service.storage.storageDir, event.UserID)
	_, err = os.Stat(userDir)
	if err == nil {
		var dirEnts []fs.FileInfo
		dirEnts, err = ioutil.ReadDir(userDir)
		if err != nil {
			return err
		}
		var lastID int
		lastID, err = strconv.Atoi(dirEnts[len(dirEnts)-1].Name())
		if err != nil {
			return err
		}
		event.EventID = strconv.Itoa(lastID + 1)
		err = service.storage.Save(event)
		if err != nil {
			return err
		}
	} else if os.IsNotExist(err) {
		event.EventID = initID
		err = service.storage.Save(event)
		if err != nil {
			return err
		}
	} else {
		return err
	}
	return nil
}
func (service *FileEventService) UpdateEvent(event Event) (err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("Cant update event: %s", err.Error())
		}
	}()
	eventDir := filepath.Join(service.storage.storageDir, event.UserID, event.EventID)
	_, err = os.Stat(eventDir)
	if err == nil {
		err = service.storage.Save(event)
		if err != nil {
			return err
		}
	} else {
		return err
	}
	return nil
}
func (service *FileEventService) DeleteEvent(event Event) (err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("Cant delete event: %s", err.Error())
		}
	}()
	eventDir := filepath.Join(service.storage.storageDir, event.UserID, event.EventID)
	_, err = os.Stat(eventDir)
	if err == nil {
		err = service.storage.Delete(event)
		if err != nil {
			return err
		}
	} else { /*может добавить else if, что если нет такого id, то это не ошибка ?*/
		return err
	}
	return nil
}
func (service *FileEventService) EventsForDay(day string) (_ []Event, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("Cant get events for day: %s", err.Error())
		}
	}()
	intDay, err := strconv.Atoi(day)
	if err != nil {
		return nil, err
	}
	return service.storage.GetBy(byDay, intDay)
}
func (service *FileEventService) EventsForWeek(week string) (_ []Event, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("Cant get events for week: %s", err.Error())
		}
	}()
	intWeek, err := strconv.Atoi(week)
	if err != nil {
		return nil, err
	}
	return service.storage.GetBy(byDay, intWeek)
}
func (service *FileEventService) EventsForMonth(month string) (_ []Event, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("Cant get events for month: %s", err.Error())
		}
	}()
	intMonth, err := strconv.Atoi(month)
	if err != nil {
		return nil, err
	}
	return service.storage.GetBy(byDay, intMonth)
}
