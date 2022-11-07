package repostitory

import (
	"time"

	"github.com/Bookings/internal/model"
)

type Repo interface {
	Cek() bool
	SearchAvailabilityRooms(start time.Time, end time.Time) ([]model.Room, error)
}
