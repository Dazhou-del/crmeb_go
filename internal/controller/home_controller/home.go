package home_controller

import (
	"crmeb_go/internal/container/service"
	"crmeb_go/internal/data/request"
	"crmeb_go/utils/ihttp"
	"github.com/gin-gonic/gin"
	"net/http"
)

func IndexDate(svc *service.Container) gin.HandlerFunc {
	return func(c *gin.Context) {
		params := request.BaseServiceParams{}
		params.SetSessionContext(c)

		res, err := svc.HomeService.IndexDate(&params)
		if err != nil {
			c.JSON(http.StatusBadRequest, ihttp.Error(err))
			return
		}

		c.JSON(http.StatusOK, ihttp.Data(res))
	}
}

func ChartOrder(svc *service.Container) gin.HandlerFunc {
	return func(c *gin.Context) {
		params := request.BaseServiceParams{}
		params.SetSessionContext(c)

		res, err := svc.HomeService.ChartOrder(&params)
		if err != nil {
			c.JSON(http.StatusBadRequest, ihttp.Error(err))
			return
		}

		c.JSON(http.StatusOK, ihttp.Data(res))
	}
}

func ChartOrderInWeek(svc *service.Container) gin.HandlerFunc {
	return func(c *gin.Context) {
		params := request.BaseServiceParams{}
		params.SetSessionContext(c)

		res, err := svc.HomeService.ChartOrderInWeek(&params)
		if err != nil {
			c.JSON(http.StatusBadRequest, ihttp.Error(err))
			return
		}

		c.JSON(http.StatusOK, ihttp.Data(res))
	}
}

func ChartOrderInMonth(svc *service.Container) gin.HandlerFunc {
	return func(c *gin.Context) {
		params := request.BaseServiceParams{}
		params.SetSessionContext(c)

		res, err := svc.HomeService.ChartOrderInMonth(&params)
		if err != nil {
			c.JSON(http.StatusBadRequest, ihttp.Error(err))
			return
		}

		c.JSON(http.StatusOK, ihttp.Data(res))
	}
}

func ChartOrderInYear(svc *service.Container) gin.HandlerFunc {
	return func(c *gin.Context) {
		params := request.BaseServiceParams{}
		params.SetSessionContext(c)

		res, err := svc.HomeService.ChartOrderInYear(&params)
		if err != nil {
			c.JSON(http.StatusBadRequest, ihttp.Error(err))
			return
		}

		c.JSON(http.StatusOK, ihttp.Data(res))
	}
}

func ChartUser(svc *service.Container) gin.HandlerFunc {
	return func(c *gin.Context) {
		params := request.BaseServiceParams{}
		params.SetSessionContext(c)

		res, err := svc.HomeService.ChartUser(&params)
		if err != nil {
			c.JSON(http.StatusBadRequest, ihttp.Error(err))
			return
		}

		c.JSON(http.StatusOK, ihttp.Data(res))
	}
}
