package repository

import "gorm.io/gorm"

type Store struct {
	gorm.Model
	Identity string `gorm:"column:identity;type:longtext;" json:"identity"`   //员工表的唯一标识
	Name     string `gorm:"column:name;type:varchar(100);" json:"name"`       //员工姓名
	Address  string `gorm:"column:address;type:varchar(255);" json:"address"` //门店地址
	Size     uint64 `gorm:"column:size;type:int(100);" json:"size"`           //平方米
	Staffs   *Staff `gorm:"foreignKey:store_identity;references:identity;"`   //联表
}
