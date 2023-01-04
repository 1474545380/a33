package handler

import (
	"context"
	"staff/internal/repository"
	"staff/internal/service"
	"staff/pkg/e"
)

type StaffService struct {
}

func NewStaffService() *StaffService {
	return &StaffService{}
}

// 用户登录
func (*StaffService) StaffLogin(ctx context.Context, req *service.StaffRequest) (resp *service.StaffDetailResponse, err error) {
	var staff repository.Staff
	resp = new(service.StaffDetailResponse)
	resp.Code = e.Success
	err = staff.ShowStaffInfo(req)
	if err != nil {
		resp.Code = e.Error
		return resp, err
	}
	resp.StaffDetail = repository.BuildStaff(staff)
	return resp, nil
}
