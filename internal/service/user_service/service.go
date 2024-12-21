package user_service

import (
	"crmeb_go/internal/data/request"
	"crmeb_go/internal/data/response"
	"crmeb_go/internal/server"
	"crmeb_go/utils/itime"
	"crmeb_go/utils/izap"
	"go.uber.org/zap"
)

type UserServiceImpl interface {
	GetRealName(params *request.GetRealNameParams) (data *response.GetRealNameResp, err error)
	GetRegisterNumByDate(params *request.DateParams) (*response.DateResp, error)
}

// UserService 用户表 模型服务
type UserService struct {
	svc *server.SvcContext
}

// NewUserService 新用户表 模型服务实例
func NewUserService(svc *server.SvcContext) *UserService {
	return &UserService{svc: svc}
}

func (u *UserService) GetRealName(params *request.GetRealNameParams) (data *response.GetRealNameResp, err error) {
	service := NewGetRealNameService(u.svc)
	data, err = service.GetRealName(params)

	return
}
func (u *UserService) GetRegisterNumByDate(params *request.DateParams) (*response.DateResp, error) {
	var data response.DateResp
	// 今日用户注册量
	nowPageViews, err := u.svc.Repo.UserRepository.GetRegisterNumByDate(params.BaseServiceParams.Ctx, params.Start, params.End)
	if err != nil {
		izap.Log.Error("查询今日访问量失败", zap.String("开始时间", itime.Format(params.Start)), zap.String("结束时间", itime.Format(params.End)), zap.Error(err))

		return &data, err
	}

	// 昨天用户注册量
	yesterdayPageViews, err := u.svc.Repo.UserRepository.GetRegisterNumByDate(params.BaseServiceParams.Ctx, params.YesterdayStart, params.YesterdayEnd)
	if err != nil {
		izap.Log.Error("查询今日访问量失败", zap.String("开始时间", itime.Format(params.YesterdayStart)), zap.String("结束时间", itime.Format(params.YesterdayEnd)), zap.Error(err))

		return &data, err
	}

	data.NowData = nowPageViews
	data.YesterdayData = yesterdayPageViews
	return &data, nil
}
