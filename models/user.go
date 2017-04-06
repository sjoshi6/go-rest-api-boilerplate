package models

import "github.com/jinzhu/gorm"

// User is a real world user
type User struct {
	gorm.Model
	Age       uint   `json:"age",gorm:"column:age;"`
	FirstName string `json:"first_name",gorm:"column:first_name;type:varchar(200):primary_key"`
	LastName  string `json:"last_name",gorm:"column:last_name;type:varchar(200):not null"`
}

// Users is a collection of multiple users
type Users []User
