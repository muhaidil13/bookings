package mysql_repo

import (
	"context"
	"time"

	"github.com/Bookings/internal/model"
)

func (m *MysqlDbrepo) Cek() bool {
	return true
}

func (m *MysqlDbrepo) SearchAvailabilityRooms(start, end time.Time) ([]model.Room, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()
	var rooms []model.Room
	stmt := `
	select r.id, r.room_name from rooms as r where r.id
	not in 
	(select rr.room_id from room_restrictions as rr where ? < rr.end_date and ? > rr.start_date);
	`
	rows, err := m.DB.QueryContext(ctx, stmt, start, end)

	if err != nil {
		return rooms, err
	}
	for rows.Next() {
		var room model.Room
		rows.Scan(
			&room.ID,
			&room.RoomName,
		)
		rooms = append(rooms, room)
	}
	if rows.Err(); err != nil {
		return rooms, err
	}
	return rooms, nil
}
