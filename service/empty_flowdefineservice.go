package service

import (
	. "github.com/pobearm/workflow/entity"
)

var _ FlowDefineService = new(EmptyFlowDefineService)

type EmptyFlowDefineService struct {
}

func NewEmptyFlowDefineService(connstr string) (*EmptyFlowDefineService, error) {
	p := &EmptyFlowDefineService{}
	return p, nil
}

// GetWorkFlows 获取流程列表, 分页
func (s *EmptyFlowDefineService) GetWorkFlows(pageindex, pagesize int32) (*FlowList, error) {
	return &FlowList{}, nil
}

// GetWorkFlowDetail 获取指定流程的详情
func (s *EmptyFlowDefineService) GetWorkFlowDetail(flowid string) (*FlowInfo, error) {
	return &FlowInfo{}, nil
}

//保存一个新的流程定义
func (s *EmptyFlowDefineService) AddFlow(flow *FlowInfo, appid string) error {
	return nil
}

//删除一个流程定义
func (s *EmptyFlowDefineService) DeleteFlow(flowid string) error {
	return nil
}

//修改一个流程定义
func (s *EmptyFlowDefineService) UpdateFlow(flow *FlowInfo) error {
	return nil
}

//启用流程
func (s *EmptyFlowDefineService) EnableFlow(flow *FlowInfo) error {
	return nil
}

//停用流程
func (s *EmptyFlowDefineService) DisableFlow(flow *FlowInfo) error {
	return nil
}
