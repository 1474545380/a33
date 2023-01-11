package handler

import (
	"context"
	"store/internal/repository"
	"store/internal/service"
	pb "store/internal/service"
	"store/pkg/e"
)

type StoreService struct {
	pb.UnimplementedStoreServiceServer
}

func NewStoreService() *StoreService {
	return &StoreService{}
}

// StoreFill 添加新店铺
func (*StoreService) StoreFill(ctx context.Context, req *service.StoreRequest) (resp *service.StoreDetailResponse, err error) {
	var store repository.Store
	resp = new(service.StoreDetailResponse)
	resp.Code = e.Success
	store, err = store.StoreCreat(req)
	if err != nil {
		resp.Code = e.Error
		resp.StoreDetail = repository.BuildStore(store)
		return resp, err
	}
	resp.StoreDetail = repository.BuildStore(store)
	return resp, err
}

// StoreDetailByIdentity 通过store_identity获取店铺信息
func (*StoreService) StoreDetailByIdentity(ctx context.Context, req *service.StoreRequest) (resp *service.StoreDetailResponse, err error) {
	var store repository.Store
	resp = new(service.StoreDetailResponse)
	resp.Code = e.Success
	store, err = store.StoreGetByIdentity(req)
	if err != nil {
		resp.Code = e.Error
		resp.StoreDetail = repository.BuildStore(store)
		return resp, err
	}
	resp.StoreDetail = repository.BuildStore(store)
	return resp, err
}

// StoreDetailModify 修改店铺信息
func (*StoreService) StoreDetailModify(ctx context.Context, req *service.StoreRequest) (resp *service.StoreDetailResponse, err error) {
	var store repository.Store
	resp = new(service.StoreDetailResponse)
	resp.Code = e.Success
	store, err = store.StoreDetailChange(req)
	if err != nil {
		resp.Code = e.Error
		resp.StoreDetail = repository.BuildStore(store)
		return resp, err
	}
	resp.StoreDetail = repository.BuildStore(store)
	return resp, err
}

// StoreDetailByStaffIdentity 通过员工identity查询店铺详细信息
func (*StoreService) StoreDetailByStaffIdentity(ctx context.Context, req *service.StaffRequest) (resp *service.StoreDetailResponse, err error) {
	var store repository.Store
	resp = new(service.StoreDetailResponse)
	resp.Code = e.Success
	store, err = store.StoreGetByStaffIdentity(req)
	if err != nil {
		resp.Code = e.Error
		resp.StoreDetail = repository.BuildStore(store)
		return resp, err
	}
	resp.StoreDetail = repository.BuildStore(store)
	return resp, err
}
