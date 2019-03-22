package structs

import "time"

type (

	// Person 會員註冊參數
	Person struct {
		ID       uint
		Account  string
		Name     string
		Password string
	}

	// User 包括了 User 的字段类型
	User struct {
		MID         uint
		MAccount    string
		MName       string
		MPassword   string
		MCreatedate time.Time
	}

	// TransformedUser 代表格式化的 member 结构体
	TransformedUser struct {
		ID         uint
		Account    string
		Name       string
		Password   string
		CreateDate time.Time
	}

	OneUser struct {
		ID uint
	}
)
