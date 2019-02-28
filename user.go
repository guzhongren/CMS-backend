package main

import (
	"database/sql"
	"fmt"

	log "github.com/sirupsen/logrus"
)

type User struct {
	Id     int    `json: "id`
	Name   string `json: "name"`
	RoleId string `json: "roleId"`
}

func (user User) getAllUser(db *sql.DB) ([]User, error) {
	rows, err := db.Query("select name from b_user")
	if err != nil {
		log.Warn("查询出错")
		return nil, err
	}
	var userList = []User{}
	var nameList = []string{}
	log.Info(rows)
	fmt.Printf("%v", rows)
	for rows.Next() {
		// user := User{}
		var name string
		err := rows.Scan(&name)

		if err != nil {
			log.Warn("处理查询结果出错", err)
			return nil, err
		}
		nameList = append(nameList, name)
	}
	return userList, nil

}
