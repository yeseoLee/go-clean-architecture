package rest

import (
	"log"
	"myservice/model"
	"myservice/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AccountRouter struct {
	*gin.RouterGroup
	AccountUC usecase.AccountUseCaseInterface
}

func NewAccountRouter(root *RootRouter, accountUC usecase.AccountUseCaseInterface) *AccountRouter {
	group := root.RouterGroup.Group("/account")
	AccountRouter := AccountRouter{RouterGroup: group, AccountUC: accountUC}

	group.POST("", AccountRouter.Create)
	group.GET("/:id", AccountRouter.Get)
	group.PATCH("/:id", AccountRouter.Update)
	group.DELETE("/:id", AccountRouter.Delete)

	return &AccountRouter
}

func (r *AccountRouter) Create(c *gin.Context) {
	req := model.AccountInput{}
	err := c.Bind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
	}

	res, err := r.AccountUC.Create(req)
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (r *AccountRouter) Get(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: "Missing request parameter error"})
	}

	res, err := r.AccountUC.FindById(id)
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (r *AccountRouter) Update(c *gin.Context) {
	c.Status(http.StatusNotFound)
}

func (r *AccountRouter) Delete(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: "Missing request parameter error"})
	}

	err := r.AccountUC.Delete(id)
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
