package main

import (
	"fmt"
	"gorm-mysql/operation"
	"gorm-mysql/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

func main() {
	err := utils.LoadEnv(".env")
	if err != nil {
		fmt.Println("load env error ...")
		return
	}
	password := os.Getenv("MYSQL_PASSWORD")
	dsn := fmt.Sprintf("root:%s@tcp(127.0.0.1:3306)/gorm-test?charset=utf8mb4&parseTime=True&loc=Local", password)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		//Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		fmt.Println("gorm open error: ", err)
		return
	}

	// basic operation
	operation.BasicOperation(db) // auto migrate, create, read, update, delete

	// user operation
	operation.UserOperation(db)

}
