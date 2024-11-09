package router

import (
	"github.com/FakJeongTeeNhoi/co-working-space-management/controller"
	"github.com/gin-gonic/gin"
)

func RoomRouterGroup(server *gin.RouterGroup) {
	space := server.Group("/room")
	space.GET("/:id",controller.DisplayRoomInfo)
	space.POST("/", controller.CreateRoom)
	space.PUT("/:id", controller.UpdateRoom)
	space.DELETE("/:id", controller.DeleteRoom)
}