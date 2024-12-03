package admin_service

import (
	"context"
	"crmeb_go/internal/data/admin_data"
	"crmeb_go/internal/server"
	"crmeb_go/utils/base64_captcha"
	"go.uber.org/zap"
)

type GetValidateCodeGetRealNameImpl interface {
	GetValidateCode(context context.Context) (data admin_data.ValidateCodeData, err error)
}

type GetValidateCodeService struct {
	svc *server.SvcContext
}

func NewGetValidateCodeService(svc *server.SvcContext) *GetValidateCodeService {
	return &GetValidateCodeService{svc: svc}
}

func (a GetValidateCodeService) GetValidateCode() (admin_data.ValidateCodeData, error) {
	var data admin_data.ValidateCodeData

	err, code, id := base64_captcha.GetCaptcha(a.svc.RedisClient, a.svc.Logger, a.svc.Ctx)
	if err != nil {
		a.svc.Logger.Error("生成验证码错误err:", zap.Error(err))

		return data, err
	}

	data.Code = code
	data.Key = id

	return data, err
}