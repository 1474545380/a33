package repository

import (
	"errors"
	"gorm.io/gorm"
	"staff/internal/service"
	"staff/pkg/help"
)

type Staff struct {
	gorm.Model
	Identity      string `grom:"column:identity;type:longtext(36);" json:"identity"`             //员工表的唯一标识
	StaffName     string `grom:"column:staff_name;type:varchar(36);" json:"staff_name"`          //员工登录用户名
	Password      string `grom:"column:password;type:varchar(36);" json:"password"`              //员工登录密码
	Name          string `grom:"column:name;type:varchar(36);" json:"name"`                      //员工真实姓名
	Email         string `grom:"column:email;type:varchar(36);" json:"email"`                    //员工邮箱
	Position      string `grom:"column:position;type:varchar(36);" json:"position"`              //员工职位
	StoreIdentity string `grom:"column:store_identity;type:longtext(36);" json:"store_identity"` //店铺唯一标识
	User          *Store `gorm:"foreignKey:identity;references:store_identity"`                  //关联门店基础表
}

const (
	PasswordCost = 12 //加密难度
)

// 判断用户是否存在
func (staff *Staff) CheckStaffExist(req *service.StaffRequest) bool {
	password := help.GetMd5(req.Password)
	if err := DB.Where("staff_name = ? and password = ?", req.StaffName, password).First(&staff).Error; err == gorm.ErrRecordNotFound {
		return false
	}
	return true
}

// 获取用户信息
func (staff *Staff) ShowStaffInfo(req *service.StaffRequest) error {
	if exist := staff.CheckStaffExist(req); exist {
		return nil
	}
	return errors.New("StaffName Not Exist")
}

// 序列化Staff
func BuildStaff(item Staff) *service.StaffModel {
	staffModel := service.StaffModel{
		StoreIdentity: item.Identity,
		Name:          item.Name,
		StaffName:     item.StaffName,
	}
	return &staffModel
}
