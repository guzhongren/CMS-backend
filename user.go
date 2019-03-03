package main

import (
	"net/http"

	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
)

type User struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	RoleId string `json:"roleid"`
}

func (user User) getUserByName(userName string) (User, error) {
	err := db.QueryRow("select id, name, roleid from b_user where name=$1", userName).Scan(&user.ID, &user.Name, &user.RoleId)
	if err != nil {
		log.Warn("查询用户出错", err)
		return User{}, err
	}
	return user, nil
}

func (user User) getAllUsers(c echo.Context) error {
	rows, err := db.Query("select id, roleid ,name from b_user")
	if err != nil {
		log.Warn("查询出错")
		return err
	}
	var userList = []User{}
	for rows.Next() {
		user := User{}
		err := rows.Scan(&user.ID, &user.RoleId, &user.Name)
		if err != nil {
			log.Warn("处理查询结果出错", err)
			return c.JSON(http.StatusExpectationFailed, &Response{
				Success: false,
				Result:  "",
				Message: "查询出错",
			})
		}
		userList = append(userList, user)
	}
	return c.JSON(http.StatusOK, &Response{
		Success: true,
		Result:  userList,
		Message: "",
	})
}
