package entity

// FlowDefineService 流程定义服务接口
type FlowDefineService interface {
	// GetWorkFlows 获取流程列表, 可按状态, 名称过滤, 分页
	GetWorkFlows(pageindex, pagesize int32) (*FlowList, error)

	// GetWorkFlowDetail 获取指定流程的详情
	GetWorkFlowDetail(flowid string) (*FlowInfo, error)

	// AddFlow 保存一个新的流程定义, appid 与业务功能关联的id
	AddFlow(flow *FlowInfo, appid string) error

	// DeleteFlow 删除一个流程定义
	DeleteFlow(flowid string) error

	// UpdateFlow 修改一个流程定义
	UpdateFlow(flow *FlowInfo) error

	// EnableFlow 启用流程
	EnableFlow(flow *FlowInfo) error

	// DisableFlow 停用流程
	DisableFlow(flow *FlowInfo) error
}
