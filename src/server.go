package main

import (
    conf "dolphin/api/conf"
    "dolphin/api/router"
)

func main() {
    port := conf.GetEnvDefault("HOST_PORT", "8080")
    router := router.InitRouter()
    router.Run(":"+port)
}
