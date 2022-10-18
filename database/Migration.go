package database

import (
	"fmt"

	"gorm.io/driver/mysql"
)

func Migration() {
	err := mysql.DB

	if err != nil {
		fmt.Println(err)
		panic("Failed to Migrate")
	}

	fmt.Println("Migration Has Been Succes")
}