package repository

import (
	"errors"
	"gorm.io/gorm"
	"log"
	"schedul_rule/internal/service"
	"schedul_rule/pkg/help"
)

type SchedulRules struct {
	gorm.Model
	Identity      string `gorm:"column:identity;type:longtext(36);" json:"identity"`         //店铺排班规则表的唯一标识
	RuleType      string `gorm:"column:rule_type;type:varchar(255);" json:"rule_type"`       //可选值:开店规则、关店规则、客流规则。可扩展
	StoreIdentity string `gorm:"column:store_identity;type:longtext;" json:"store_identity"` //店铺唯一标识，为空时，为系统通用规则。 不为空时，为门店规则。当门店有门店规则时，使用门店规则进行排班，没有门店规则时，使用系统通用规则进行排班
	RuleDate      string `gorm:"column:rule_date;type:text;" json:"rule_date"`               //需要开发者自行设计。 示例:- 客流规则:"3.8" 表示按照业务预测数据，每 3.8 个客流必须安排至少一个员工当值- 开店规则:"1.5,23.5" 表示开店 1 个半小时前需要 有员工当值，当值员工数为门店面积除以 23.5- 关店规则:"2.5,3,13" 表示关店 2 个半小时内需要 有员工当值，当值员工数不小于 3 并且不小于门店 面积除以 13为了提高规则的灵活性，建议使用 json 格式保存规 则值，如关店规则: {"after":"2.5","count":"3","fomula":"size/13"}
	Store         *Store `gorm:"foreignKey:identity;references:store_identity"`              //关联门店基础表
}

// SchedulRuleGet 获取所有店铺规则
func (schedulRules *SchedulRules) SchedulRuleGet(req *service.SchedulRuleRequest) ([]SchedulRules, error) {
	var scheduDetail []SchedulRules
	err := DB.Model(&SchedulRules{}).Find(&scheduDetail).Error
	if err != nil {
		return []SchedulRules{}, errors.New("data find error")
	}
	return scheduDetail, err
}

// SchedulRuleInsert 店铺排班规则添加
func (schedulRules *SchedulRules) SchedulRuleInsert(req *service.SchedulRuleRequest) (SchedulRules, error) {
	if req.RuleType == "" || req.RuleDate == "" {
		return SchedulRules{}, errors.New("InvalidParams")
	}
	schedulRuleIdentity := help.GetUUID()
	s := SchedulRules{
		Identity:      schedulRuleIdentity,
		RuleDate:      req.RuleDate,
		RuleType:      req.RuleType,
		StoreIdentity: req.StoreIdentity,
	}
	err := DB.Model(&SchedulRules{}).Create(&s).Error
	if err != nil {
		return SchedulRules{}, errors.New("schedulRuleInsert error")
	}
	return s, err
}

// SchedulRuleChange 店铺排版规则修改
func (schedulRules *SchedulRules) SchedulRuleChange(req *service.SchedulRuleRequest) (SchedulRules, error) {
	if req.Identity == "" || req.RuleType == "" || req.RuleDate == "" {
		return SchedulRules{}, errors.New("InvalidParams")
	}
	s := SchedulRules{}
	err := DB.Model(&SchedulRules{}).Where("identity = ?", req.Identity).Take(&s).Error
	s.RuleType = req.RuleType
	s.StoreIdentity = req.StoreIdentity
	s.RuleDate = req.RuleDate
	err = DB.Model(&SchedulRules{}).Save(&s).Error
	if err != nil {
		return SchedulRules{}, errors.New("schedulRuleInsert error")
	}
	return s, err
}

// SchedulRuleDetail 店铺排班规则删除
func (schedulRules *SchedulRules) SchedulRuleDetail(req *service.SchedulRuleRequest) (SchedulRules, error) {
	if req.Identity == "" {
		return SchedulRules{}, errors.New("InvalidParams")
	}
	s := SchedulRules{}
	err := DB.Model(&SchedulRules{}).Where("identity = ?", req.Identity).Find(&s).Error
	err = DB.Model(&Staff{}).Delete(&s).Error
	if err != nil {
		return SchedulRules{}, errors.New("data delete error")
	}
	return s, err
}

// SchedulRuleGetByStoreIdentity 通过store_identity获取店铺排班规则
func (schedulRules *SchedulRules) SchedulRuleGetByStoreIdentity(req *service.SchedulRuleRequest) ([]SchedulRules, error) {
	if req.StoreIdentity == "" {
		return []SchedulRules{}, errors.New("InvalidParams")
	}
	log.Println(req.StoreIdentity)
	var scheduDetail []SchedulRules
	err := DB.Model(&SchedulRules{}).Where("store_identity = ? or store_identity = ?", req.StoreIdentity, "").Find(&scheduDetail).Error
	if err != nil {
		return []SchedulRules{}, errors.New("data find error")
	}
	return scheduDetail, err
}

func BuildSchedulDetails(item []SchedulRules) []*service.SchedulRuleModel {
	schedulRuleModel := make([]*service.SchedulRuleModel, 0)
	for _, i := range item {
		schedulRuleModel = append(schedulRuleModel, &service.SchedulRuleModel{
			Identity:      i.Identity,
			RuleType:      i.RuleType,
			StoreIdentity: i.StoreIdentity,
			RuleDate:      i.RuleDate,
		})
	}
	return schedulRuleModel
}

func BuildSchedulDetail(item SchedulRules) *service.SchedulRuleModel {
	schedulModel := service.SchedulRuleModel{
		Identity:      item.Identity,
		RuleType:      item.RuleType,
		RuleDate:      item.RuleDate,
		StoreIdentity: item.StoreIdentity,
	}
	return &schedulModel
}
