package model

import (
	"gorm.io/gorm"
)

type Table struct {
	gorm.Model
	PlayerUsername string
	RoomID         uint
}
