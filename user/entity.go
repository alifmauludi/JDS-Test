package user

import "time"

type User struct {
	UserID        int
	Username      string
	Password      string
	Role          string
	CreatedDate   time.Time
	UpdatedDate   time.Time
	PlainPassword string
}
