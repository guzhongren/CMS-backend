package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

type System struct {
}
type SystemResponse struct {
	UserCount     int64 `json:"userCount"`
	MaterialCount int64 `json:"materialCount"`
	MaterialPrice int64 `json:"price"`
	PersonalCount int64 `json:"personalCount"`
}

func (system System) Statistic(c echo.Context) error {
	utils := Utils{}
	currentUser, err := utils.GetUserFromContext(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &Response{
			Success: false,
			Result:  "",
			Message: "无效的请求",
		})
	}
	userCount, err := system.getUserCount()
	if err != nil {
		return c.JSON(http.StatusBadRequest, &Response{
			Success: false,
			Result:  "",
			Message: "无效的请求",
		})
	}
	materialCount, price, err := system.getMaterialCount()
	if err != nil {
		return c.JSON(http.StatusBadRequest, &Response{
			Success: false,
			Result:  "",
			Message: "无效的请求",
		})
	}
	personalCount, err := system.getPersonalCount(currentUser.ID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &Response{
			Success: false,
			Result:  "",
			Message: "无效的请求",
		})
	}
	return c.JSON(http.StatusOK, &Response{
		Success: true,
		Result: SystemResponse{
			userCount, materialCount, price, personalCount,
		},
		Message: "",
	})

}
func (system System) getPersonalCount(id string) (int64, error) {
	var count int64 = 0
	err := db.QueryRow(`SELECT count("id")FROM public.b_material as m  where m."ownerId"=$1`, id).Scan(&count)
	if err != nil {
		log.Warn("查询出错", err)
		return 0, err
	}
	return count, nil
}

func (system System) getUserCount() (int64, error) {
	var count int64 = 0
	err := db.QueryRow(`SELECT count("id") FROM public.b_user as b where b."isdeleted"=$1`, false).Scan(&count)
	if err != nil {
		log.Warn("查询出错", err)
		return 0, err
	}
	return count, nil
}
func (system System) getMaterialCount() (int64, int64, error) {
	var count int64 = 0
	var totalPrice int64 = 0
	err := db.QueryRow(`SELECT count("id"), sum(price) FROM public.b_material`).Scan(&count, &totalPrice)
	if err != nil {
		log.Warn("查询出错", err)
		return 0, 0, err
	}
	return count, totalPrice, nil
}
