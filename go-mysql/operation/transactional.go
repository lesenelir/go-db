package operation

import (
	"database/sql"
	"fmt"
)

func Transaction(db *sql.DB) {
	conn, err := db.Begin()
	if err != nil {
		fmt.Println("Error in begin transaction ", err)
		return
	}

	exec, err := conn.Exec("insert into person(username, sex, email) values (?, ?, ?)", "s01", "man", "s01@qq.com")
	if err != nil {
		fmt.Println("Error in insert data ", err)
		return
	}

	id, err := exec.LastInsertId()
	if err != nil {
		fmt.Println("Error in get last insert id ", err)
		err = conn.Rollback()
		if err != nil {
			fmt.Println("rollback failed: ", err)
			return
		}
		return
	}

	fmt.Println("insert success: ", id)

	exec, err = conn.Exec("insert into person(username, sex, email)values(?, ?, ?)", "s02", "man", "s02@qq.com")
	if err != nil {
		fmt.Println("exec failed: ", err)
		if err = conn.Rollback(); err != nil {
			fmt.Println("rollback failed: ", err)
			return
		}
		return
	}

	id, err = exec.LastInsertId()
	if err != nil {
		fmt.Println("exec failed: ", err)
		if err = conn.Rollback(); err != nil {
			fmt.Println("rollback failed: ", err)
			return
		}
	}

	fmt.Println("insert success: ", id)

	conn.Commit()
}
