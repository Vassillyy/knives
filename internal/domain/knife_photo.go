package domain

import "time"

type KnifePhoto struct {
	ID        string
	KnifeID   string
	S3Key     string
	Filename  string
	CreatedAt time.Time
	DeletedAt *time.Time
}