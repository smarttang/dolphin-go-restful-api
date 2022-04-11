package router

import (
	"github.com/gin-gonic/gin"
	apis "dolphin/api/apis"
)

func InitRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.GET("/", apis.DefaultIndex)
	router.POST("/add", apis.JobAdd)
	router.POST("/status", apis.JobStatus)
	router.POST("/report", apis.JobReport)
	return router
}