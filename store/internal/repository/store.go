package repository

import (
	"errors"
	"gorm.io/gorm"
	"store/internal/service"
	"store/pkg/help"
)

type Store struct {
	gorm.Model
	Identity string `grom:"column:identity;type:longtext(36);" json:"identity"` //员工表的唯一标识
	Name     string `grom:"column:name;type:varchar(100);" json:"name"`         //员工姓名
	Address  string `grom:"column:address;type:varchar(255);" json:"address"`   //门店地址
	Size     uint64 `grom:"column:size;type:int(100);" json:"size"`             //平方米
}

// StoreCreat 添加店铺
func (s Store) StoreCreat(req *service.StoreRequest) (store Store, err error) {
	if req.Name == "" || req.Size == 0 || req.Address == "" {
		return Store{}, errors.New("InvalidParams")
	}
	//var count int64
	//err = DB.Where("name = ?", req.Name).Count(&count).Error
	//if err != nil {
	//	return Store{}, errors.New("creat error")
	//}
	//if count != 0 {
	//	return Store{}, errors.New("store is exist")
	//}
	storeIdentity := help.GetUUID()
	store = Store{
		Identity: storeIdentity,
		Name:     req.Name,
		Address:  req.Address,
		Size:     req.Size,
	}
	err = DB.Create(&store).Error
	if err != nil {
		return Store{}, errors.New("store creat error")
	}
	return store, err
}

// StoreDetailChange 店铺信息修改
func (s Store) StoreDetailChange(req *service.StoreRequest) (Store, error) {
	if req.Identity == "" || req.Name == "" || req.Size == 0 || req.Address == "" {
		return Store{}, errors.New("InvalidParams")
	}
	store := Store{}
	DB.Where("identity = ?", req.Identity).Take(&store)
	store.Address = req.Address
	store.Size = req.Size
	store.Name = req.Name
	err := DB.Save(&store).Error
	if err != nil {
		return Store{}, errors.New("store update error")
	}
	return store, err
}

// StoreGetByIdentity 通过identity获取店铺信息
func (s Store) StoreGetByIdentity(req *service.StoreRequest) (Store, error) {
	if req.Identity == "" {
		return Store{}, errors.New("InvalidParams")
	}
	storeDetail := new(Store)
	err := DB.Where("identity = ?", req.Identity).Find(&storeDetail).Error
	if err != nil {
		return Store{}, err
	}
	return *storeDetail, err
}

// BuildStore 序列化Store
func BuildStore(item Store) *service.StoreModel {
	storeModel := service.StoreModel{
		Identity: help.GetUUID(),
		Name:     item.Name,
		Address:  item.Address,
		Size:     item.Size,
	}
	return &storeModel
}
