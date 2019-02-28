package main

import (
	"database/sql"

	log "github.com/sirupsen/logrus"
)

type User struct {
	Id     int    `json: "id`
	Name   string `json: "name"`
	RoleId string `json: "roleid"`
}

func (user User) getAllUser(db *sql.DB) ([]User, error) {
	rows, err := db.Query("select * from b_user")
	if err != nil {
		log.Warn("查询出错")
		return nil, err
	}
	var userList = []User{}
	for rows.Next() {
		user := User{}
		err := rows.Scan(&user.Name, &user.RoleId, &user.Id)
		if err != nil {
			log.Warn("处理查询结果出错")
			return nil, err
		}
		userList = append(userList, user)
	}
	return userList, nil

}
