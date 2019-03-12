package main

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
)

type Material struct {
	ID           string `json:"id"`
	UserId       string `json:"user_id"`
	Name         string `json:"name"`
	Location     string `json:"location"`
	Type         string `json:"type"`
	Count        int64  `json:"count"`
	Provider     string `json:"provider"`
	ProviderLink string `json:"providerLink"`
	Images       string `json:"images"`
	CreateTime   int64  `json:"createTime"`
	UpdateTime   int64  `json:"updateTime"`
}

type MaterialResponse struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Location     string `json:"location"`
	Type         string `json:"type"`
	Count        int64  `json:"count"`
	Provider     string `json:"provider"`
	ProviderLink string `json:"providerLink"`
	Images       string `json:"images"`
	CreateTime   int64  `json:"createTime"`
	UpdateTime   int64  `json:"updateTime"`
	User         struct {
		UserId   string `json:"user_id"`
		UserName string `json:"user_name"`
	}
}

func (material Material) Add(c echo.Context) error {
	m := new(Material)
	userInfo := c.Get("user").(*jwt.Token)
	claims := userInfo.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	user := User{}
	getedUser, err := user.GetUserByName(name)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &Response{
			Success: false,
			Result:  "",
			Message: "参数错误，请填写正确的参数！",
		})
	}
	userId := getedUser.ID
	if err := c.Bind(m); err != nil {
		return c.JSON(http.StatusBadRequest, &Response{
			Success: false,
			Result:  "",
			Message: "参数错误",
		})
	}
	var utils = Utils{}
	m.ID = utils.GetGUID()
	m.CreateTime = time.Now().Unix()
	m.UserId = userId
	insertedMaterial, err := m.insert(*m)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &Response{
			Success: false,
			Result:  "",
			Message: "服务器内部错误",
		})
	}
	return c.JSON(http.StatusOK, &Response{
		Success: true,
		Result:  insertedMaterial,
		Message: "",
	})
}

func (m Material) insert(material Material) (Material, error) {
	stmt, err := db.Prepare(`INSERT INTO "public"."b_material" ("id", "userId", "location", "type", "count", "provider", "providerLink", "images", "name", "createTime", "updateTime", "price") VALUES ($1, $2, $3, $4, $5, $6, $7, $8, &9, &10, &11, &12)`)
	if err != nil {
		log.Warn("插入物料前错误，", err)
		return Material{}, err
	}
	defer stmt.Close()
	_, err = stmt.Exec(material.ID)
	if err != nil {
		log.Warn("插入物料错误时错误！", err)
		return Material{}, err
	}
	return material, nil
}
