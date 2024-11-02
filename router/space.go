package router

import (
	"github.com/FakJeongTeeNhoi/co-working-space-management/controller"
	"github.com/gin-gonic/gin"
)

func SpaceRouterGroup(server *gin.RouterGroup) {
	space := server.Group("/space")
	space.GET("/search", controller.SearchSpaceWithParam)
	space.GET("/:id", controller.DisplaySpaceInfo)
}
