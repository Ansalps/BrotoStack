package utils

import "github.com/gin-gonic/gin"

func SuccessResponse(c *gin.Context,stauscode int,message string,data any){
	c.JSON(stauscode,gin.H{
		"success":true,
		"message":message,
		"data":data,
	})
}
func ErrorResponse(c *gin.Context,statuscode int,message string,err any){
	var ErrorMessage string
	v,ok:=err.(error)
	if ok{
		ErrorMessage=v.Error()
	} else {
		ErrorMessage="not able to do type assertion aon erro"
	}
	c.JSON(statuscode,gin.H{
		"success":false,
		"meassage":message,
		"data":ErrorMessage,
	})
}