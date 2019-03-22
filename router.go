package main

import (
	"member_center/controller/member"

	"github.com/gin-gonic/gin"
)

// InitRouter router init
func InitRouter(r *gin.Engine) {
	v1 := r.Group("/api/v1/member")
	{
		// 註冊
		v1.POST("/", member.SignUpMember)

		// 會員列表
		v1.GET("/", member.FetchAllMember)

		// 查詢單人資料 by id
		v1.GET("/:id", member.FetchSingleMember)

		// 編輯會員資料
		v1.POST("/edit/:id", member.UpdateMember)

		// 刪除會員
		v1.POST("/delete/:id", member.DeleteMember)
	}
}
