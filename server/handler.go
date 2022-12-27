package server

import (
	"context"
	"reservation/dao"
	"reservation/github.com/reservation/api"
)

func (svcImpl *apiServiceImpl) GetSystemInfo(ctx context.Context, req *api.GetSystemInfoReq) (*api.GetSystemInfoResp, error){
	return dao.GetSystemInfo()
}

func (svcImpl *apiServiceImpl) QueryActivityList(ctx context.Context, req *api.QueryActivityListReq) (*api.QueryActivityListResp, error){
	return dao.QueryActivityList()
}

func (svcImpl *apiServiceImpl) QueryActivityListByDay(ctx context.Context, req *api.QueryActivityListByDayReq) (*api.QueryActivityListByDayResp, error){
	return dao.QueryActivityListByDay(req)
}


