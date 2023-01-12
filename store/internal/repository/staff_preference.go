package repository

import "gorm.io/gorm"

type StaffPreference struct {
	gorm.Model
	PreferenceType  string `gorm:"column:preference_type;type:varchar(255);" json:"preference_type"` //员工偏好类型，可选值:工作日偏好、工作时间偏好、班次时长偏好。可扩展
	StaffIdentity   string `gorm:"column:staff_identity;type:longtext(255);" json:"staff_identity"`  //员工唯一标识，当员工没有选择对应的偏好类型时，表示员工对该 类型无特殊偏好。
	PreferenceValue string `gorm:"column:preference_value;type:text(255);" json:"preference"`        //员工偏好值需要开发者自行设计。 示例:- 工作日偏好的值:1,3,4 表示偏好周一、周 三、周四工作。- 工作时间偏好的值:08:00-12:00,18:00:22:00 表 示偏好上午 9 点到 12 点和晚上 6 点到 10 点工作
	Staff           *Staff `gorm:"foreignKey:staff_identity;references:identity"`                    //关联员工
}
