package rest

import (
	"myservice/rest/middleware"
	"myservice/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type StockRouter struct {
	*gin.RouterGroup
	StockUC usecase.StockUseCaseInterface
}

func NewStockRouter(root *RootRouter, stockUC usecase.StockUseCaseInterface) *StockRouter {
	group := root.RouterGroup.Group("/stock")
	StockRouter := StockRouter{RouterGroup: group, StockUC: stockUC}

	group.Use(middleware.Authentication())
	group.GET("/:id", StockRouter.Get)

	return &StockRouter
}

func (r *StockRouter) Get(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: "Missing request parameter error"})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}
