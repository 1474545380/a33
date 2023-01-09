package repository

import "gorm.io/gorm"

type Store struct {
	gorm.Model
	Identity string `grom:"column:identity;type:longtext(36);" json:"identity"` //员工表的唯一标识
	Name     string `grom:"column:name;type:varchar(100);" json:"name"`         //员工姓名
	Address  string `grom:"column:address;type:varchar(255);" json:"address"`   //门店地址
	Size     uint64 `grom:"column:size;type:int(100);" json:"size"`             //平方米
}
