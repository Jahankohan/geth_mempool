package controller

import (
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"github.com/jahankohan/geth_mempool/modules"
	"github.com/jahankohan/geth_mempool/utils"
)

type BlocksController struct {}

func (bc BlocksController) GetLatestBlock(c *gin.Context) {
	variables := utils.GetConfig()
	client, err := ethclient.Dial(variables.LocalNetwork)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error" : "Problem with connecting to Geth",
			"description" : err.Error(),
		})
	}
	handler := modules.ClientHandler{Client: client}
	latestBlock := handler.GetLatestBlock()
	c.JSON(http.StatusOK, latestBlock)
}

func (bc BlocksController) GetTransactionID(c *gin.Context) {
	hashStr := c.Query("hash")
	variables := utils.GetConfig()
	client, err := ethclient.Dial(variables.LocalNetwork)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error" : "Problem with connecting to Geth",
			"description" : err.Error(),
		})
	}
	handler := modules.ClientHandler{Client: client}
	tx := handler.GetTxByHash(common.HexToHash(hashStr))
	c.JSON(http.StatusOK, tx)
}