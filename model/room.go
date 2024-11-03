package model

import "gorm.io/gorm"

type Room struct {
	gorm.Model
	Name               string `json:"name" gorm:"not null"`
	Description        string `json:"description"`
	RoomNumber         string `json:"room_number" gorm:"not null"`
	Capacity           int    `json:"capacity" gorm:"not null"`
	MinReserveCapacity int    `json:"min_reserve_capacity" gorm:"not null"`
	IsAvailable        bool   `json:"is_available" gorm:"not null"`
}

func (s *Room) GetOneRoom(filter interface{}) error {
	result := MainDB.Model(&Room{}).Where(filter).First(s)
	return result.Error
}
