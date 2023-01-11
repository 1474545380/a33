package repository

import "gorm.io/gorm"

type SchedulRules struct {
	gorm.Model
	Identity      string `grom:"column:identity;type:longtext(36);" json:"identity"`             //员工表的唯一标识
	RuleType      string `grom:"column:rule_type;type:varchar(255);" json:"rule_type"`           //员工登录密码
	StoreIdentity string `grom:"column:store_identity;type:longtext(36);" json:"store_identity"` //店铺唯一标识
	RuleDate      string `grom:"column:rule_date;type:text;" json:"rule_date"`                   //员工真实姓名
	Store         *Store `gorm:"foreignKey:identity;references:store_identity"`                  //关联门店基础表
}
