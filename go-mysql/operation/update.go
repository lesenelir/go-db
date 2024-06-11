package operation

import (
	"database/sql"
	"fmt"
)

func Update(db *sql.DB) {
	exec, err := db.Exec("update person set username = ?, sex = ? where user_id = ?", "Tom-update", "sex-update", 1)
	if err != nil {
		fmt.Println("update person failed, err:", err)
		return
	}

	row, err := exec.RowsAffected() // be affected rows number
	if err != nil {
		fmt.Println("get person table RowsAffected failed, err:", err)
		return
	}

	fmt.Println("update person success, affected rows:", row)
}
