package member

import (
	"member_center/model"

	"github.com/gin-gonic/gin"
)

// FetchAllMember get all member
func FetchAllMember(c *gin.Context) {

	result, err := model.SQLListMember()

	if err != nil {
		c.JSON(200, gin.H{
			"code":    "40004",
			"message": "db error !! ",
		})
		return
	}

	c.JSON(200, gin.H{
		"code":    "200",
		"message": "create user success !!",
		"data":    result,
	})

}
