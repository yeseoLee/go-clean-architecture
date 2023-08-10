package rest

import (
	"myservice/usecase"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type RootRouter struct{ *gin.RouterGroup }

func NewGin(accountUC usecase.AccountUseCaseInterface,
	stockUC usecase.StockUseCaseInterface,
	store sessions.Store) *gin.Engine {
	g := gin.Default()

	// middlewares
	g.Use(sessions.Sessions("mysession", store))

	// routes
	g.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/api")
	})
	g.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})
	root := NewRootRouter(g)

	// TODO: 의존성 주입
	_ = NewAccountRouter(root, accountUC)
	_ = NewLoginRouter(root, accountUC)
	_ = NewStockRouter(root, stockUC)

	return g
}

func NewRootRouter(g *gin.Engine) *RootRouter {
	group := g.Group("/api")
	rootRouter := RootRouter{RouterGroup: group}

	rootRouter.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "my rest api")
	})
	return &rootRouter
}

// TODO: controller, service, route ???
