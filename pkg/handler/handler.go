package handler

import (
	"github.com/Ansalps/BrotoStack/pkg/service"
	"github.com/gin-gonic/gin"
)
 type UserHandler struct{
	handler *service.UserService
 }
 func NewUserHandler(service *service.UserService)*UserHandler{
	return &UserHandler{
		handler: service,
	}
 }

 func (h *UserHandler) SignUp(c *gin.Context){

 }