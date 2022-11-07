package model

import "time"

type ReservationRoom struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	Phone     string
	StartDate time.Time
	EndDate   time.Time
	Processed int
	RoomId    int
	CratedAt  time.Time
	UpdateAt  time.Time
	Room      Room
}
type Room struct {
	ID       int
	RoomName string
	CreateAt time.Time
	UpdateAt time.Time
}
