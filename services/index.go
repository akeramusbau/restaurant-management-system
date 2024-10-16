package services

import (
	"github.com/akeramusbau/restaurant-management-system/services/table"
	"gorm.io/gorm"
)

var (
	TableService table.TableService
)

func InjectDBIntoServices(db *gorm.DB) {
	TableService.DB = db
}
