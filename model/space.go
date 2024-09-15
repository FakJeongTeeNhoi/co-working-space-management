package model

import (
	"github.com/lib/pq"
	gorm "gorm.io/gorm"
)

type Space struct {
	gorm.Model
	Name 				string `json:"name" gorm:"not null"`
	Description 		string `json:"description"`
	WorkingHour 		pq.StringArray `json:"working_hour" gorm:"type:text[]"`
	Latitude 			float64 `json:"latitude" gorm:"not null"`
	Longitude 			float64 `json:"longitude" gorm:"not null"`
	Faculty 	 		string `json:"faculty" gorm:"not null"`
	Floor 		 		int `json:"floor" gorm:"not null"`
	Building 	 	  	string `json:"building" gorm:"not null"`
	Type 		 	  	string `json:"type" gorm:"not null"`
	HeadStaff 	 	  	string `json:"head_staff"`
	FacultyAccessList 	pq.StringArray `json:"faculty_access_list" gorm:"type:text[]"`
	StaffList 	 		pq.Int64Array `json:"staff_list" gorm:"type:integer[]"`
	RoomList 	 		pq.StringArray `json:"room_list" gorm:"type:text[]"`
	IsAvailable 	 	bool `json:"is_available" gorm:"not null"`
}
