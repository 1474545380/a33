package handler

import (
	pb "schedu_date/internal/service"
)

type StaffService struct {
	pb.UnimplementedSchedulDateServiceServer
}

func NewStaffService() *StaffService {
	return &StaffService{}
}
