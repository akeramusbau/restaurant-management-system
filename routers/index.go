package routers

import "github.com/gin-gonic/gin"

type RouterInterface interface {
	Route(*gin.Engine)
}

type RouteLoader struct{}

func (loader RouteLoader) LoadRoutes() []RouterInterface {
	table := new(table)
	return []RouterInterface{
		table,
	}
}
