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

type RoomCreateRequest struct {
	Name               string `json:"name" binding:"required"`
	Description        string `json:"description"`
	RoomNumber         string `json:"room_number" binding:"required"`
	Capacity           int    `json:"capacity" binding:"required"`
	MinReserveCapacity int    `json:"min_reserve_capacity" binding:"required"`
}

type RoomResponse struct {
	ID                 uint   `json:"id"`
	Name               string `json:"name"`
	Description        string `json:"description"`
	RoomNumber         string `json:"room_number"`
	Capacity           int    `json:"capacity"`
	MinReserveCapacity int    `json:"min_reserve_capacity"`
	IsAvailable        bool   `json:"is_available"`
}

func ToRoomResponse(s *Room, sr *RoomResponse) {
	sr.ID = s.ID
	sr.Name = s.Name
	sr.Description = s.Description
	sr.RoomNumber = s.RoomNumber
	sr.Capacity = s.Capacity
	sr.MinReserveCapacity = s.MinReserveCapacity
	sr.IsAvailable = s.IsAvailable
}

func (s *Room) GetOneRoom(filter interface{}) error {
	result := MainDB.Model(&Room{}).Where(filter).First(s)
	return result.Error
}

func (s *Room) GetAllRooms(filter interface{}) error{
	result := MainDB.Model(&Room{}).Where(filter).Find(s)
	return result.Error
}

func (s *RoomCreateRequest) CreateRoom() (*Room, error) {
	room := Room{
		Name:               s.Name,
		Description:        s.Description,
		RoomNumber:         s.RoomNumber,
		Capacity:           s.Capacity,
		MinReserveCapacity: s.MinReserveCapacity,
		IsAvailable:        true,
	}
	result := MainDB.Model(&Room{}).Create(&room)
	return &room, result.Error
}

func (s *Room) UpdateRoom(filter interface{}) error {
	result := MainDB.Model(&Room{}).Where(filter).Updates(s)
	return result.Error
}

func (s *Room) DeleteRoom(filter interface{}) error {
	result := MainDB.Model(&Room{}).Where(filter).Delete(s)
	return result.Error
}
