package service

import (
	"log"
	"strconv"
	"time"
)

// func GetUintID(s string, c *gin.Context) uint {
// 	i, err := strconv.ParseUint(s, 10, 64)
// 	if err != nil {
// 		response.BadRequest("Invalid ID").AbortWithError(c)
// 		return -1
// 	}
// 	return uint(i)
// }

func ParseToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Println("Failed to parse string to int: ", err)
	}
	return i
}

func ParseToInt64(s string) int64 {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		log.Println("Failed to parse string to int64: ", err)
	}
	return i
}


func ParseToTime(s string) time.Time {
	t, err := time.Parse(time.RFC3339, s)
	if err != nil {
		log.Println("Failed to parse string to time: ", err)
	}
	return t
}

func ParseToFloat64Array(s string) [2]float64 {
	var arr [2]float64
	if s == "" {
		return arr
	}
	for i, v := range s {
		if v == ',' {
			arr[0] = ParseToFloat64(s[:i])
			arr[1] = ParseToFloat64(s[i+1:])
			break
		}
	}
	return arr
}

func ParseToFloat64(s string) float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		log.Println("Failed to parse string to float64: ", err)
	}
	return f
}

func ParsToString(i int) string {
	return strconv.Itoa(i)
}