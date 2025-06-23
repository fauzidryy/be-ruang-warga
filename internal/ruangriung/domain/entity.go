package domain

import (
	"time"

	"github.com/google/uuid"
)

type RuangRiung struct {
	ID uuid.UUID `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	ScheduleTime time.Time `json:"schedule_time"`
	Location string `json:"location"`
	PosterPath string `json:"poster_path"`
	Performers string `json:"performers"`
	CreatedAt time.Time `json:"created_at"`
}