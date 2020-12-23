package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// Base used for all struct
type Base struct {
	ID        uuid.UUID  `json:"id" gorm:"type:varchar(36);primarykey"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-"`
}
