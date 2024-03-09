package models

import (
    "github.com/jinzhu/gorm"
)

type Comment struct {
	gorm.Model
	UserId    	uint
	ParentId    uint	`gorm:"default:0"`
	UserName    string
	Content     string
	Uid 		uint  	`gorm:"not null"`
	// Video		Video 	`gorm:"ForeignKey:Uid"`
}