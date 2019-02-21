package service

import (
	"errors"

	. "github.com/pobearm/workflow/entity"
)

const (
	SERVICETYPE_EMPTY int = 0 //服务的实现是空的，用于代码调试
	SERVICETYPE_PG    int = 1 //服务的实现基于postgresql数据库
)

// ServiceFactory 服务工厂
type ServiceFactory struct {
	connStr string
}

func New_ServiceFactory(conn string) *ServiceFactory {
	return &ServiceFactory{
		connStr: conn,
	}
}

// NewFlowDefineService 创建流程定义服务
func (sf ServiceFactory) GetFlowDefineService(serviceType int) (FlowDefineService, error) {
	switch serviceType {
	case SERVICETYPE_PG:
		//return New_PgFlowDefineService(sf.connStr), nil
		return nil, errors.New("未实现的FlowDefine servicetype")
	default:
		return NewEmptyFlowDefineService(sf.connStr)
	}
}

// NewOrgService 创建组织服务
func (sf ServiceFactory) GetOrgService(serviceType int) (OrgService, error) {
	switch serviceType {
	case SERVICETYPE_PG:
		//return New_PgOrgService(sf.connStr), nil
		return nil, errors.New("未实现的FlowDefine servicetype")
	default:
		return NewEmptyOrgService(sf.connStr)
	}
}

// NewCaseService 创建组织服务
func (sf ServiceFactory) GetCaseService(serviceType int) (FlowCaseService, error) {
	switch serviceType {
	case SERVICETYPE_PG:
		// return New_PgFLowCaseService(sf.connStr), nil
		return nil, errors.New("未实现的FlowDefine servicetype")
	default:
		return NewEmptyFlowCaseService(sf.connStr)
	}
}
