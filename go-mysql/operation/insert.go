package operation

import (
	"database/sql"
	"fmt"
)

func insertPerson(db *sql.DB, person Person) error {
	exec, err := db.Exec("insert into person(username, sex, email) values(?, ?, ?)", person.UserName, person.Sex, person.Email)
	if err != nil {
		return fmt.Errorf("insert person err: %v", err)
	}

	id, err := exec.LastInsertId()
	if err != nil {
		return fmt.Errorf("get person table last insert id failed, err: %v", err)
	}

	fmt.Println("insert person success, the id is:", id)
	return nil
}

func insertPlace(db *sql.DB, place Place) error {
	exec, err := db.Exec("insert into place(country, city, telcode) values(?, ?, ?)", place.Country, place.City, place.TelCode)
	if err != nil {
		return fmt.Errorf("insert place err: %v", err)
	}

	id, err := exec.LastInsertId()
	if err != nil {
		return fmt.Errorf("get place table last insert id failed, err: %v", err)
	}

	fmt.Println("insert place success, the id is:", id)
	return nil
}

func Insert(db *sql.DB) {
	person := Person{
		UserId:   1,
		UserName: "Tom",
		Sex:      "male",
		Email:    "123@qq.com",
	}

	err := insertPerson(db, person)
	if err != nil {
		fmt.Println(err)
		return
	}

	// ---------------------

	place := Place{
		Country: "China",
		City:    "Beijing",
		TelCode: 010,
	}

	if err = insertPlace(db, place); err != nil {
		fmt.Println(err)
		return
	}
	//err = insertPlace(db, place)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
}
