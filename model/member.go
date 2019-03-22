package model

import (
	"fmt"
	"io/ioutil"
	"log"
	"member_center/structs"

	"github.com/jinzhu/gorm"
	"gopkg.in/yaml.v2"
)

var db *gorm.DB

func connectDB() (db *gorm.DB, err error) {

	conf := new(structs.Config)

	// get db setting
	yamlFile, err := ioutil.ReadFile("./config/local.yaml")

	log.Println("yamlFile:", string(yamlFile))

	if err != nil {
		log.Printf("yamlFile.Get err #%v ", err)
	}

	err = yaml.Unmarshal(yamlFile, conf)

	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	log.Println("conf", conf.SQLConf.User)

	//open a db connection
	db, err = gorm.Open("mysql", conf.SQLConf.User+":"+conf.SQLConf.Password+"@/"+conf.SQLConf.Name+"?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic("failed to connect database")
	}

	// 全局禁用表名复数
	db.SingularTable(true)

	// show sql exec command
	db.LogMode(true)

	return db, nil

}

// SQLSignUpMember Add member
func SQLSignUpMember(p *structs.Person) (err error) {
	db, err := connectDB()
	if err != nil {
		return err
	}
	defer db.Close()

	user := structs.User{MAccount: p.Account, MName: p.Name, MPassword: p.Password}

	if err := db.Create(&user).Error; err != nil {
		return err
	}
	return err

}

// SQLListMember get All member
func SQLListMember() (userlist *[]structs.TransformedUser, err error) {
	var user []structs.User
	var _user []structs.TransformedUser

	db, err := connectDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	if err := db.Find(&user).Error; err != nil {
		return nil, err
	}

	for _, item := range user {
		_user = append(_user, structs.TransformedUser{
			ID:         item.MID,
			Account:    item.MAccount,
			Name:       item.MName,
			Password:   item.MPassword,
			CreateDate: item.MCreatedate,
		})
	}

	fmt.Println(&_user)
	return &_user, err

}

// SQLOneMember get one member data
func SQLOneMember(o *structs.OneUser) (userlist *[]structs.TransformedUser, err error) {
	var user []structs.User
	var _user []structs.TransformedUser
	db, err := connectDB()

	if err != nil {
		return nil, err
	}
	defer db.Close()

	if err := db.Where("m_id = ?", o.ID).Find(&user).Error; err != nil {
		return nil, err
	}

	for _, item := range user {
		_user = append(_user, structs.TransformedUser{
			ID:         item.MID,
			Account:    item.MAccount,
			Name:       item.MName,
			Password:   item.MPassword,
			CreateDate: item.MCreatedate,
		})
	}

	fmt.Println(&_user)
	return &_user, err
}

// SQLEditMember edit member data
func SQLEditMember(p *structs.Person) (err error) {
	db, err := connectDB()

	if err != nil {
		return err
	}
	defer db.Close()

	user := structs.User{MAccount: p.Account, MName: p.Name, MPassword: p.Password}

	fmt.Println(user)

	if err := db.Model(&user).Where("m_id = ?", p.ID).Updates(user).Error; err != nil {
		return err
	}
	return err
}

// SQLDeleteMember delete member
func SQLDeleteMember(p *structs.OneUser) (err error) {
	db, err := connectDB()
	if err != nil {
		return err
	}
	defer db.Close()

	user := structs.User{}

	if err := db.Delete(&user, "m_id = ?", p.ID).Error; err != nil {
		return err
	}
	return err
}
