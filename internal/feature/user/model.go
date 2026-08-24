package user

import "time"

type User struct {
	ID                 int64
	FullName           string
	Age                int
	GoogleRefreshToken *string
	StepsGoal          *int
	RestDays           int
	Streak             int
	CreatedAt          time.Time
	UpdatedAt          *time.Time
}

func NewUser(fullName string, age int, stepsGoal *int, restDays int) *User {
	return &User{
		FullName:  fullName,
		Age:       age,
		StepsGoal: stepsGoal,
		RestDays:  restDays,
		Streak:    0,
	}
}
