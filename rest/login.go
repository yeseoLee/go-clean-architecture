package rest

import (
	"log"
	"myservice/model"
	"myservice/usecase"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type LoginRouter struct {
	*gin.RouterGroup
	AccountUC usecase.AccountUseCaseInterface
}

func NewLoginRouter(root *RootRouter, accountUC usecase.AccountUseCaseInterface) *LoginRouter {
	group := root.RouterGroup.Group("")
	LoginRouter := LoginRouter{RouterGroup: group, AccountUC: accountUC}

	group.POST("/login", LoginRouter.Login)
	group.POST("/logout", LoginRouter.Logout)

	return &LoginRouter
}

func (r *LoginRouter) Login(c *gin.Context) {
	req := model.LoginInput{}
	err := c.Bind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
	}

	err = r.AccountUC.Login(req)
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
		return
	}

	session := sessions.Default(c)
	session.Set("id", req.ID)
	session.Save()

	c.JSON(http.StatusOK, gin.H{
		"message": "User Sign In Successfully",
	})
}

func (r *LoginRouter) Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.JSON(http.StatusOK, gin.H{
		"message": "User Sign Out Successfully",
	})
}
