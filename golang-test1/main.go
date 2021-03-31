package main

import (
	"encoding/json"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)


type User struct {
	ID 			int `json:"id"`
	Username 	string `json:"username"`
	Parent 		NullString `json:"parent"`
}

type NullString struct {
	sql.NullString
}

func (ns *NullString) MarshalJSON() ([]byte, error) {
	if !ns.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(ns.String)
}

func main() {
	fmt.Println("Running SQL")
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/test")
	defer db.Close()

	checkErr(err)

	var version string

	errVer := db.QueryRow("SELECT VERSION()").Scan(&version)

	checkErr(errVer)

	fmt.Println(version)

	// Menjalankan Query
	results, err := db.Query(`
	SELECT
    		child.id,
    		child.username,
    		parent.username AS 'parentUsername'
	FROM
    	USER child
    	left join user parent on child.parent = parent.ID`)
	checkErr(err)

	for results.Next() {
		var user User
		err = results.Scan(&user.ID, &user.Username, &user.Parent)
		checkErr(err)

		log.Printf("user instance := %#v\n", user)
		userJSON, err := json.Marshal(&user)
		if err != nil {
			log.Fatal(err)
		} else {
			log.Printf("json := %s\n\n", userJSON)
		}
	}

	err = results.Err()
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
