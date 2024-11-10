package router

import (
	"github.com/FakJeongTeeNhoi/co-working-space-management/controller"
	"github.com/gin-gonic/gin"
)

func SpaceRouterGroup(server *gin.RouterGroup) {
	space := server.Group("/space")
	space.GET("/search", controller.SearchSpaceWithParam)
	space.GET("/",controller.DisplayAllSpace)
	space.POST("/", controller.CreateSpace)
	space.GET("/:id", controller.DisplaySpaceInfo)
	space.PUT("/:id", controller.UpdateSpace)
	space.POST("/:id/addStaff", controller.AddStaffToSpace)
	space.POST("/:id/removeStaff", controller.RemoveStaffFromSpace)
	space.POST("/:id/addRoom", controller.AddRoomToSpace)
	space.POST("/:id/removeRoom", controller.RemoveRoomFromSpace)
	space.GET("room/:id", controller.DisplaySpaceWithRoomInfo)
}
