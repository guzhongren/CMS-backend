package main

import (
	"net/http"

	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
)

type Role struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	RoleId int    `json:"roleId"`
}

type Roleesponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (role Role) GetAll(c echo.Context) error {
	roleList, err := role.getAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &Response{
			Success: false,
			Result:  "",
			Message: "所有用户获取错误",
		})
	}
	return c.JSON(http.StatusOK, &Response{
		Success: true,
		Result:  roleList,
		Message: "",
	})
}

func (role Role) getAll() ([]Roleesponse, error) {
	rows, err := db.Query(`SELECT br.id, br.name FROM public.b_role as br`)
	if err != nil {
		log.Warn("查询出错", err)
		return []Roleesponse{}, err
	}
	var roleList = []Roleesponse{}
	for rows.Next() {
		role := Roleesponse{}
		err := rows.Scan(&role.ID, &role.Name)
		if err != nil {
			log.Warn("处理查询结果出错", err)
			return []Roleesponse{}, err
		}
		roleList = append(roleList, role)
	}
	return roleList, nil
}
