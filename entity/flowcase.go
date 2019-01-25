package entity

import (
	"time"
)

// FlowCase 流程实例的数据
type FlowCase struct {
	CaseInfo  *Case               //`json:"case"`      //流程实例信息
	CaseItems map[int32]*CaseItem //`json:"caseitems"` //流程的步骤记录
}

// Case 流程实例对象
type Case struct {
	CaseID         string    `json:"caseid"` //实例id
	ItemID         int32     `json:"itemid"` //当前实例顺序
	AppID          string    `json:"appid"`
	BizID1         string    `json:"bizid1"`
	BizID2         string    `json:"bizid2"`
	FlowID         string    `json:"flowid"`         //流程定义id
	FlowName       string    `json:"flowname"`       //流程名称
	CreatorID      string    `json:"creatorid"`      //创建人id
	CreatorName    string    `json:"creatorname"`    //创建人名字
	Step           string    `json:"stepname"`       //当前状态
	Status         int32     `json:"status"`         //状态 0:审批中 1:通过 2:不通过
	CreateTime     time.Time `json:"createtime"`     //创建时间
	EndTime        time.Time `json:"endtime"`        //结束时间
	CopyUser       []int     `json:"copyuser"`       //抄送人
	AppData        string    `json:"appdata"`        //表单json数据
	SendTime       string    `json:"sendtime"`       //发送时间
	ChoiceItems    string    `json:"choiceitems"`    //审核选项
	SerialNumber   string    `json:"serialnumber"`   //流水号
	HandleUserid   string    `json:"handleuserid"`   //步骤处理人
	HandleUserName string    `json:"handleusername"` //处理人名字
	HandleTime     string    `json:"handletime"`     //处理时间
	StepStatus     int32     `json:"stepstatus"`     //处理状态 0未读1已读2已处理
	PluginID       string    `json:"pluginid"`
	VersionNo      int32     `json:"versionno"`
}

// CaseItem 流程实例的步骤数据
type CaseItem struct {
	ItemID         int32     `json:"itemid"`         //步骤id
	HandleUserid   string    `json:"handleuserid"`   //步骤处理人
	HandleUserName string    `json:"handleusername"` //处理人名字
	StepName       string    `json:"stepname"`       //这个步骤的状态名字
	Choice         string    `json:"choice"`         //用户的选择结果
	Mark           string    `json:"mark"`           //处理人的备注
	CreateTime     time.Time `json:"createtime"`     //发起时间
	HandleTime     string    `json:"handletime"`     //处理时间
	AgentUserid    string    `json:"agentuserid"`    //代理人id
	AgentUserName  string    `json:"agentusername"`  //代理人名字
	StepStatus     int32     `json:"stepstatus"`     //流程步骤的状态
	SysEnterInfo   string    `json:"sysenterinfo"`   //进入步骤,系统信息记录
	SysExitInfo    string    `json:"sysexitinfo"`    //离开步骤,系统信息记录
	ChoiceItems    string    `json:"choiceitems"`    //审核选项
}

// SetAgent 设置代理人
func (c *CaseItem) SetCaseItemAgent(userid, username string) {
	c.AgentUserid = userid
	c.AgentUserName = username
}

// CaseInfo DTO对象,流程事务
type CaseInfo struct {
	CaseID         string    `json:"caseid"`         //流程实例id
	ItemID         int32     `json:"itemid"`         //流程当前步骤id
	FlowID         string    `json:"flowid"`         //流程id
	Name           string    `json:"flowname"`       //流程名称
	Creator        string    `json:"creator"`        //流程发起人账号
	Creatorname    string    `json:"creatorname"`    //流程发起人姓名
	Createtime     time.Time `json:"createtime"`     //流程发起时间
	Handleuserid   string    `json:"handleuserid"`   //步骤原处理人(有代理人)
	Handleusername string    `json:"handleusername"` //步骤原处理人姓名
	Handletime     string    `json:"handletime"`     //处理时间
	ChoiceItems    string    `json:"choiceitems"`    //审核选项
	Stepname       string    `json:"stepname"`       //当前步骤名称
	Stepstatus     int32     `json:"stepstatus"`     //当前步骤的状态，0:未处理 1:已读 2:已处理
	Status         int32     `json:"status"`         //状态,0:审批中 1:通过 2:不通过
	Appid          string    `json:"appid"`          //流程关联的业务对象(记录在crm_t_entityreg)
	Bizid1         string    `json:"bizid1"`         //业务主键1
	Bizid2         string    `json:"bizid2"`         //业务主键2
	SendTime       string    `json:"sendtime"`       //发送时间
	SerialNumber   string    `json:"serialnumber"`   //流水号
	Choice         string    `json:"choice"`         //审核
	FlowStatus     int32     `json:"flowstatus"`     //流程状态 1启用0停用
}

// CaseList 代办事务
type CaseList struct {
	Items      []*CaseInfo
	TotalItems int32
}

// FlowCaseList 流程实例信息
type FlowCaseList struct {
	CaseInfo  *Case       //`json:"case"`      //流程实例信息
	CaseItems []*CaseItem //`json:"caseitems"` //流程的步骤记录
}

// NextStatuInfo 流程下一步的信息
type NextStatuInfo struct {
	// StepName 步骤名称
	StepName string
	// Users 流程参与人
	Users []*FlowUser
	// IsFree 是否自由流程
	IsFree bool
	// SelectType 如何选择人员 0表示只能选Users数组内的人 1表示可以选择所有的人
	SelectType bool
}
