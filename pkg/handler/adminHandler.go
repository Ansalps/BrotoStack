package handler

import (
	"fmt"

	HandlerInterface "github.com/Ansalps/BrotoStack/pkg/handler/interface"
	"github.com/Ansalps/BrotoStack/pkg/models"
	"github.com/Ansalps/BrotoStack/pkg/utils"
	"github.com/gin-gonic/gin"
)

type AdminHandler struct {
	adminHandler HandlerInterface.Admin_Handler_Interface
}

func NewAdminHanlder(service HandlerInterface.Admin_Handler_Interface) *AdminHandler {
	return &AdminHandler{
		adminHandler: service,
	}
}
func (h *AdminHandler) AdmnSignUp(c *gin.Context) {
	var signuprequest models.UserSignUpRequest
	if err := c.BindJSON(&signuprequest); err != nil {
		fmt.Println(err)
	}

	err := h.adminHandler.ValidateAdminSignUpRequest(signuprequest)
	if err != nil {
		fmt.Println("check in handler if error is nil or not", err)
		utils.ErrorResponse(c, 400, "failed to validate admin signup request", err)
		return
	}
	
	fmt.Println("checking in handler last")
	utils.SuccessResponse(c, 200, "unverified admin successfully inserted into database. front end developer, please let me know if you need any detalis of the unverified user", nil)
}
