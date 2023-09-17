package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

type EventStorage interface {
	Save(e Event) error
	Delete(e Event) error
	GetByDay(day int) ([]Event, error)
	GetByWeek(week int) ([]Event, error)
	GetByMonth(month int) ([]Event, error)
}

type JsonEventStorage struct {
	storageDir string
}

func NewJsonEventStorage(dir string) *JsonEventStorage {
	return &JsonEventStorage{
		storageDir: dir,
	}
}
func (storage JsonEventStorage) fileName(e Event) string {
	storePath := filepath.Join(storage.storageDir, e.UserID, e.EventID)
	return storePath
}
func (storage *JsonEventStorage) Save(e Event) error {
	storePath := storage.fileName(e)
	err := os.MkdirAll(filepath.Dir(storePath), os.ModePerm)
	if err != nil {
		return err
	}
	file, err := os.Create(storePath)
	if err != nil {
		return err
	}
	defer file.Close()

	saveToJSON := json.NewEncoder(file)
	if err := saveToJSON.Encode(e); err != nil {
		return err
	}
	return nil
}
func (storage *JsonEventStorage) Delete(e Event) error {
	storePath := storage.fileName(e)
	err := os.Remove(storePath)
	if err != nil {
		return err
	}
	return nil
}
func (storage *JsonEventStorage) GetByUserID(id string) ([]Event, error) {
	userDir := filepath.Join(storage.storageDir, id)
	files, err := ioutil.ReadDir(userDir)
	if err != nil {
		return nil, err
	}
	events := []Event{}
	for _, file := range files {

		fOpened, err := os.Open(filepath.Join(userDir, file.Name()))
		if err != nil {
			return nil, err
		}
		var tmpEvent Event
		err = json.NewDecoder(fOpened).Decode(&tmpEvent)
		if err != nil {
			return nil, err
		}
		events = append(events, tmpEvent)
	}
	return events, nil

}
func (storage *JsonEventStorage) GetByDay(day int) ([]Event, error) {
	events, err := storage.GetBy("day", day)
	if err != nil {
		return nil, err
	}
	return events, nil
}
func (storage *JsonEventStorage) GetByWeek(week int) ([]Event, error) {
	events, err := storage.GetBy("week", week)
	if err != nil {
		return nil, err
	}
	return events, nil
}
func (storage *JsonEventStorage) GetByMonth(month int) ([]Event, error) {
	events, err := storage.GetBy("month", month)
	if err != nil {
		return nil, err
	}
	return events, nil
}

func (storage *JsonEventStorage) GetBy(filter string, value int) ([]Event, error) {
	files, err := ioutil.ReadDir(storage.storageDir)
	if err != nil {
		return nil, err
	}
	ans := []Event{}
	for _, file := range files {
		events, err := storage.GetByUserID(file.Name())
		if err != nil {
			return nil, err
		}
		for _, event := range events {
			var getFunc func(string) (int, error)
			switch filter {
			case "day":
				getFunc = getDayFromDate
			case "week":
				getFunc = getWeekFromDate
			case "month":
				getFunc = getMonthFromDate
			default:
				return nil, fmt.Errorf("incorrect filter for getting events")

			}
			curEventValue, err := getFunc(event.Date)
			if err != nil {
				return nil, err
			}
			if value == curEventValue {
				ans = append(ans, event)
			}

		}
	}
	return ans, nil
}

/*
используем формат "2006-01-02", потому что в Golang есть специальная константа time.RFC3339,
которая представляет этот формат. Формат состоит из чисел,
разделенных дефисами или точками, и использует специальные
символы для обозначения годов, месяцев и дней.
*/
func getDayFromDate(date string) (int, error) {
	layout := "2006-01-02"
	parsedDate, err := time.Parse(layout, date)
	if err != nil {
		return -1, err
	}
	return parsedDate.Day(), nil

}
func getWeekFromDate(date string) (int, error) {
	layout := "2006-01-02"
	parsedDate, err := time.Parse(layout, date)
	if err != nil {
		return -1, err
	}
	_, week := parsedDate.ISOWeek()
	return week, nil

}
func getMonthFromDate(date string) (int, error) {
	layout := "2006-01-02"
	parsedDate, err := time.Parse(layout, date)
	if err != nil {
		return -1, err
	}
	month := parsedDate.Month()
	return int(month), nil

}
