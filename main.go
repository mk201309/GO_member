package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {

	r := gin.Default()

	// 註冊 Router
	InitRouter(r)

	r.Run()
}
