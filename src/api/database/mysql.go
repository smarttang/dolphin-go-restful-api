package database

import (
    _ "github.com/go-sql-driver/mysql"
    "github.com/jinzhu/gorm"
    "fmt"
    conf "dolphin/api/conf"
)

var Eloquent *gorm.DB

func init() {
    var err error

    username := conf.GetEnvDefault("MYSQL_USERNAME", "root")
    password := conf.GetEnvDefault("MYSQL_PASSWORD", "root")
    host := conf.GetEnvDefault("MYSQL_HOST", "127.0.0.1")
    port := conf.GetEnvDefault("MYSQL_PORT", "3306")

    Eloquent, err = gorm.Open("mysql", username+":"+password+"@tcp("+host+":"+port+")/jobs?charset=utf8mb4&parseTime=True&timeout=60ms")

    if err != nil {
        fmt.Printf("[*]mysql connect error %v", err)
    }

    if Eloquent.Error != nil {
        fmt.Printf("[*]database error %v", Eloquent.Error)
    }
}