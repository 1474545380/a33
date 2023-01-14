package handler

import (
	"context"
	"schedul_rule/internal/repository"
	"schedul_rule/internal/service"
	pb "schedul_rule/internal/service"
	"schedul_rule/pkg/e"
)

type SchedulRuleService struct {
	pb.UnimplementedSchedulRuleServiceServer
}

func NewStaffService() *SchedulRuleService {
	return &SchedulRuleService{}
}

// SchedulRule 店铺排班规则获取全部
func (*SchedulRuleService) SchedulRule(ctx context.Context, req *pb.SchedulRuleRequest) (resp *service.SchedulRuleResponse, err error) {
	var schedulRules repository.SchedulRules
	resp = new(service.SchedulRuleResponse)
	resp.Code = e.Success
	schedulRulesDetail, err := schedulRules.SchedulRuleGet(req)
	if err != nil {
		resp.Code = e.Error
		resp.SchedulRuleDetail = repository.BuildSchedulDetails(schedulRulesDetail)
		return resp, err
	}
	resp.SchedulRuleDetail = repository.BuildSchedulDetails(schedulRulesDetail)
	return resp, err
}

// SchedulRuleInsert 店铺排班规则添加
func (*SchedulRuleService) SchedulRuleInsert(ctx context.Context, req *pb.SchedulRuleRequest) (resp *service.SchedulRuleResponse, err error) {
	var schedulRules repository.SchedulRules
	resp = new(service.SchedulRuleResponse)
	resp.Code = e.Success
	schedulRulesDetail, err := schedulRules.SchedulRuleInsert(req)
	if err != nil {
		resp.Code = e.Error
		resp.SchedulRuleDetail = append(resp.SchedulRuleDetail, repository.BuildSchedulDetail(schedulRulesDetail))
		return resp, err
	}
	resp.SchedulRuleDetail = append(resp.SchedulRuleDetail, repository.BuildSchedulDetail(schedulRulesDetail))
	return resp, err
}

// SchedulRuleModify 店铺排班规则修改
func (*SchedulRuleService) SchedulRuleModify(ctx context.Context, req *pb.SchedulRuleRequest) (resp *service.SchedulRuleResponse, err error) {
	var schedulRules repository.SchedulRules
	resp = new(service.SchedulRuleResponse)
	resp.Code = e.Success
	schedulRulesDetail, err := schedulRules.SchedulRuleChange(req)
	if err != nil {
		resp.Code = e.Error
		resp.SchedulRuleDetail = append(resp.SchedulRuleDetail, repository.BuildSchedulDetail(schedulRulesDetail))
		return resp, err
	}
	resp.SchedulRuleDetail = append(resp.SchedulRuleDetail, repository.BuildSchedulDetail(schedulRulesDetail))
	return resp, err
}

// SchedulRuleDelete 店铺排班规则删除
func (*SchedulRuleService) SchedulRuleDelete(ctx context.Context, req *pb.SchedulRuleRequest) (resp *service.SchedulRuleResponse, err error) {
	var schedulRules repository.SchedulRules
	resp = new(service.SchedulRuleResponse)
	resp.Code = e.Success
	schedulRulesDetail, err := schedulRules.SchedulRuleDetail(req)
	if err != nil {
		resp.Code = e.Error
		resp.SchedulRuleDetail = append(resp.SchedulRuleDetail, repository.BuildSchedulDetail(schedulRulesDetail))
		return resp, err
	}
	resp.SchedulRuleDetail = append(resp.SchedulRuleDetail, repository.BuildSchedulDetail(schedulRulesDetail))
	return resp, err
}

// SchedulRuleByStoreIdentity 通过store_identity查询店铺规则，会同时查询出store_identity为null的字段
func (*SchedulRuleService) SchedulRuleByStoreIdentity(ctx context.Context, req *pb.SchedulRuleRequest) (resp *service.SchedulRuleResponse, err error) {
	var schedulRules repository.SchedulRules
	resp = new(service.SchedulRuleResponse)
	resp.Code = e.Success
	schedulRulesDetail, err := schedulRules.SchedulRuleGetByStoreIdentity(req)
	if err != nil {
		resp.Code = e.Error
		resp.SchedulRuleDetail = repository.BuildSchedulDetails(schedulRulesDetail)
		return resp, err
	}
	resp.SchedulRuleDetail = repository.BuildSchedulDetails(schedulRulesDetail)
	return resp, err
}
