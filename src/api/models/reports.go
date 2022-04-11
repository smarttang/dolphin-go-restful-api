package models

import (
	orm "dolphin/api/database"
)

type Report struct {
	ID			string `json:"id"`						// ID
	Report_Info	string `json:"report_info"`				// 任务扫描的结果
}

var Reports []Report

//设置表名
func (Report) TableName() string {
    return "report"
}

// 查询任务结果
func (report Report) GetReports() (results Report) {
	orm.Eloquent.First(&results, "id = ?", report.ID)
	return 
}
