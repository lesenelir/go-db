package operation

import (
	"database/sql"
	"fmt"
)

func Delete(db *sql.DB) {
	exec, err := db.Exec("delete from person where user_id = ?", 2)
	if err != nil {
		fmt.Println("delete person failed, err:", err)
		return
	}

	row, err := exec.RowsAffected()
	if err != nil {
		fmt.Println("get person table RowsAffected failed, err:", err)
	}

	fmt.Println("delete person success, affected rows:", row)
}
