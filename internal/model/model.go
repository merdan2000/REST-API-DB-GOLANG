package model

import "time"

type Users struct {
	Id        int       `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Age       int       `json:"age"`
	Password  string    `json:"password"`
	Created   time.Time `json:"created"`
}
