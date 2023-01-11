package server

import (
	"context"
	"landscape/constant"
	"landscape/dal"
	"landscape/github.com/landscape/api"
)

func (svcImpl *apiServiceImpl) QueryCityList(ctx context.Context, req *api.QueryCityListReq) (*api.QueryCityListResp, error) {
	provinces := make([]*api.Province, 0)
	provinceMap := constant.GetProvinceMap()
	for key, value := range provinceMap {
		provinces = append(provinces, &api.Province{
			Name:   key,
			Cities: value,
		})
	}
	resp := api.QueryCityListResp{
		Provinces: provinces,
		Success:   true,
		ErrorCode: "200",
	}
	return &resp, nil
}

func (svcImpl *apiServiceImpl) QueryEssayList(ctx context.Context, req *api.QueryEssayListReq) (*api.QueryEssayListResp, error) {
	return dal.QueryEssayList(ctx, req)
}

func (svcImpl *apiServiceImpl) QueryOneEssay(ctx context.Context, req *api.QueryOneEssayReq) (*api.QueryOneEssayResp, error) {
	return dal.QueryOneEssay(ctx, req)
}

func (svcImpl *apiServiceImpl) QuerySwiperImageList(ctx context.Context, req *api.QuerySwiperImageListReq) (*api.QuerySwiperImageListResp, error) {
	return dal.QuerySwiperImageList()
}
