package main

import (
	"github.com/gin-gonic/gin"
)

func BlockModelRouters(r *gin.Engine) *gin.Engine {

	r.GET("/blockModel")

	return r
}

func CollectRoute(r *gin.Engine) *gin.Engine {
	r = BlockRouters(r)

	return r
}
