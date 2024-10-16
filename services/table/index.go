package table

import (
	"gorm.io/gorm"
)

type TableService struct {
	DB *gorm.DB
}
