package routers

import (
	"github.com/akeramusbau/restaurant-management-system/controller"
	"github.com/gin-gonic/gin"
)

type Table struct{}

func (p *Table) Route(route *gin.Engine) {
	router := route.Group("/table")
	Controller := controller.TableController{}

	router.POST("/join", Controller.JoinTable)
}
