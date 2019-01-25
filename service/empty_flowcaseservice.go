package service

import (
	. "github.com/pobearm/workflow/entity"
)

var _ FlowCaseService = new(EmptyFlowCaseService)

type EmptyFlowCaseService struct {
}

func NewEmptyFlowCaseService(connstr string) (*EmptyFlowCaseService, error) {
	p := &EmptyFlowCaseService{}
	return p, nil
}

// GetTodoCases 获取用户的代办列表---flowname查询条件
func (s *EmptyFlowCaseService) GetTodoCases(flowname, usernumber string, pageindex, pagesize int32) (*CaseList, error) {
	return &CaseList{}, nil
}

// LoadFlowCase 加载一个流程的完整信息
func (s *EmptyFlowCaseService) LoadFlowCase(caseid string) (fc *FlowCase, err error) {
	return &FlowCase{}, nil
}

// SaveNewCase 保存一个新的流程实例
func (s *EmptyFlowCaseService) SaveNewCase(fc *FlowCase, versionno int32) error {
	return nil
}

// ComitFlow 在一个事务中提交流程数据,case,当前item,下一步item
func (s *EmptyFlowCaseService) ComitFlow(c *Case, ci *CaseItem, ni *CaseItem) error {
	return nil
}

// FindAgent 找到步骤处理人的代理人
func (s *EmptyFlowCaseService) FindAgent(userid string) (user *FlowUser, find bool) {
	return &FlowUser{}, false
}

// StepHandled 记录步骤进入,退出的消息
func (s *EmptyFlowCaseService) StepHandled(ca *Case, ci *CaseItem, ni *CaseItem) error {
	return nil
}

// GetCaseDetail 流程实例详情
func (s *EmptyFlowCaseService) GetCaseDetail(caseid string) (*FlowCaseList, error) {
	return &FlowCaseList{}, nil
}

// PreCommitCase 预提交, 选择审批选项, 返回下一步去到的步骤和可选审批人
func (s *EmptyFlowCaseService) PreCommitCase(caseid, choice string, itemid int32, appdata string) (nsif *NextStatuInfo, err error) {
	return &NextStatuInfo{}, nil
}

// CommitCase 处理待办项, 返回进入的状态名称
func (s *EmptyFlowCaseService) CommitCase(enterprise, usernumber, caseid, choice, remark string,
	itemid int32, flowuser *FlowUser, appdata string) (stepname string, err error) {
	return "", nil
}

//作废流程实列
func (s *EmptyFlowCaseService) AbandonCase(enterprise, usernumber, caseid, choice, remark string,
	itemid int32, appdata string) error {
	return nil
}

//结束流程实列
func (s *EmptyFlowCaseService) FinishCase(caseid, choice, remark string, itemid int32, appdata string) error {
	return nil
}

//流程实列, 退回到发起人
func (s *EmptyFlowCaseService) SendbackCase(enterprise, usernumber, caseid, choice, remark string,
	itemid int32, appdata string) error {
	return nil
}

//流程实列, 退回给上一个步骤
func (s *EmptyFlowCaseService) FallbackCase(caseid, choice, remark string, itemid int32, appdata string) error {
	return nil
}

//标记流程步骤为已读
func (s *EmptyFlowCaseService) Readed(itemid int32, caseid, usernumber string) error {
	return nil
}

//设置代理人
func (s *EmptyFlowCaseService) SetAgent(userid, agentid string) error {
	return nil
}

//取消代理人
func (s *EmptyFlowCaseService) UnsetAgent(userid string) error {
	return nil
}
