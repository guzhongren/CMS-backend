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
	rows, err := db.Query("select id, roleid ,name from b_user")
	if err != nil {
		log.Warn("查询出错")
		return nil, err
	}
	var userList = []User{}
	for rows.Next() {
		user := User{}
		err := rows.Scan(&user.Id, &user.RoleId, &user.Name)
		if err != nil {
			log.Warn("处理查询结果出错", err)
			return nil, err
		}
		userList = append(userList, user)
	}
	log.Info(userList)
	return userList, nil

}
