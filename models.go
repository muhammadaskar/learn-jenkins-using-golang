package main

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Id    int
	Name  string
	Email string
}
