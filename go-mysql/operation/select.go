package operation

import (
	"database/sql"
	"fmt"
)

func Select(db *sql.DB) {
	var persons []Person

	query, err := db.Query("select * from person")
	if err != nil {
		fmt.Println("query person table failed, err:", err)
		return
	}
	defer query.Close()

	for query.Next() {
		var person Person
		err = query.Scan(&person.UserId, &person.UserName, &person.Sex, &person.Email)
		if err != nil {
			fmt.Println("scan person failed, err:", err)
			return
		}
		persons = append(persons, person)
	}

	fmt.Println(persons)
}
