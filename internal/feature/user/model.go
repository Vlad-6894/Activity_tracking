package user

import "time"

type User struct {
	ID                 int
	TelegramID         int64
	FullName           string
	TelegramUserName   string
	Age                int
	GoogleRefreshToken string
	StepsGoal          int
	RestDays           int
	Streak             int
	CreatedAt          time.Time
	UpdatedAt          time.Time
}

func RegUser(
	telegramID int64,
	fullName string,
	tgUserName string,
	age int,
	googleRefreshToken string,
	stepsGoal int,
	restDays int,
) User {
	return User{
		TelegramID:         telegramID,
		FullName:           fullName,
		TelegramUserName:   tgUserName,
		Age:                age,
		GoogleRefreshToken: googleRefreshToken,
		StepsGoal:          stepsGoal,
		RestDays:           restDays,
		Streak:             0,
	}
}

func NewUser(
	id int,
	telegramID int64,
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
		TelegramID:         telegramID,
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
