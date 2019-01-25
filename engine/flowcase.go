package engine

import (
	. "github.com/pobearm/workflow/entity"
	"time"
)

// NewStartFlowCase 新启动一个流程实例
func NewStartFlowCase(creatorid, creatorname, caseid, flowid, step string, versionno int32) *FlowCase {
	c := &Case{
		CaseID:      caseid,
		FlowID:      flowid,
		CreatorID:   creatorid,
		CreatorName: creatorname,
		Step:        step,
		Status:      0,
		CreateTime:  time.Now(),
		VersionNo:   versionno,
	}
	ci := NewCaseItem(0, step, creatorid, creatorname)
	cis := make(map[int32]*CaseItem)
	cis[ci.ItemID] = ci
	return &FlowCase{CaseInfo: c, CaseItems: cis}
}

//流程步骤的状态
const (
	StepStatusNew int32 = iota
	StepStatusRead
	StepStatusFinish
)

// NewCaseItem 新创建步骤流转信息
func NewCaseItem(itemid int32, stepname, userid, username string) *CaseItem {
	return &CaseItem{
		ItemID:         itemid,
		StepName:       stepname,
		HandleUserid:   userid,
		HandleUserName: username,
		CreateTime:     time.Now(),
		StepStatus:     StepStatusNew,
	}
}
