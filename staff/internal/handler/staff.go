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

// StaffLogin 员工登录
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

// StaffRegister 员工注册
func (*StaffService) StaffRegister(ctx context.Context, req *service.StaffRequest) (resp *service.StaffDetailResponse, err error) {
	var staff repository.Staff
	resp = new(service.StaffDetailResponse)
	resp.Code = e.Success
	staff, err = staff.StaffCreat(req)
	if err != nil {
		resp.Code = e.Error
		return resp, err
	}
	resp.StaffDetail = repository.BuildStaff(staff)
	return resp, err
}

// StaffDetails 员工详细信息
func (*StaffService) StaffDetails(ctx context.Context, req *service.StaffRequest) (resp *service.StaffDetailResponse, err error) {
	var staff repository.Staff
	resp = new(service.StaffDetailResponse)
	resp.Code = e.Success
	staff, err = staff.StaffDetailGet(req)
	if err != nil {
		resp.Code = e.Error
		return resp, err
	}
	resp.StaffDetail = repository.BuildStaff(staff)
	return resp, err
}

// StaffDetailsModify 修改员工详细信息
func (*StaffService) StaffDetailsModify(ctx context.Context, req *service.StaffRequest) (resp *service.StaffDetailResponse, err error) {
	var staff repository.Staff
	resp = new(service.StaffDetailResponse)
	resp.Code = e.Success
	err = staff.StaffDetailChange(req)
	if err != nil {
		resp.Code = e.Error
		return resp, err
	}
	resp.StaffDetail = repository.BuildStaff(staff)
	return resp, err
}

// StaffPreference 获取员工偏好信息
func (*StaffService) StaffPreference(ctx context.Context, req *service.StaffRequest) (resp *service.StaffPreferenceResponse, err error) {
	var staffPreference repository.StaffPreference
	resp = new(service.StaffPreferenceResponse)
	resp.Code = e.Success
	staffPreferenceDetail, err := staffPreference.StaffPreferenceGet(req)
	if err != nil {
		resp.Code = e.Error
		return resp, err
	}
	resp.StaffPreferenceDetail = repository.BuildStaffPreference(staffPreferenceDetail)
	return resp, err
}

// StaffPreferenceModify 修改员工偏好信息
func (*StaffService) StaffPreferenceModify(ctx context.Context, req *service.StaffPreferenceRequest) (resp *service.StaffPreferenceResponse, err error) {
	var staffPreference repository.StaffPreference
	resp = new(service.StaffPreferenceResponse)
	resp.Code = e.Success
	err = staffPreference.StaffPreferenceChange(req)
	if err != nil {
		resp.Code = e.Error
		return resp, err
	}
	resp.StaffDetail = repository.BuildStaff(staff)
	return resp, err
}
