package repository

import "gorm.io/gorm"

type Staff struct {
	gorm.Model
	Identity      string `grom:"column:identity;type:longtext(36);" json:"identity"`             //员工表的唯一标识
	Name          string `grom:"column:name;type:varchar(36);" json:"name"`                      //员工姓名
	Email         string `grom:"column:email;type:varchar(36);" json:"email"`                    //员工邮箱
	Position      string `grom:"column:position;type:varchar(36);" json:"position"`              //员工职位
	StoreIdentity string `grom:"column:store_identity;type:longtext(36);" json:"store_identity"` //店铺唯一标识
	User          *Store `gorm:"foreignKey:identity;references:store_identity"`                  //关联门店基础表
}