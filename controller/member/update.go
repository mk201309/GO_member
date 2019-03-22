package member

import (
	"member_center/model"
	"member_center/structs"
	"strconv"

	"github.com/gin-gonic/gin"
)

// UpdateMember edit member data
func UpdateMember(c *gin.Context) {

	p := &structs.Person{}
	id := c.Params.ByName("id")
	p.Account = c.PostForm("account")
	p.Password = c.PostForm("password")
	p.Name = c.PostForm("name")

	uid, err := strconv.ParseUint(id, 10, 32)
	p.ID = uint(uid)

	if err != nil {
		c.JSON(200, gin.H{
			"status":  "N",
			"message": "error",
		})
	}

	err = model.SQLEditMember(p)

	if err != nil {
		c.JSON(200, gin.H{
			"status":  "N",
			"message": "db error !! ",
		})
		return
	}

	c.JSON(200, gin.H{
		"status":  "Y",
		"message": "create user success !!",
	})
}
