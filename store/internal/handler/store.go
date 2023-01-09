package handler

import (
	"context"
	"store/internal/repository"
	"store/internal/service"
	"store/pkg/e"
)

type StoreService struct {
}

func NewStoreService() *StoreService {
	return &StoreService{}
}

// StoreFill 添加新商店
func (*StoreService) StoreFill(ctx context.Context, req *service.StoreRequest) (resp *service.StoreDetailResponse, err error) {
	var store repository.Store
	resp = new(service.StoreDetailResponse)
	resp.Code = e.Success
	store, err = store.StoreCreat(req)
	if err != nil {
		resp.Code = e.Error
		return resp, err
	}
	resp.StoreDetail = repository.BuildStore(store)
	return resp, err
}

// StoreDetailByIdentity StoreDetail 通过store_identity获取商店信息
func (*StoreService) StoreDetailByIdentity() {

}

// StoreDetailModify 修改商店信息
func (*StoreService) StoreDetailModify() {

}
