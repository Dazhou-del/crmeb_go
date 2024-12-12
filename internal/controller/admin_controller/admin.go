package admin_controller

import (
	"crmeb_go/internal/container/service"
	service_data "crmeb_go/internal/data/sevice_data"
	"crmeb_go/internal/validation"
	"crmeb_go/utils/ihttp"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetLoginPic(svc *service.Container) gin.HandlerFunc {
	return func(c *gin.Context) {
		var param validation.GetLoginPicParams

		if err := c.ShouldBindQuery(&param); err != nil {
			c.JSON(http.StatusOK, ihttp.Error(err))
			return
		}

		params := service_data.GetLoginPicParams{
			Temp: param.Temp,
		}

		params.SetSessionContext(c)

		res, err := svc.AdminService.GetLoginPic(params)
		if err != nil {
			c.JSON(http.StatusOK, ihttp.Error(err))
			return
		}

		c.JSON(http.StatusOK, ihttp.Data(res))
	}
}

func UploadFile(svc *service.Container) gin.HandlerFunc {
	return func(c *gin.Context) {
		file, fileHear, err := c.Request.FormFile("file")
		if err != nil {
			c.JSON(http.StatusOK, ihttp.Error(err))
		}

		key := c.Request.FormValue("key")            // 完整文件路径以及名称
		pathName := c.Request.FormValue("path_name") // 路径名
		fileType := c.Request.FormValue("type")      // 路径类型
		contentLength := c.Request.ContentLength

		params := service_data.UploadParams{
			File:          file,
			FileHeader:    fileHear,
			FileKey:       key,
			PathName:      pathName,
			FileType:      fileType,
			ContentLength: contentLength,
		}

		params.SetSessionContext(c)

		res, err := svc.AdminService.UploadFile(params)
		if err != nil {
			c.JSON(http.StatusOK, ihttp.Error(err))
			return
		}

		c.JSON(http.StatusOK, ihttp.Data(res))
	}
}
