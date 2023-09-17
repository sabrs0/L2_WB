package main

import (
	"fmt"
	"strconv"
	"time"
)

type Validator struct {
}

func (v Validator) ValidateID(id string) error {
	intID, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	if intID < 0 {
		return fmt.Errorf("Negative userID")
	}
	return nil
}

/*
используем формат "2006-01-02", потому что в Golang есть специальная константа time.RFC3339,
которая представляет этот формат. Формат состоит из чисел,
разделенных дефисами или точками, и использует специальные
символы для обозначения годов, месяцев и дней.
*/
func (v Validator) ValidateDate(date string) error {
	if date == "" {
		return fmt.Errorf("Empty date recieved")
	}
	layout := "2006-01-02"
	_, err := time.Parse(layout, date)
	if err != nil {
		return err
	}
	return nil
}
func (v Validator) ValidateEventName(name string) error {
	if name == "" {
		return fmt.Errorf("Empty name recieved")
	}
	return nil
}
func (v Validator) ValidateCreate(userID, date, name string) error {
	err := v.ValidateID(userID)
	if err != nil {
		return err
	}
	err = v.ValidateDate(date)
	if err != nil {
		return err
	}
	err = v.ValidateEventName(name)
	if err != nil {
		return err
	}
	return nil
}
func (v Validator) ValidateUpdate(userID, eventID, date, name string) error {
	err := v.ValidateID(userID)
	if err != nil {
		return err
	}
	err = v.ValidateID(eventID)
	if err != nil {
		return err
	}
	err = v.ValidateDate(date)
	if err != nil {
		return err
	}
	err = v.ValidateEventName(name)
	if err != nil {
		return err
	}
	return nil
}
