package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" // 导入mysql驱动
	"go-mysql/operation"
	"go-mysql/utils"
	"os"
)

func main() {
	err := utils.LoadEnv(".env")
	if err != nil {
		fmt.Println("load env error ...")
		return
	}

	password := os.Getenv("MYSQL_PASSWORD")
	// root:19970122@/test
	db, err := sql.Open("mysql", fmt.Sprintf("root:%s@/test", password))
	if err != nil {
		panic(err.Error())
		return
	}
	defer db.Close()

	// insert
	operation.Insert(db)
	// select
	operation.Select(db)
	// update
	operation.Update(db)
	// delete
	operation.Delete(db)
	// transition
	operation.Transaction(db)
}
