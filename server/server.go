package server

import "github.com/jahankohan/geth_mempool/utils"

func Init() {
	utils.InitConfig()

	r := NewRouter()
	r.Run(":8000")
}