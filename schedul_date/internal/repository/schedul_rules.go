package repository

import "gorm.io/gorm"

type SchedulRules struct {
	gorm.Model
	Identity      string `gorm:"column:identity;type:longtext(36);" json:"identity"`         //店铺排班规则表的唯一标识
	RuleType      string `gorm:"column:rule_type;type:varchar(255);" json:"rule_type"`       //可选值:开店规则、关店规则、客流规则。可扩展
	StoreIdentity string `gorm:"column:store_identity;type:longtext;" json:"store_identity"` //店铺唯一标识，为空时，为系统通用规则。 不为空时，为门店规则。当门店有门店规则时，使用门店规则进行排班，没有门店规则时，使用系统通用规则进行排班
	RuleDate      string `gorm:"column:rule_date;type:text;" json:"rule_date"`               //需要开发者自行设计。 示例:- 客流规则:"3.8" 表示按照业务预测数据，每 3.8 个客流必须安排至少一个员工当值- 开店规则:"1.5,23.5" 表示开店 1 个半小时前需要 有员工当值，当值员工数为门店面积除以 23.5- 关店规则:"2.5,3,13" 表示关店 2 个半小时内需要 有员工当值，当值员工数不小于 3 并且不小于门店 面积除以 13为了提高规则的灵活性，建议使用 json 格式保存规 则值，如关店规则: {"after":"2.5","count":"3","fomula":"size/13"}
	Store         *Store `gorm:"foreignKey:identity;references:store_identity"`              //关联门店基础表
}
