package handler

import (
	"fmt"

	HandlerInterface "github.com/Ansalps/BrotoStack/pkg/handler/interface"
	"github.com/Ansalps/BrotoStack/pkg/models"
	"github.com/Ansalps/BrotoStack/pkg/utils"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userhandler HandlerInterface.User_Handler_Interface
}

func NewUserHandler(service HandlerInterface.User_Handler_Interface) *UserHandler {
	return &UserHandler{
		userhandler: service,
	}
}

func (h *UserHandler) SignUp(c *gin.Context) {
	var signuprequest models.UserSignUpRequest
	if err := c.BindJSON(&signuprequest); err != nil {
		fmt.Println(err)
	}

	err := h.userhandler.ValidateSignUpRequest(signuprequest)
	if err != nil {
		fmt.Println("check in handler if error is nil or not",err)
		utils.ErrorResponse(c, 400,"failed to validate user signup request",  err)
		return
	}
	err=h.userhandler.StoreUnverifiedUserInDb(signuprequest)
	if err!=nil{
		fmt.Println("checking in handler middle")
		utils.ErrorResponse(c,500,"failed to store unverified user in database",err)
		return
	}
	fmt.Println("checking in handler last")
	utils.SuccessResponse(c,200,"unverified user successfully inserted into database. front end developer, please let me know if you need any detalis of the unverified user",nil)
}
