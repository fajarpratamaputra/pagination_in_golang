package config

import "github.com/gin-gonic/gin"

func Response(c *gin.Context, Message interface{}, Data interface{}) {
	c.JSON(200, gin.H{
		"responseCode":        1000,
		"responseStatus":      "SUCCESS",
		"responseDescription": Message,
		"data":                Data,
	})
}
func ResponseTest(c *gin.Context, Message interface{}, Data interface{}) {
	c.JSON(200, gin.H{
		"responseCode":        1000,
		"responseStatus":      "SUCCESS",
		"responseDescription": Message,
		"data":                Data,
	})
}
func ErrorResponse(c *gin.Context, Message interface{}) {
	c.AbortWithStatusJSON(200, gin.H{
		"responseCode":        ResponseStruct.Get().ResponseCode,
		"responseStatus":      "FAILED",
		"responseDescription": Message,
	})
}
func ResponseSingle(c *gin.Context, Data interface{}) {
	c.JSON(200, Data)
}

var ResponseStruct = responseStruct{9999}

type responseStruct struct {
	ResponseCode int
}

func (r *responseStruct) Get() *responseStruct {
	return r
}

func (r *responseStruct) Set(code int) *responseStruct {
	r.ResponseCode = code
	return r
}
