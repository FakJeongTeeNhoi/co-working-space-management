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

	c.JSON(200, response.CommonResponse{
		Success: true,
	}.AddInterfaces(map[string]interface{}{
		"space": spaceResponse,
	}))
}
