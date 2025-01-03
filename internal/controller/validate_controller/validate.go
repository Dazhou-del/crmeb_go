package validate_controller

import (
	"crmeb_go/internal/container/service"
	"crmeb_go/utils/ihttp"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetValidateCode(svc *service.Container) gin.HandlerFunc {
	return func(c *gin.Context) {
		res, err := svc.AdminService.GetValidateCode()
		if err != nil {
			c.JSON(http.StatusOK, ihttp.Error(err))
			return
		}

		c.JSON(http.StatusOK, ihttp.Data(res))
	}
}
