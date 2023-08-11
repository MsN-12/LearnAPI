package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type Data struct {
	Id   int
	Name string
}

func main() {
	connectionString := fmt.Sprintf("%v:%v@tcp(127.0.0.1:3306)/%v", DBUser, DBPassword, DBName)
	db, err := sql.Open("mysql", connectionString)
	checkError(err)
	defer db.Close()

	result, err := db.Exec("insert into data values(4, 'xyz')")
	checkError(err)
	lastInsertId, err := result.LastInsertId()
	fmt.Println("lansInsertedId: ", lastInsertId)
	rowsAffected, err := result.RowsAffected()
	fmt.Println("rowsAffected: ", rowsAffected)
	checkError(err)

	rows, err := db.Query("SELECT * from data")
	checkError(err)
	for rows.Next() {
		var data Data
		err := rows.Scan(&data.Id, &data.Name)
		checkError(err)
		fmt.Println(data)
	}
}

func checkError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
