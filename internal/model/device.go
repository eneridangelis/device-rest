package model

import (
	"time"
)

type Device struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string    `gorm:"type:varchar(100);not null" json:"name"`
	Brand     string    `gorm:"type:varchar(50);not null" json:"brand"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}
