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

// RegUser конструктор для создания нового пользователя при регистрации (ID = 0, streak = 0)
func RegUser(fullName string, age int, stepsGoal *int, restDays int) *User {
	return &User{
		ID:        0,
		FullName:  fullName,
		Age:       age,
		StepsGoal: stepsGoal,
		RestDays:  restDays,
		Streak:    0,
	}
}

// NewUser конструктор для полной инициализации модели со всеми полями
func NewUser(
	id int64,
	fullName string,
	age int,
	googleRefreshToken *string,
	stepsGoal *int,
	restDays int,
	streak int,
	createdAt time.Time,
	updatedAt *time.Time,
) *User {
	return &User{
		ID:                 id,
		FullName:           fullName,
		Age:                age,
		GoogleRefreshToken: googleRefreshToken,
		StepsGoal:          stepsGoal,
		RestDays:           restDays,
		Streak:             streak,
		CreatedAt:          createdAt,
		UpdatedAt:          updatedAt,
	}
}
