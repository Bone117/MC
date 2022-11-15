package router

import (
	"server/api"

	"github.com/gin-gonic/gin"
)

type PeriodRouter struct{}

func (p *PeriodRouter) InitPeriodRouter(Router *gin.RouterGroup) {
	periodRouter := Router.Group("period")
	periodApi := api.ApiGroupApp.PeriodApi
	{
		periodRouter.POST("createPeriod", periodApi.CreatePeriod)
		periodRouter.POST("deletePeriod", periodApi.DeletePeriod)
		periodRouter.POST("updatePeriod", periodApi.UpdatePeriod)
		periodRouter.POST("getPeriod", periodApi.GetPeriod)
		periodRouter.POST("getPeriodList", periodApi.GetPeriodList)
		periodRouter.POST("createCpTime", periodApi.CreateCpTime)
		periodRouter.POST("updateCpTime", periodApi.UpdateCpTime)
		periodRouter.POST("deleteCpTime", periodApi.DeleteCpTime)
		periodRouter.POST("getCpTimeList", periodApi.GetCpTimeList)

	}
}
