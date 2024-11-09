package controller

import (
	"github.com/FakJeongTeeNhoi/co-working-space-management/model"
	"github.com/FakJeongTeeNhoi/co-working-space-management/model/response"
	"github.com/FakJeongTeeNhoi/co-working-space-management/service"
	"github.com/gin-gonic/gin"
)

func CreateRoom(c *gin.Context) {
	roomRequest := model.RoomCreateRequest{}
	if err := c.ShouldBindJSON(&roomRequest); err != nil {
		response.BadRequest("Invalid request").AbortWithError(c)
		return
	}

	room, err := roomRequest.CreateRoom()
	if err != nil {
		response.InternalServerError("Failed to create room").AbortWithError(c)
		return
	}

	c.JSON(200, response.CommonResponse{
		Success: true,
	}.AddInterfaces(map[string]interface{}{
		"room": room,
	}))
}

func UpdateRoom(c *gin.Context) {
	id := c.Param("id")
	room := model.Room{}
	if err := c.ShouldBindJSON(&room); err != nil {
		response.BadRequest("Invalid request").AbortWithError(c)
		return
	}

	if err := room.UpdateRoom(map[string]interface{}{"id": service.ParseToInt(id)}); err != nil {
		response.InternalServerError("Failed to update room").AbortWithError(c)
		return
	}

	c.JSON(200, response.CommonResponse{
		Success: true,
	}.AddInterfaces(map[string]interface{}{
		"room": room,
	}))
}

func DeleteRoom(c *gin.Context) {
	id := c.Param("id")
	room := model.Room{}
	if err := room.DeleteRoom(map[string]interface{}{"id": service.ParseToInt(id)}); err != nil {
		response.InternalServerError("Failed to delete room").AbortWithError(c)
		return
	}

	c.JSON(200, response.CommonResponse{
		Success: true,
	})
}

func DisplayRoomInfo(c *gin.Context) {
	id := c.Param("id")
	room := model.Room{}

	if err := room.GetOneRoom(map[string]interface{}{"id": service.ParseToInt(id)}); err != nil {
		response.NotFound("Room not found").AbortWithError(c)
		return
	}

	roomResponse := model.RoomResponse{}
	model.ToRoomResponse(&room, &roomResponse)

	c.JSON(200, response.CommonResponse{
		Success: true,
	}.AddInterfaces(map[string]interface{}{
		"room": roomResponse,
	}))
}