package user_service

import (
	"crmeb_go/internal/data/eb_user_data"
	service_data "crmeb_go/internal/data/sevice_data"
	"crmeb_go/internal/server"
	"crmeb_go/utils/izap"
	"go.uber.org/zap"
)

type GetRealNameImpl interface {
	GetRealName(params service_data.GetRealNameParams) (data eb_user_data.GetRealNameData, err error)
}

type GetRealNameService struct {
	svc *server.SvcContext
}

func NewGetRealNameService(svc *server.SvcContext) GetRealNameService {
	return GetRealNameService{svc: svc}
}

func (s GetRealNameService) GetRealName(params service_data.GetRealNameParams) (data eb_user_data.GetRealNameData, err error) {
	userInfoModel, err := s.svc.Repo.UserRepository.GetRealName(params.Ctx, params.UserId)
	if err != nil {
		izap.Log.Error("EbUserRepository.QueryOne [err]:%v", zap.Error(err))

		return
	}

	data.RealName = userInfoModel.RealName
	data.UserId = uint64(userInfoModel.UID)

	return
}
