package main

import (
	"net/http"

	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
)

type File struct {
}

func (file File) Upload(c echo.Context) error {
	utils := Utils{}
	// 文件上传
	form, err := c.MultipartForm()
	if err != nil {
		log.Warn("获取form 出错！", err)

		return c.JSON(http.StatusInternalServerError, &Response{
			Success: false,
			Result:  "",
			Message: "服务器内部错误",
		})
	}
	files := form.File["images"]
	savedFileIDArr, err := utils.UploadFiles(files)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &Response{
			Success: false,
			Result:  "",
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, &Response{
		Success: true,
		Result:  savedFileIDArr,
		Message: "",
	})
}
