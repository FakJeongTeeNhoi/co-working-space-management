package gRPC

import (
	"context"

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
