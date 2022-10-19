package database

import (
	"BE-foodways/pkg/mysql"
	"BE-foodways/models"
	"fmt"
)

func RunMigration() {
	err := mysql.DB.AutoMigrate(&models.User{})

	if err != nil {
		fmt.Println(err)
		panic("Migration Failed")
	}

	fmt.Println("Migration Success")
}
