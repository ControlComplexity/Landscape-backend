package dal

import (
	"context"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"landscape/constant"
	"landscape/github.com/landscape/api"
	"landscape/model"
)

func Init() *gorm.DB {
	fmt.Println("init gorm db...")
	db, err := gorm.Open("mysql", "root:root@(127.0.0.1:3306)/landscape?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	return db
}

func QueryEssayList(ctx context.Context, req *api.QueryEssayListReq) (*api.QueryEssayListResp, error) {
	db := Init()
	do := []*api.Essay{}
	err := db.Model(&api.Essay{}).Find(&do).Error
	if err == nil {
		return &api.QueryEssayListResp{
			EssayList: do,
			Success:   true,
			ErrorCode: "200",
		}, nil
	} else {
		return &api.QueryEssayListResp{
			ErrorMsg: "获取文章详情失败，错误是" + err.Error(),
		}, err
	}
}

func QueryOneEssay(ctx context.Context, req *api.QueryOneEssayReq) (*api.QueryOneEssayResp, error) {
	db := Init()
	var essay model.EssayDO
	err := db.Model(&model.EssayDO{}).Find(&essay).Where("uuid = ", req.Uuid).Limit(1).Error
	fmt.Println("req.Uuid: ", req.Uuid, " essay: ", essay)
	if err == nil {
		return &api.QueryOneEssayResp{
			Essay: &api.Essay{
				Uuid:  essay.Uuid,
				Title: essay.Title,
			},
			Success:   true,
			ErrorCode: "200",
		}, nil
	} else {
		return &api.QueryOneEssayResp{
			ErrorMsg: "获取文章详情失败，错误是" + err.Error(),
		}, err
	}
}

func QuerySwiperImageList() (*api.QuerySwiperImageListResp, error) {
	db := Init()
	var imgs []model.Image
	err := db.Model(&model.Image{}).Find(&imgs).Error
	fmt.Println("imgs: ", imgs)
	var images []string

	for _, img := range imgs {
		images = append(images, img.URL)
	}
	if err == nil {
		return &api.QuerySwiperImageListResp{
			Image:     images,
			Success:   true,
			ErrorCode: constant.SUCCESS_ERROR_CODE,
		}, nil
	} else {
		return &api.QuerySwiperImageListResp{
			ErrorMsg: "获取轮播图图片失败，错误是" + err.Error(),
		}, err
	}
}
