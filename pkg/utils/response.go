package utils

import "github.com/gin-gonic/gin"

func Response(c *gin.Context,success string,stauscode int,data any,err any){
	var errorMsg string
	if err==nil{
		errorMsg=""
	} else{
		errorMsg=err.(error).Error()
	}
	c.JSON(stauscode,gin.H{
		"success":success,
		"status_code":stauscode,
		"data":data,
		"error":errorMsg,
	})
}