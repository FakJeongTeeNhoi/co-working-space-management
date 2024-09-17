package gRPC

import (
	"context"
	"errors"

	pb "github.com/FakJeongTeeNhoi/co-working-space-management/generated/space"
	"github.com/FakJeongTeeNhoi/co-working-space-management/model"
	"gorm.io/gorm"
)

func ToSpace(req *pb.EditSpaceRequest, space *model.Space) {
	space.Name = req.Name
	space.Description = req.Description
	space.WorkingHour = req.WorkingHours
	space.Latitude = float64(req.Latitude)
	space.Longitude = float64(req.Longitude)
	space.Faculty = req.Faculty
	space.Floor = int(req.Floor)
	space.Building = req.Building
	space.Type = req.Type
	space.HeadStaff = req.HeadStaff
	space.FacultyAccessList = req.FacultyAccessList
	space.RoomList = req.RoomList
	space.IsAvailable = req.IsAvailable
}

func validateCreateSpaceRequest(req *pb.CreateSpaceRequest) error {
	if req.Name == "" {
		return errors.New("name is required")
	}
	if req.Latitude < -90 || req.Latitude > 90 {
		return errors.New("latitude must be between -90 and 90")
	}
	if req.Longitude < -180 || req.Longitude > 180 {
		return errors.New("longitude must be between -180 and 180")
	}
	if req.Faculty == "" {
		return errors.New("faculty is required")
	}
	if req.Floor <= 0 {
		return errors.New("floor must be greater than 0")
	}
	if req.Building == "" {
		return errors.New("building is required")
	}
	if req.Type == "" {
		return errors.New("type is required")
	}
	if !req.IsAvailable {
		return errors.New("availability status is required")
	}
	return nil
}

func (s *SpaceServer) CreateSpace(ctx context.Context, req *pb.CreateSpaceRequest) (*pb.GetSpaceResponse, error) {
	// Proceed with validation and space creation
	if err := validateCreateSpaceRequest(req); err != nil {
		return &pb.GetSpaceResponse{Success: false}, err
	}

	// Create a new space object to save in the database
	var space model.Space

	space.Name = req.Name
	space.Description = req.Description
	space.WorkingHour = req.WorkingHours
	space.Latitude = float64(req.Latitude)
	space.Longitude = float64(req.Longitude)
	space.Faculty = req.Faculty
	space.Floor = int(req.Floor)
	space.Building = req.Building
	space.Type = req.Type
	space.HeadStaff = req.HeadStaff
	space.StaffList = req.StaffList
	space.FacultyAccessList = req.FacultyAccessList
	space.RoomList = req.RoomList
	space.IsAvailable = req.IsAvailable

	if err := s.db.Create(&space).Error; err != nil {
		return &pb.GetSpaceResponse{Success: false}, err
	}

	return &pb.GetSpaceResponse{
		Success: true,
		SpaceId: int64(space.ID),
		Name: space.Name,
		Description: space.Description,
		WorkingHours: space.WorkingHour,
		Latitude: float32(space.Latitude),
		Longitude: float32(space.Longitude),
		Faculty: space.Faculty,
		Floor: int64(space.Floor),
		Building: space.Building,
		Type: space.Type,
		HeadStaff: space.HeadStaff,
		FacultyAccessList: space.FacultyAccessList,
		RoomList: space.RoomList,
		IsAvailable: space.IsAvailable,
	}, nil
}


func (s *SpaceServer) GetAllSpace(ctx context.Context, req *pb.GetAllSpaceRequest) (*pb.GetAllSpaceResponse, error) {
    var spaces []model.Space

    // Retrieve all spaces from the database
    if err := s.db.Find(&spaces).Error; err != nil {
        return &pb.GetAllSpaceResponse{Success: false, Message: "Failed to get all Co-Working Spaces"}, err
    }

    // Convert the model spaces to protobuf spaces
    var pbSpaces []*pb.Space
    for _, space := range spaces {
        pbSpace := &pb.Space{
			SpaceId: 		   int64(space.ID),
            Name:              space.Name,
            Description:       space.Description,
            WorkingHours:      space.WorkingHour,
            Latitude:          float32(space.Latitude),
            Longitude:         float32(space.Longitude),
            Faculty:           space.Faculty,
            Floor:             int64(space.Floor),
            Building:          space.Building,
            Type:              space.Type,
            HeadStaff:         space.HeadStaff,
			StaffList:         space.StaffList,
            FacultyAccessList: space.FacultyAccessList,
            RoomList:          space.RoomList,
            IsAvailable:       space.IsAvailable,
        }
        pbSpaces = append(pbSpaces, pbSpace)
    }

    return &pb.GetAllSpaceResponse{Success: true, Message: "Spaces retrieved successfully", Spaces: pbSpaces}, nil
}

func (s *SpaceServer) EditSpaceDetail(ctx context.Context, req *pb.EditSpaceRequest) (*pb.SpaceServiceResponse, error) {
	var space model.Space
	if err := s.db.First(&space, req.SpaceId).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return &pb.SpaceServiceResponse{Success: false, Message: "Co-Working Space not found"}, nil
		}
		return &pb.SpaceServiceResponse{Success: false, Message: "Error retrieving Co-Working Space"}, err
	}

	ToSpace(req, &space)

	if err := s.db.Save(&space).Error; err != nil {
		return &pb.SpaceServiceResponse{Success: false, Message: "Failed to update Co-Working Space"}, err
	}

	return &pb.SpaceServiceResponse{Success: true, Message: "Co-Working Space updated successfully"}, nil
}

func (s *SpaceServer) DeleteSpace(ctx context.Context, req *pb.DeleteSpaceRequest) (*pb.SpaceServiceResponse, error) {
	if err := s.db.Delete(&model.Space{}, req.SpaceId).Error; err != nil {
		return &pb.SpaceServiceResponse{Success: false, Message: "Failed to delete Co-Working Space"}, err
	}

	return &pb.SpaceServiceResponse{Success: true, Message: "Co-Working Space deleted successfully"}, nil
}

func (s *SpaceServer) GetSpace(ctx context.Context, req *pb.GetSpaceRequest) (*pb.GetSpaceResponse, error) {
	var space model.Space
	if err := s.db.First(&space, req.SpaceId).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return &pb.GetSpaceResponse{Success: false}, nil
		}
		return nil, err
	}

	return &pb.GetSpaceResponse{
		Success: true,
		SpaceId: int64(space.ID),  // Include space_id in the response
		Name: space.Name,
		Description: space.Description,
		WorkingHours: space.WorkingHour,
		Latitude: float32(space.Latitude),
		Longitude: float32(space.Longitude),
		Faculty: space.Faculty,
		Floor: int64(space.Floor),
		Building: space.Building,
		Type: space.Type,
		HeadStaff: space.HeadStaff,
		FacultyAccessList: space.FacultyAccessList,
		RoomList: space.RoomList,
		IsAvailable: space.IsAvailable,
	}, nil
}
