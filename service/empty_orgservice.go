package service

import (
	. "github.com/pobearm/workflow/entity"
)

type EmptyOrgService struct {
}

var _ OrgService = new(EmptyOrgService)

func NewEmptyOrgService(connstr string) (*EmptyOrgService, error) {
	return &EmptyOrgService{}, nil
}

//根据角色,  部门id找到用户
func (s *EmptyOrgService) FindUser(role, departid string) (us []*FlowUser, err error) {
	return make([]*FlowUser, 0), nil
}

//找到用户直属部门
func (s *EmptyOrgService) FindUserDept(userid string) (deptid string, err error) {
	return "", nil
}

//找到用户直属部门的父部门
func (s *EmptyOrgService) FindUserParentDept(userid string) (deptid string, err error) {
	return "", nil
}

//根据用户id,获取用户
func (s *EmptyOrgService) GetUser(userid string) (us []*FlowUser, err error) {
	return make([]*FlowUser, 0), nil
}

//根据部门找到所有用户
func (s *EmptyOrgService) FindUserByDept(departid string) (us []*FlowUser, err error) {
	return make([]*FlowUser, 0), nil
}

//根据角色找到所有用户
func (s *EmptyOrgService) FindUserByRole(role string) (us []*FlowUser, err error) {
	return make([]*FlowUser, 0), nil
}
