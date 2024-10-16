package match

import (
	"github.com/akeramusbau/restaurant-management-system/model"
)

func (s TableService) PlayerJoinRoom(username string, roomID uint) (model.Table, error) {
	table := model.Table{
		PlayerUsername: username,
		RoomID:         roomID,
	}
	if err := s.DB.Create(&table).Error; err != nil {
		return table, err
	}
	return table, nil
}
