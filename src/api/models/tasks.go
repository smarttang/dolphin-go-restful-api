package models

import (
	"github.com/google/uuid"
	orm "dolphin/api/database"
	"fmt"
)

type Task struct {
	ID			string `json:"id"`						// ID
	Task_Type 	int64 	`json:"task_type"`				// 任务类型(0:代码漏洞扫描,1:端口扫描,2:sonar数据抓取,3:调用链路分析)
	Task_Info	string `json:"task_info"`				// 扫描任务的基础信息,JSON的方式提交，后端自己拆分解析
	Task_Status	int64 	`json:"task_status"`			// 扫描任务状态(0:待处理,1:处理中,2:处理完成,3:处理失败)
}

var Tasks []Task

//设置表名
func (Task) TableName() string {
    return "task"
}

// 新增任务
func (task Task) AddTask() (id string, err error) {
	u, err := uuid.NewRandom()
	if err != nil {
		fmt.Println(err)
		return
	}

	task.ID = u.String()
	task.Task_Status = 0

	result := orm.Eloquent.Create(&task)

	id = task.ID
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}

// 查询任务状态
func (task Task) CheckStatus() (results Task) {
	orm.Eloquent.Select([]string{"task_status"}).First(&results, "id = ?", task.ID)
	return 
}
