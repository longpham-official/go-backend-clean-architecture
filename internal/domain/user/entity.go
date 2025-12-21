package user

import "time"

type User struct {
	ID        int64
	Email     string
	Password  string // hashed password
	CreatedAt time.Time
	UpdatedAt time.Time
}