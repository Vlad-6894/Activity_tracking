package user

import "time"

type User struct {
	ID                 int
	FullName           string
	TelegramUserName   string // [ИЗМЕНЕНО]: Добавлено поле TelegramUserName (tg_user_name)
	Age                int
	GoogleRefreshToken string
	StepsGoal          int
	RestDays           int
	Streak             int
	CreatedAt          time.Time
	UpdatedAt          time.Time
}

// [ИЗМЕНЕНО]: Добавлен аргумент tgUserName в конструктор регистрации
func RegUser(fullName string, tgUserName string, age int, googleRefreshToken string, stepsGoal int, restDays int) User {
	return User{
		FullName:           fullName,
		TelegramUserName:   tgUserName,
		Age:                age,
		GoogleRefreshToken: googleRefreshToken,
		StepsGoal:          stepsGoal,
		RestDays:           restDays,
		Streak:             0,
	}
}

// [ИЗМЕНЕНО]: Добавлен аргумент tgUserName в конструктор полной модели
func NewUser(
	id int,
	fullName string,
	tgUserName string,
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
		TelegramUserName:   tgUserName,
		Age:                age,
		GoogleRefreshToken: googleRefreshToken,
		StepsGoal:          stepsGoal,
		RestDays:           restDays,
		Streak:             streak,
		CreatedAt:          createdAt,
		UpdatedAt:          updatedAt,
	}
}
