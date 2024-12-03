package user_controller

import (
	"crmeb_go/internal/container/service"
	service_data "crmeb_go/internal/data/sevice_data"
	"crmeb_go/internal/validation"
	"crmeb_go/utils/xhttp"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetRealName(svc *service.Container) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 参数校验
		var param validation.GetRealNameParams
		if err := c.ShouldBindJSON(&param); err != nil {
			c.JSON(http.StatusOK, xhttp.Error(err))
			return
		}

		params := service_data.GetRealNameParams{
			UserId: param.UserId,
		}

		params.SetSessionContext(c)
		res, err := svc.UserService.GetRealName(params)
		if err != nil {
			c.JSON(http.StatusOK, xhttp.Error(err))
			return
		}

		c.JSON(http.StatusOK, xhttp.Data(res))
	}
}