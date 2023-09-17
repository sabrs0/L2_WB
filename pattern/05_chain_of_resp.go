package pattern

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

type SystemDB map[string]User

var db SystemDB = make(SystemDB)

type User struct {
	login    string
	password string
}

type Validator interface {
	isCorrect(u User) bool
	getNext() Validator
}

type BaseValidator struct {
	next Validator
}

func (validator BaseValidator) getNext() Validator {
	return validator.next
}
func (validator *BaseValidator) setNext(next Validator) {
	validator.next = next
}
func (validator BaseValidator) isCorrect(u User) bool {
	return true
}

type LoginValidator struct {
	BaseValidator
}

func (validator LoginValidator) isCorrect(u User) bool {
	toCheck := strings.ToLower(u.login)
	if toCheck != u.login {
		fmt.Println("All letters should be in low case")
		return false
	}
	if _, isInUse := db[u.login]; isInUse {
		fmt.Println("Sorry, this login is in use, try again")
	}
	return true

}

type PasswordValidator struct {
	BaseValidator
}

func (validator PasswordValidator) isCorrect(u User) bool {
	passwd := u.password
	if strings.IndexRune(passwd, '@') < 0 &&
		strings.IndexRune(passwd, '!') < 0 &&
		strings.IndexRune(passwd, '&') < 0 {
		fmt.Println("Password should have @ ! or & character")
		return false
	}
	if utf8.RuneCount([]byte(passwd)) < 8 {
		fmt.Println("Password should have at least 8 characters")
		return false
	}
	return true

}

func Validate(validator Validator, u User) {
	if validator.isCorrect(u) {
		if validator.getNext() == nil {
			fmt.Println(u.login, ": Successfully registered")
		} else {
			Validate(validator.getNext(), u)
		}

	} else {
		fmt.Println(u.login, ": Registration Error")
	}
}

func chainOfRespPattern() {
	loginVal := LoginValidator{}
	paswwdVal := PasswordValidator{}
	loginVal.setNext(&paswwdVal)

	worstUser := User{
		login:    "BADUSER",
		password: "12345",
	}
	Validate(loginVal, worstUser)

	notGoodUser := User{
		login:    "notgood",
		password: "12345",
	}
	Validate(loginVal, notGoodUser)
	goodUser := User{
		login:    "good",
		password: "12345678!",
	}
	Validate(loginVal, goodUser)

}
