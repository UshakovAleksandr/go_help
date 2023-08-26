package main

import (
	"fmt"
)

const tagName = "validate"

type User struct {
	ID     int     `json:"id,omitempty"`
	Name   string  `json:"name,omitempty" validate:"required,minLen=1"`
	Email  string  `json:"email,omitempty" validate:"required"`
	Age    int     `json:"age,omitempty" validate:"required,minValue=10,maxValue=100"`
	Salary float32 `json:"salary,omitempty" validate:"required,minValue=100.25"`
	Loan   float64 `json:",omitempty" validate:"required,minValue=200.11"`
}

func NewUser(id, age int, name, email string,
	salary float32, loan float64) *User {
	return &User{
		ID:     id,
		Name:   name,
		Email:  email,
		Age:    age,
		Salary: salary,
		Loan:   loan,
	}
}

func main() {
	user := NewUser(1, 25, "Alex",
		"email@host.ru", 2220, 2005,
	)
	err := Validate(user)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("ok")
}

// разъименования указателя в поле для получения значения
//field := elem.Field(i)
//if field.Kind() == reflect.Ptr {
//	field = field.Elem()
//}
//fmt.Println(field)
