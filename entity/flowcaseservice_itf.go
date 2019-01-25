package entity

// FlowCaseService 流程实例服务，接口定义
type FlowCaseService interface {

	// GetTodoCases 获取用户的代办列表---flowname查询条件
	GetTodoCases(flowname, usernumber string, pageindex, pagesize int32) (*CaseList, error)

	// LoadFlowCase 加载一个流程的完整信息
	LoadFlowCase(caseid string) (fc *FlowCase, err error)

	// SaveNewCase 保存一个新的流程实例
	SaveNewCase(fc *FlowCase, versionno int32) error

	// ComitFlow 在一个事务中提交流程数据,case,当前item,下一步item
	ComitFlow(c *Case, ci *CaseItem, ni *CaseItem) error

	// FindAgent 找到步骤处理人的代理人
	FindAgent(userid string) (user *FlowUser, find bool)

	// StepHandled 记录步骤进入,退出的消息
	StepHandled(ca *Case, ci *CaseItem, ni *CaseItem) error

	// GetCaseDetail 流程实例详情
	GetCaseDetail(caseid string) (*FlowCaseList, error)

	// PreCommitCase 预提交, 选择审批选项, 返回下一步去到的步骤和可选审批人
	PreCommitCase(caseid, choice string, itemid int32, appdata string) (nsif *NextStatuInfo, err error)

	// CommitCase 处理待办项, 返回进入的状态名称
	CommitCase(enterprise, usernumber, caseid, choice, remark string, itemid int32, flowuser *FlowUser, appdata string) (string, error)

	//作废流程实列
	AbandonCase(enterprise, usernumber, caseid, choice, remark string, itemid int32, appdata string) error

	//结束流程实列
	FinishCase(caseid, choice, remark string, itemid int32, appdata string) error

	//流程实列, 退回到发起人
	SendbackCase(enterprise, usernumber, caseid, choice, remark string, itemid int32, appdata string) error

	//流程实列, 退回给上一个步骤
	FallbackCase(caseid, choice, remark string, itemid int32, appdata string) error

	//标记流程步骤为已读
	Readed(itemid int32, caseid, usernumber string) error

	//设置代理人
	SetAgent(userid, agentid string) error

	//取消代理人
	UnsetAgent(userid string) error
}
