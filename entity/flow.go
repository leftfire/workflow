package entity

import (
	"time"
)

// FlowInfo 流程的信息
type FlowInfo struct {
	FlowID     string    `json:"flowid"`
	Name       string    `json:"flowname"`
	Descript   string    `json:"descript"`
	FlowXML    string    `json:"flowxml"`
	StepCount  int32     `json:"stepcount"`
	CreateTime time.Time `json:"createtime"`
	Creator    string    `json:"creator"`
	Status     int32     `json:"status"`
	UpdateTime string    `json:"updatetime"`
	Updator    string    `json:"updator"`
	FlowType   int32     `json:"flowtype"`
	AppID      string    `json:"appid"`
}

// FlowList 流程列表
type FlowList struct {
	Items      []*FlowInfo
	TotalItems int32
}
