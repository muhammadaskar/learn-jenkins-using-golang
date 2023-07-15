package main

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Configuration struct {
	DB struct {
		Name     string `default:"db_test_golang"`
		User     string `default:"root"`
		Password string `default:"admin"`
		Host     string `default:"mysql-container"`
		Port     string `default:"3306"`
	}
}

var DB *gorm.DB

func InitDatabase() {
	var err error

	// config := Configuration{}
	// // err = configor.Load(&config, "config.yml") // Create a config.yml file with your MySQL credentials

	// // if err != nil {
	// // 	fmt.Println("Error loading database configuration:", err)
	// // 	os.Exit(1)
	// // }

	// connectionString := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
	// 	config.DB.User,
	// 	config.DB.Password,
	// 	config.DB.Host,
	// 	config.DB.Port,
	// 	config.DB.Name,
	// )

	DB, err = gorm.Open("mysql", "root:admind@tcp(172.17.0.3:3306)/db_test_golang")

	if err != nil {
		fmt.Println("Database connection failed:", err)
		os.Exit(1)
	}

	DB.AutoMigrate(&User{}) // Migrate User model to the database
}
