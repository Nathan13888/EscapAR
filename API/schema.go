package main

import (
	"time"
	"uuid"
)

type UserEntry struct {
	Time         time.Time
	ProfileImage string
	BestScore    int
	ID           uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
}
