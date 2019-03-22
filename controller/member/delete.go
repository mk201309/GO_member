package member

import (
	"member_center/model"
	"member_center/structs"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DeleteMember(c *gin.Context) {
	id := c.Params.ByName("id")

	uid, err := strconv.ParseUint(id, 10, 32)

	if err != nil {
		c.JSON(200, gin.H{
			"status":  "N",
			"message": "error",
		})
	}

	o := &structs.OneUser{
		ID: uint(uid),
	}

	err = model.SQLDeleteMember(o)

	if err != nil {
		c.JSON(200, gin.H{
			"status":  "N",
			"message": "db error",
		})
		return
	}

	c.JSON(200, gin.H{
		"status":  "Y",
		"message": "delete success",
	})
}
