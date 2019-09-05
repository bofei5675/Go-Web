package model

import (
	"fmt"
	"log"
)

// insert
func Insert(sql string, args ...interface{}) (int64, error) {
	/*
		Used an interface to insert data into sql statement
		then run sql statement by DB connection
	*/
	stmt, err := DB.Prepare(sql)
	// defer this function execution after this Insert() functions finished
	defer stmt.Close()
	if err != nil {
		return 1, nil
	}
	result, err := stmt.Exec(args...)
	if err != nil {
		return 1, nil
	}
	id, err := result.LastInsertId()

	if err != nil {
		return 1, err
	}

	return id, nil
}

// Delete
func Delete(sql string, args ...interface{}) {
	stmt, err := DB.Prepare(sql)
	defer stmt.Close()
	CheckErr(err, "Failed at setup of SQL statement.")
	result, err := stmt.Exec(args...)
	CheckErr(err, "Failed at setup of parameters.")
	num, err := result.RowsAffected()
	CheckErr(err, "Delete Failed")
	fmt.Printf("Delete successfully, affect %d rows \n", num)
}

// Update
func Update(sql string, args ...interface{}) {
	stmt, err := DB.Prepare(sql)
	defer stmt.Close()
	CheckErr(err, "Failed at setup of SQL statement")
	result, err := stmt.Exec(args...)
	CheckErr(err, "Failed at setup of parameters")
	num, err := result.RowsAffected()
	CheckErr(err, "Insert Failed")
	fmt.Printf("Update successfully, affect %d rows \n", num)

}

func CheckErr(err error, msg string) {
	if err != nil {
		log.Panicln(msg)
	}
}


