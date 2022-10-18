package mysql

import (
	"fmt"

	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

func Database() {
	var DB *gorm.DB

	// dsn := "root@tcp(127.0.0.1:3306)/dumbflix_2?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := "{user}:{password}@tcp({Host}:{Port})/{Database name}?charset=utf8mb4&parseTime=True&loc=Local"

	DB, err := gorm.Open(mysql.Open(dsn), &gorm.DB{})

	if err != nil {
		panic(err)
	}

	fmt.Println("Database Has Connected")


}