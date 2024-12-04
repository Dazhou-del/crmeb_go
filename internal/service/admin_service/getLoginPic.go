package admin_service

import (
	"crmeb_go/internal/data/admin_data"
	service_data "crmeb_go/internal/data/sevice_data"
	"crmeb_go/internal/server"
	"crmeb_go/internal/service/eb_system_config_service"
	"go.uber.org/zap"
)

type GetLoginPicImpl interface {
	GetLoginPic(params service_data.GetLoginPicParams) (data admin_data.GetLoginPicResp, err error)
}

type GetLoginPicService struct {
	svc *server.SvcContext
}

func NewGetLoginPicService(svc *server.SvcContext) *GetLoginPicService {
	return &GetLoginPicService{svc: svc}
}

func (a GetLoginPicService) GetLoginPic(params service_data.GetLoginPicParams) (data admin_data.GetLoginPicResp, err error) {
	result := make(map[string]any)

	//背景图
	var systemConfigParam service_data.GetSystemConfigParams
	systemConfigParam.BaseServiceParams = params.BaseServiceParams
	systemConfigParam.Name = service_data.CONFIG_KEY_ADMIN_LOGIN_BACKGROUND_IMAGE
	backgroundImage, err := eb_system_config_service.NewGetSystemConfigInfoService(a.svc).GetSystemConfigInfo(systemConfigParam)

	if err != nil {
		a.svc.Logger.Error("查询系统配置错误:", zap.String("name", service_data.CONFIG_KEY_ADMIN_LOGIN_BACKGROUND_IMAGE), zap.Error(err))

		return data, err
	}

	result["backgroundImage"] = backgroundImage.Value

	// logo
	systemConfigParam.Name = service_data.CONFIG_KEY_ADMIN_LOGIN_LOGO_LEFT_TOP
	logo, err := eb_system_config_service.NewGetSystemConfigInfoService(a.svc).GetSystemConfigInfo(systemConfigParam)

	if err != nil {
		a.svc.Logger.Error("查询系统配置错误:", zap.String("name", service_data.CONFIG_KEY_ADMIN_LOGIN_LOGO_LEFT_TOP), zap.Error(err))

		return data, err
	}

	result["logo"] = logo.Value

	// loginLogo
	systemConfigParam.Name = service_data.CONFIG_KEY_ADMIN_LOGIN_LOGO_LOGIN
	loginLogo, err := eb_system_config_service.NewGetSystemConfigInfoService(a.svc).GetSystemConfigInfo(systemConfigParam)

	if err != nil {
		a.svc.Logger.Error("查询系统配置错误:", zap.String("name", service_data.CONFIG_KEY_ADMIN_LOGIN_LOGO_LOGIN), zap.Error(err))

		return data, err
	}

	result["loginLogo"] = loginLogo.Value

	// todo 轮播图
	data.Map = result
	return data, nil
}
