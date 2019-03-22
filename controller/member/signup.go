package member

import (
	"fmt"
	"member_center/model"
	"member_center/structs"

	"github.com/gin-gonic/gin"
)

// SignUpMember add member
func SignUpMember(c *gin.Context) {

	account := c.PostForm("account")
	password := c.PostForm("password")
	name := c.PostForm("name")

	p := &structs.Person{
		Account:  account,
		Password: password,
		Name:     name,
	}

	fmt.Println(account)

	err := model.SQLSignUpMember(p)

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
