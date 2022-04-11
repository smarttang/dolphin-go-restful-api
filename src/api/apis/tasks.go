package apis

import (
	"github.com/gin-gonic/gin"
	"net/http"
	model "dolphin/api/models"
	"strconv"
)

func DefaultIndex(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"data": "dolphin PaaS Service v1.0",
	})
}

// 添加任务
func JobAdd(c *gin.Context) {
	var task model.Task

	task.Task_Type, _ = strconv.ParseInt(c.Param("task_type"), 10, 64)
	task.Task_Info = c.Request.FormValue("task_info")

	result, err := task.AddTask()

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg": "create task filed!!",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":1,
		"data": result,
	})
}

// 获取扫描状态
func JobStatus(c *gin.Context) {
	var task model.Task
	task.ID = c.Request.FormValue("id")
	task = task.CheckStatus()
	c.JSON(http.StatusOK, gin.H{
		"code":1,
		"data":task,
	})
}

// 粗暴获取结果
func JobReport(c *gin.Context) {
	var report model.Report
	report.ID = c.Request.FormValue("id")
	report = report.GetReports()
	c.JSON(http.StatusOK, gin.H{
		"code":1,
		"data":report,
	})
}