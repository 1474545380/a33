package repository

import (
	"gorm.io/gorm"
)

type Staff struct {
	gorm.Model
	Identity      string `gorm:"column:identity;type:longtext(36);" json:"identity"`             //员工表的唯一标识
	Password      string `gorm:"column:password;type:varchar(36);" json:"password"`              //员工登录密码
	Name          string `gorm:"column:name;type:varchar(36);" json:"name"`                      //员工真实姓名
	Email         string `gorm:"column:email;type:varchar(36);" json:"email"`                    //员工邮箱
	Position      string `gorm:"column:position;type:varchar(36);" json:"position"`              //员工职位
	StoreIdentity string `gorm:"column:store_identity;type:longtext(36);" json:"store_identity"` //店铺唯一标识
}
