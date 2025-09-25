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
		utils.Response(c, "false", 400, nil, err)
		return
	}
	err=h.userhandler.StoreUnverifiedUserInDb(signuprequest)
	if err!=nil{
		fmt.Println("checking in handler middle")
		utils.Response(c,"false",500,nil,err)
		return
	}
	fmt.Println("checking in handler last")
	utils.Response(c,"true",200,"unverified user successfully inserted into database. front end developer, please let me know if you need any detalis of the unverified user",err)
}
