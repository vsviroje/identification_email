package models

import (
	"time"
)

// Config ...
type Config struct {
	ID        int       `db:"id" json:"configId"`
	Type      string    `db:"type" json:"type"`
	Key       string    `db:"key" json:"key"`
	Value     string    `db:"value" json:"value"`
	CreatedAt time.Time `db:"createdAt" json:"createdAt"`
	UpdatedAt time.Time `db:"updatedAt" json:"updatedAt"`
}
