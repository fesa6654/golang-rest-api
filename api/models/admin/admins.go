package models

import (
	"golang-rest-api/connections/mysql"
	"time"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Admin struct {
	Id        uint      `json:"id"`
	Name      string    `json:"name"`
	Surname   string    `json:"surname"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt"`
}

func init() {
	mysql.MySQLConnection()
	db = mysql.GetDB()
	db.AutoMigrate(&Admin{})
}

func (a *Admin) CreateNewAdmin() *Admin {

	db.NewRecord(a)
	db.Create(&a)

	return a
}

func GetAllAdmins() []Admin {

	var AllAdmins []Admin
	db.Find(&AllAdmins)

	return AllAdmins
}

func GetAdminById(Id int64) (*Admin, *gorm.DB) {

	var getAdmins Admin
	db := db.Where("id=?", Id).Find(&getAdmins)

	return &getAdmins, db
}

func DeleteAdminById(Id int64) Admin {

	var admin Admin
	db.Where("id=?", Id).Delete(admin)

	return admin
}
