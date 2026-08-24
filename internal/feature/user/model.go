package user

import "time"

// [ИЗМЕНЕНО]: Все поля теперь без указателей (значимые типы)
type User struct {
	ID                 int
	FullName           string
	Age                int
	GoogleRefreshToken string
	StepsGoal          int
	RestDays           int
	Streak             int
	CreatedAt          time.Time
	UpdatedAt          time.Time
}

// [ИЗМЕНЕНО]: Конструктор регистрации теперь принимает googleRefreshToken и возвращает User по значению
func RegUser(fullName string, age int, googleRefreshToken string, stepsGoal int, restDays int) User {
	return User{
		FullName:           fullName,
		Age:                age,
		GoogleRefreshToken: googleRefreshToken,
		StepsGoal:          stepsGoal,
		RestDays:           restDays,
		Streak:             0,
	}
}

// [ИЗМЕНЕНО]: Конструктор полной модели возвращает User по значению без указателей
func NewUser(
	id int,
	fullName string,
	age int,
	googleRefreshToken string,
	stepsGoal int,
	restDays int,
	streak int,
	createdAt time.Time,
	updatedAt time.Time,
) User {
	return User{
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
