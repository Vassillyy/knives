package models

import "time"

type KnifePhoto struct {
	ID        string     `db:"id"         json:"id"`
	KnifeID   string     `db:"knife_id"   json:"knife_id"`
	S3Key     string     `db:"s3_key"      json:"s3_key"`
	Filename  string     `db:"filename"   json:"filename"`
	CreatedAt time.Time  `db:"created_at" json:"created_at"`
	DeletedAt *time.Time `db:"deleted_at" json:"-"`
}
