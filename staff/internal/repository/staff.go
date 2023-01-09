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
	Password      string `grom:"column:password;type:varchar(36);" json:"password"`              //员工登录密码
	Name          string `grom:"column:name;type:varchar(36);" json:"name"`                      //员工真实姓名
	Email         string `grom:"column:email;type:varchar(36);" json:"email"`                    //员工邮箱
	Position      string `grom:"column:position;type:varchar(36);" json:"position"`              //员工职位
	StoreIdentity string `grom:"column:store_identity;type:longtext(36);" json:"store_identity"` //店铺唯一标识
	Store         *Store `gorm:"foreignKey:identity;references:store_identity"`                  //关联门店基础表
}

// CheckStaffExist 判断员工登录信息是否正确
func (staff *Staff) CheckStaffExist(req *service.StaffRequest) bool {
	password := help.GetMd5(req.Password)
	if err := DB.Where("staff_name = ? and password = ?", req.Email, password).First(&staff).Error; err == gorm.ErrRecordNotFound {
		return false
	}
	return true
}

// ShowStaffInfo 获取员工信息
func (staff *Staff) ShowStaffInfo(req *service.StaffRequest) error {
	if exist := staff.CheckStaffExist(req); exist {
		return nil
	}
	return errors.New("staff not exist")
}

// StaffCreat 创建员工
func (*Staff) StaffCreat(req *service.StaffRequest) (staff Staff, err error) {
	var count int64
	err = DB.Where("email = ?", req.Email).Count(&count).Error
	if err != nil {
		return Staff{}, errors.New("creat error")
	}
	if count != 0 {
		return Staff{}, errors.New("staff is exist")
	}
	staffIdentity := help.GetUUID()
	staff = Staff{
		Identity:      staffIdentity,
		Password:      help.GetMd5(req.Password),
		Name:          req.Name,
		Email:         req.Email,
		Position:      req.Position,
		StoreIdentity: req.StoreIdentity,
	}
	err = DB.Create(&staff).Error
	return staff, err
}

// StaffDetailGet 获取员工具体信息
func (staff *Staff) StaffDetailGet(req *service.StaffRequest) (Staff, error) {
	staffDetail := new(Staff)
	err := DB.Omit("password").Where("identity = ?", req.Identity).Find(&staffDetail).Error
	if err != nil {
		return Staff{}, err
	}
	return *staffDetail, err
}

// StaffDetailChange 修改员工具体信息
func (staff *Staff) StaffDetailChange(req *service.StaffRequest) error {
	if req.Email == "" || req.Identity == "" || req.Name == "" || req.StoreIdentity == "" || req.Position == "" {
		return errors.New("incomplete information")
	}
	err := DB.Omit("password").Where("identity = ?", req.Identity).Updates(req).Error
	if err != nil {
		return errors.New("data change error")
	}
	return nil
}

// BuildStaff 序列化Staff
func BuildStaff(item Staff) *service.StaffModel {
	staffModel := service.StaffModel{
		StoreIdentity: item.Identity,
		Name:          item.Name,
	}
	return &staffModel
}
