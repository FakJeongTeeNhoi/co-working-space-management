package controller

import (
	"github.com/FakJeongTeeNhoi/co-working-space-management/model"
	"github.com/FakJeongTeeNhoi/co-working-space-management/model/response"
	"github.com/FakJeongTeeNhoi/co-working-space-management/service"
	"github.com/gin-gonic/gin"
)

func SearchSpaceWithParam(c *gin.Context) {
	SearchParam := model.SpaceSearchParam{}
	spaces := model.SpaceResponses{}

	SearchParam.Name = c.Query("name")
	SearchParam.Faculty = c.Query("faculty")
	SearchParam.Start_datetime = service.ParseToTime(c.Query("start_datetime"))
	SearchParam.End_datetime = service.ParseToTime(c.Query("end_datetime"))
	SearchParam.Capacity = service.ParseToInt(c.Query("capacity"))
	SearchParam.Latitude_range = service.ParseToFloat64Array(c.Query("latitude_range"))
	SearchParam.Longitude_range = service.ParseToFloat64Array(c.Query("longitude_range"))

	if err := spaces.GetAllWithSearchParam(SearchParam); err != nil {
		response.NotFound("Space not found").AbortWithError(c)
		return
	}

	c.JSON(200, response.CommonResponse{
		Success: true,
	}.AddInterfaces(map[string]interface{}{
		"count": len(spaces),
		"space": spaces,
	}))
}

func DisplaySpaceInfo(c *gin.Context) {
	id := c.Param("id")
	space := model.Space{}

	if err := space.GetOne(map[string]interface{}{"id": service.ParseToInt(id)}); err != nil {
		response.NotFound("Space not found").AbortWithError(c)
		return
	}

	spaceResponse := model.SpaceResponse{}
	model.ToSpaceResponse(&space, &spaceResponse)
	model.EmbeddedRoomList(&space, &spaceResponse, 0)

	c.JSON(200, response.CommonResponse{
		Success: true,
	}.AddInterfaces(map[string]interface{}{
		"space": spaceResponse,
	}))
}

func CreateSpace(c *gin.Context) {
	spaceRequest := model.CreateSpaceRequest{}
	if err := c.ShouldBindJSON(&spaceRequest); err != nil {
		response.BadRequest("Invalid request").AbortWithError(c)
		return
	}
	space, err := spaceRequest.CreateSpace();
	if err != nil {
		response.InternalServerError("Failed to create space").AbortWithError(c)
		return
	}
	c.JSON(200, response.CommonResponse{
		Success: true,
	}.AddInterfaces(map[string]interface{}{
		"space": space,
	}))
}

func UpdateSpace(c *gin.Context) {
	id := c.Param("id")
	space := model.Space{}
	if err := c.ShouldBindJSON(&space); err != nil {
		response.BadRequest("Invalid request").AbortWithError(c)
		return
	}

	if err := space.UpdateSpace(map[string]interface{}{"id": service.ParseToInt(id)}); err != nil {
		response.InternalServerError("Failed to update space").AbortWithError(c)
		return
	}

	c.JSON(200, response.CommonResponse{
		Success: true,
	}.AddInterfaces(map[string]interface{}{
		"space": space,
	}))
}

func AddStaffToSpace(c *gin.Context) {
	id := c.Param("id")
	addStaffRequest := model.AddStaffToSpaceRequest{}
	if err := c.ShouldBindJSON(&addStaffRequest); err != nil {
		response.BadRequest("Invalid request").AbortWithError(c)
		return
	}

	if err := addStaffRequest.AddStaffToSpace(service.ParseToInt(id)); err != nil {
		response.InternalServerError("Failed to add staff to space").AbortWithError(c)
		return
	}

	c.JSON(200, response.CommonResponse{
		Success: true,
	})
}

func RemoveStaffFromSpace(c *gin.Context) {
	id := c.Param("id")
	removeStaffRequest := model.RemoveStaffFromSpaceRequest{}
	if err := c.ShouldBindJSON(&removeStaffRequest); err != nil {
		response.BadRequest("Invalid request").AbortWithError(c)
		return
	}

	if err := removeStaffRequest.RemoveStaffFromSpace(service.ParseToInt(id)); err != nil {
		response.InternalServerError("Failed to remove staff from space").AbortWithError(c)
		return
	}

	c.JSON(200, response.CommonResponse{
		Success: true,
	})
}

func AddRoomToSpace(c *gin.Context) {
	id := c.Param("id")
	addRoomRequest := model.AddRoomToSpaceRequest{}
	if err := c.ShouldBindJSON(&addRoomRequest); err != nil {
		response.BadRequest("Invalid request").AbortWithError(c)
		return
	}

	if err := addRoomRequest.AddRoomToSpace(service.ParseToInt(id)); err != nil {
		response.InternalServerError("Failed to add room to space").AbortWithError(c)
		return
	}

	c.JSON(200, response.CommonResponse{
		Success: true,
	})
}

func RemoveRoomFromSpace(c *gin.Context) {
	id := c.Param("id")
	removeRoomRequest := model.RemoveRoomFromSpaceRequest{}
	if err := c.ShouldBindJSON(&removeRoomRequest); err != nil {
		response.BadRequest("Invalid request").AbortWithError(c)
		return
	}

	if err := removeRoomRequest.RemoveRoomFromSpace(service.ParseToInt(id)); err != nil {
		response.InternalServerError("Failed to remove room from space").AbortWithError(c)
		return
	}

	c.JSON(200, response.CommonResponse{
		Success: true,
	})
}

func DisplaySpaceWithRoomInfo(c *gin.Context) {
	id := c.Param("id")
	space := model.Space{}
	if err := space.GetOneWithRoomInfo(id); err != nil {
		response.NotFound("Space not found").AbortWithError(c)
		return
	}

	spaceResponse := model.SpaceResponse{}
	model.ToSpaceResponse(&space, &spaceResponse)
	model.EmbeddedRoomListWithSpace(&space, &spaceResponse,uint(service.ParseToInt(id)))

	c.JSON(200, response.CommonResponse{
		Success: true,
	}.AddInterfaces(map[string]interface{}{
		"space": spaceResponse,
	}))

}