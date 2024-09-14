package service

import (
	"github.com/FakJeongTeeNhoi/co-working-space-management/model/response"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetUintID(s string, c *gin.Context) uint {
	i, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		response.BadRequest("Invalid ID").AbortWithError(c)
		return -1
	}
	return uint(i)
}
