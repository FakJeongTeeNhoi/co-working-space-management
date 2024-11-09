package model

import (
	"fmt"
	"time"

	"github.com/FakJeongTeeNhoi/co-working-space-management/service"
	"github.com/lib/pq"
	gorm "gorm.io/gorm"
)

type Spaces []Space
type Space struct {
	gorm.Model
	Name              string         `json:"name" gorm:"not null"`
	Description       string         `json:"description"`
	WorkingHour       pq.StringArray `json:"working_hour" gorm:"type:text[]"`
	Latitude          float64        `json:"latitude" gorm:"not null"`
	Longitude         float64        `json:"longitude" gorm:"not null"`
	Faculty           string         `json:"faculty" gorm:"not null"`
	Floor             int            `json:"floor" gorm:"not null"`
	Building          string         `json:"building" gorm:"not null"`
	Type              string         `json:"type" gorm:"not null"`
	HeadStaff         string         `json:"head_staff"`
	FacultyAccessList pq.StringArray `json:"faculty_access_list" gorm:"type:text[]"`
	StaffList         pq.Int64Array  `json:"staff_list" gorm:"type:integer[]"`
	RoomList          pq.StringArray `json:"room_list" gorm:"type:text[]"`
	IsAvailable       bool           `json:"is_available" gorm:"not null"`
}

type CreateSpaceRequest struct {
	Name              string         `json:"name" binding:"required"`
	Description       string         `json:"description"`
	WorkingHour       pq.StringArray `json:"working_hour" binding:"required"`
	Latitude          float64        `json:"latitude" binding:"required"`
	Longitude         float64        `json:"longitude" binding:"required"`
	Faculty           string         `json:"faculty" binding:"required"`
	Floor             int            `json:"floor" binding:"required"`
	Building          string         `json:"building" binding:"required"`
	Type              string         `json:"type" binding:"required"`
	HeadStaff         string         `json:"head_staff"`
	FacultyAccessList pq.StringArray `json:"faculty_access_list"`
	StaffList         pq.Int64Array  `json:"staff_list"`
	RoomList          pq.StringArray `json:"room_list"`
}

type SpaceResponses []SpaceResponse
type SpaceResponse struct {
	gorm.Model
	Name              string         `json:"name"`
	Description       string         `json:"description"`
	Opening_day       []string       `json:"opening_day"`
	WorkingHour       pq.StringArray `json:"working_hour"`
	Latitude          float64        `json:"latitude"`
	Longitude         float64        `json:"longitude"`
	Faculty           string         `json:"faculty"`
	Floor             int            `json:"floor"`
	Building          string         `json:"building"`
	Type              string         `json:"type"`
	HeadStaff         string         `json:"head_staff"`
	FacultyAccessList pq.StringArray `json:"faculty_access_list"`
	StaffList         pq.Int64Array  `json:"staff_list"`
	RoomList          []Room         `json:"room_list"`
	IsAvailable       bool           `json:"is_available"`
}

type SpaceSearchParam struct {
	Name            string     `json:"name"`
	Faculty         string     `json:"faculty"`
	Start_datetime  time.Time  `json:"start_datetime"`
	End_datetime    time.Time  `json:"end_datetime"`
	Capacity        int        `json:"capacity"`
	Latitude_range  [2]float64 `json:"latitude_range"`
	Longitude_range [2]float64 `json:"longitude_range"`
}

func (s *Space) GetOne(filter interface{}) error {
	result := MainDB.Model(&Space{}).Where(filter).First(s)
	return result.Error
}

func ToOpenDay(workingHour pq.StringArray) []string {
	opening_day := []string{}
	for i, working_hour := range workingHour {
		if working_hour != "Closed" {
			opening_day = append(opening_day, time.Weekday(i).String())
		}
	}
	return opening_day
}

func EmbeddedRoomList(space *Space, spaceResponse *SpaceResponse, capacity int) error {
	for _, room_name := range space.RoomList {
		room := Room{}
		room.GetOneRoom(map[string]interface{}{"name": room_name})
		if room.Capacity >= capacity {
			spaceResponse.RoomList = append(spaceResponse.RoomList, room)
		}
	}
	return nil
}

func EmbeddedRoomListWithSpace(space *Space, spaceResponse *SpaceResponse, ID uint) error {
	room := Room{}
	room.GetOneRoom(map[string]interface{}{"id": ID})
	spaceResponse.RoomList = append(spaceResponse.RoomList, room)
	return nil
}

func ToSpaceResponse(req *Space, res *SpaceResponse) {
	res.ID = req.ID
	res.Name = req.Name
	res.Description = req.Description
	res.Opening_day = ToOpenDay(req.WorkingHour)
	res.WorkingHour = req.WorkingHour
	res.Latitude = req.Latitude
	res.Longitude = req.Longitude
	res.Faculty = req.Faculty
	res.Floor = req.Floor
	res.Building = req.Building
	res.Type = req.Type
	res.HeadStaff = req.HeadStaff
	res.FacultyAccessList = req.FacultyAccessList
	res.StaffList = req.StaffList
	res.RoomList = []Room{}
	res.IsAvailable = req.IsAvailable
}

func (s *SpaceResponses) GetAllWithSearchParam(params SpaceSearchParam) error {

	query := MainDB.Model(&Space{})

	// Check each parameter in SpaceSearchParam and add to the query if provided
	if params.Name != "" {
		query = query.Where("name = ?", params.Name)
	}
	if params.Faculty != "" {
		query = query.Where("faculty = ?", params.Faculty)
	}
	// Extract the day of the week from start_datetime
	dayOfWeek := params.Start_datetime.Weekday() // returns a value from 0 (Sunday) to 6 (Saturday)

	// Get the corresponding working hour string for that day
	workingHourColumn := fmt.Sprintf("working_hour[%d]", dayOfWeek) // GORM uses 1-based indexing for PostgreSQL arrays

	if !params.Start_datetime.IsZero() && !params.End_datetime.IsZero() {
		// Prepare the time range to compare with the working hours
		startTime := params.Start_datetime.Format("15:04") // HH:MM format
		endTime := params.End_datetime.Format("15:04")

		// Query to check if the requested time range falls within the working hours
		query = query.Where(fmt.Sprintf("%s IS NOT NULL AND (? BETWEEN substring(%s from 1 for 5) AND substring(%s from 7 for 5) OR ? BETWEEN substring(%s from 1 for 5) AND substring(%s from 7 for 5) OR (? <= substring(%s from 7 for 5) AND ? >= substring(%s from 1 for 5))",
			workingHourColumn, workingHourColumn, workingHourColumn,
			startTime, endTime, startTime, endTime))
	}

	if params.Latitude_range != [2]float64{} {
		query = query.Where("latitude BETWEEN ? AND ?", params.Latitude_range[0], params.Latitude_range[1])
	}
	if params.Longitude_range != [2]float64{} {
		query = query.Where("longitude BETWEEN ? AND ?", params.Longitude_range[0], params.Longitude_range[1])
	}

	queried_spaces := Spaces{}
	// Execute the query
	result := query.Find(&queried_spaces)

	// Get the rooms associated with each space
	for _, space := range queried_spaces {
		temporary_space := SpaceResponse{}
		// Get the rooms associated with the space
		ToSpaceResponse(&space, &temporary_space)
		if params.Capacity > 0 {
			EmbeddedRoomList(&space, &temporary_space, params.Capacity)
		} else {
			EmbeddedRoomList(&space, &temporary_space, 0)
		}
		if len(temporary_space.RoomList) > 0 {
			*s = append(*s, temporary_space)
		}
	}

	return result.Error
}

func (s *CreateSpaceRequest) CreateSpace() (*SpaceResponse, error) {
	space := Space{
		Name:              s.Name,
		Description:       s.Description,
		WorkingHour:       s.WorkingHour,
		Latitude:          s.Latitude,
		Longitude:         s.Longitude,
		Faculty:           s.Faculty,
		Floor:             s.Floor,
		Building:          s.Building,
		Type:              s.Type,
		HeadStaff:         s.HeadStaff,
		FacultyAccessList: s.FacultyAccessList,
		StaffList:         s.StaffList,
		RoomList:          s.RoomList,
		IsAvailable:       true,
	}
	result := MainDB.Model(&Space{}).Create(&space)
	spaceResponse := SpaceResponse{}
	ToSpaceResponse(&space, &spaceResponse)
	return &spaceResponse, result.Error
}

func (s *Space) UpdateSpace(filter interface{}) error {
	result := MainDB.Model(&Space{}).Where(filter).Updates(s)
	return result.Error
}

type AddStaffToSpaceRequest struct {
	StaffIDs []int `json:"staff_id" binding:"required"`
}

func (s *AddStaffToSpaceRequest) AddStaffToSpace(spaceID int) error {
	space := Space{}
	space.GetOne(map[string]interface{}{"id": spaceID})
	for _, staffID := range s.StaffIDs {
		space.StaffList = append(space.StaffList, int64(staffID))
	}
	result := MainDB.Model(&Space{}).Where("id = ?", spaceID).Updates(&space)
	return result.Error
}

type RemoveStaffFromSpaceRequest struct {
	StaffIDs []int `json:"staff_id" binding:"required"`
}

func (s *RemoveStaffFromSpaceRequest) RemoveStaffFromSpace(spaceID int) error {
	space := Space{}
	space.GetOne(map[string]interface{}{"id": spaceID})
	for _, staffID := range s.StaffIDs {
		for i, staff := range space.StaffList {
			if staff == int64(staffID) {
				space.StaffList = append(space.StaffList[:i], space.StaffList[i+1:]...)
			}
		}
	}
	result := MainDB.Model(&Space{}).Where("id = ?", spaceID).Updates(&space)
	return result.Error
}

type AddRoomToSpaceRequest struct {
	RoomList []int `json:"room_list" binding:"required"`
}

func (s *AddRoomToSpaceRequest) AddRoomToSpace(spaceID int) error {
	space := Space{}
	space.GetOne(map[string]interface{}{"id": spaceID})
	for _, room := range s.RoomList {
		space.RoomList = append(space.RoomList, service.ParsToString(room))
		fmt.Println(string(room))
		fmt.Println(room)
	}
	result := MainDB.Model(&Space{}).Where("id = ?", spaceID).Updates(&space)
	return result.Error
}

type RemoveRoomFromSpaceRequest struct {
	RoomList []int `json:"room_list" binding:"required"`
}

func (s *RemoveRoomFromSpaceRequest) RemoveRoomFromSpace(spaceID int) error {
	space := Space{}
	space.GetOne(map[string]interface{}{"id": spaceID})
	for _, room := range s.RoomList {
		for i, room_name := range space.RoomList {
			if room_name == string(room) {
				space.RoomList = append(space.RoomList[:i], space.RoomList[i+1:]...)
			}
		}
	}
	result := MainDB.Model(&Space{}).Where("id = ?", spaceID).Updates(&space)
	return result.Error
}

func (s *Space) GetOneWithRoomInfo(roomID string) error {
	result := MainDB.Model(&Space{}).Where("room_list @> ARRAY[?]::text[]", roomID).First(s)
	return result.Error
}