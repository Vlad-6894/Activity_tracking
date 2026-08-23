package user

import "time"

type User struct {
	ID                 int64      `json:"id"`
	FullName           string     `json:"full_name"`
	Age                int        `json:"age"`
	GoogleRefreshToken *string    `json:"google_refresh_token,omitempty"`
	StepsGoal          *int       `json:"steps_goal,omitempty"`
	RestDays           int        `json:"rest_days"`
	Streak             int        `json:"streak"`
	CreatedAt          time.Time  `json:"created_at"`
	UpdatedAt          *time.Time `json:"updated_at,omitempty"`
}
