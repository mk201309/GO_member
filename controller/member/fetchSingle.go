package member

import (
	"fmt"
	"member_center/model"
	"member_center/structs"
	"strconv"

	"github.com/gin-gonic/gin"
)

// FetchSingleMember rr
func FetchSingleMember(c *gin.Context) {
	id := c.Params.ByName("id")

	fmt.Println(id)
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

	result, err := model.SQLOneMember(o)
	if err != nil {

	}

	c.JSON(200, gin.H{
		"status":  "Y",
		"message": result,
	})
}
