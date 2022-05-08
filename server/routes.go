package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jahankohan/geth_mempool/controller"
	"github.com/jahankohan/geth_mempool/utils"
)

func NewRouter() *gin.Engine {
	variables := utils.GetConfig()

	if variables.DebugMode {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()

	// MiddleWares
	router.Use(gin.Recovery())
	router.Use(gin.Logger())
	router.Use(cors.Default())

	// In Case we need to serve static data
	// router.Static("/static", "./static")

	v1 := router.Group("v1") 
	{
		blocks := v1.Group("blocks") 
		{
			bc := controller.BlocksController{}

			blocks.GET("/latest-block", bc.GetLatestBlock)
			blocks.GET("/get-tx", bc.GetTransactionID)
		}
	}
	return router
}