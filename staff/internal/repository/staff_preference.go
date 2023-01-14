package repository

import (
	"errors"
	"gorm.io/gorm"
	"staff/internal/service"
)

type StaffPreference struct {
	gorm.Model
	StaffPreferenceIdentity string `gorm:"column:staff_preference_identity;type:longtext(255);" json:"staff_preference_identity"` //员工偏好唯一标识
	PreferenceType          string `grom:"column:preference_type;type:varchar(255);" json:"preference_type"`                      //员工偏好类型，可选值:工作日偏好、工作时间偏好、班次时长偏好。可扩展
	StaffIdentity           string `gorm:"column:staff_identity;type:longtext(255);" json:"staff_identity"`                       //员工唯一标识，当员工没有选择对应的偏好类型时，表示员工对该 类型无特殊偏好。
	PreferenceValue         string `gorm:"column:preference_value;type:text(255);" json:"preference"`                             //员工偏好值需要开发者自行设计。 示例:- 工作日偏好的值:1,3,4 表示偏好周一、周 三、周四工作。- 工作时间偏好的值:08:00-12:00,18:00:22:00 表 示偏好上午 9 点到 12 点和晚上 6 点到 10 点工作
	Staff                   *Staff `gorm:"foreignKey:staff_identity;references:identity"`                                         //关联员工
}

// StaffPreferenceGet 获取员工偏好信息
func (staffPreference *StaffPreference) StaffPreferenceGet(req *service.StaffRequest) ([]StaffPreference, error) {
	if req.Identity == "" {
		return []StaffPreference{}, errors.New("InvalidParams")
	}
	var staffPreferenceDetail []StaffPreference
	err := DB.Model(&StaffPreference{}).Where("staff_identity = ?", req.Identity).Find(&staffPreferenceDetail).Error
	if err != nil {
		return []StaffPreference{}, errors.New("data find error")
	}
	return staffPreferenceDetail, err
}

// StaffPreferenceAdd 添加员工偏好信息
func (staffPreference *StaffPreference) StaffPreferenceAdd(req *service.StaffPreferenceRequest) (StaffPreference, error) {
	if req.StaffIdentity == "" || req.StaffPreferenceIdentity == "" || req.PreferenceType == "" || req.PreferenceValue == "" {
		return StaffPreference{}, errors.New("InvalidParams")
	}
	staffPreferenceDetail := StaffPreference{
		StaffPreferenceIdentity: req.StaffPreferenceIdentity,
		PreferenceType:          req.PreferenceType,
		StaffIdentity:           req.StaffIdentity,
		PreferenceValue:         req.PreferenceValue,
	}
	err := DB.Model(&StaffPreference{}).Create(&staffPreferenceDetail).Error
	if err != nil {
		return StaffPreference{}, errors.New("data find error")
	}
	return StaffPreference{}, err
}

// StaffPreferenceChange 修改员工偏好信息
func (staffPreference *StaffPreference) StaffPreferenceChange(req *service.StaffPreferenceRequest) (StaffPreference, error) {
	if req.StaffPreferenceIdentity == "" {
		return StaffPreference{}, errors.New("InvalidParams")
	}
	s := StaffPreference{}
	err := DB.Model(&StaffPreference{}).Where("staff_preference_identity = ?", req.StaffPreferenceIdentity).Take(&s).Error
	s.StaffIdentity = req.StaffIdentity
	s.PreferenceType = req.PreferenceType
	s.PreferenceValue = req.PreferenceValue
	DB.Model(&StaffPreference{}).Save(&s)
	if err != nil {
		return StaffPreference{}, errors.New("data update error")
	}
	return StaffPreference{}, err
}

// BuildStaffPreference 序列化[]StaffPreference
func BuildStaffPreference(item []StaffPreference) []*service.StaffPreferenceModel {
	staffPreferenceModel := make([]*service.StaffPreferenceModel, 0)
	for _, i := range item {
		staffPreferenceModel = append(staffPreferenceModel, &service.StaffPreferenceModel{
			StaffPreferenceIdentity: i.StaffPreferenceIdentity,
			PreferenceType:          i.PreferenceType,
			StaffIdentity:           i.StaffIdentity,
			PreferenceValue:         i.PreferenceValue,
		})
	}
	return staffPreferenceModel
}

// BuildNormalStaffPreference 序列化staffpreference
func BuildNormalStaffPreference(item StaffPreference) *service.StaffPreferenceModel {
	staffPreferenceModel := service.StaffPreferenceModel{
		StaffPreferenceIdentity: item.StaffPreferenceIdentity,
		PreferenceType:          item.PreferenceType,
		StaffIdentity:           item.StaffIdentity,
		PreferenceValue:         item.PreferenceValue,
	}
	return &staffPreferenceModel

}
