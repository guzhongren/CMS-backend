package main

import (
	"net/http"

	"database/sql"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

var (
	db   *sql.DB
	conf *Conf
)

func main() {
	utils := Utils{}
	utils.LoadConfig()
	// log.Info(conf)

	log.SetLevel(log.DebugLevel)
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: conf.APP.CORS_Origins,
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	auth := Auth{}
	user := User{}
	material := Material{}
	var IsLoggedIn = middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(conf.Secret),
		Skipper:    auth.skipper,
	})
	dbInfo := conf.DB
	db = getDB(dbInfo.Host, dbInfo.Port, dbInfo.Username, dbInfo.Password, dbInfo.Db)
	defer db.Close()

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Hello! Welcome to CMS!")
	})

	h := &handler{}
	apiGroup := e.Group("/api/" + conf.Version)
	apiGroup.POST("/login", auth.Login)
	// 查询所有用户
	apiGroup.GET("/users", user.GetAllUsers, IsLoggedIn)
	apiGroup.GET("/users/:id", user.GetUser, IsLoggedIn)
	apiGroup.POST("/users", user.AddUser, IsLoggedIn)
	apiGroup.DELETE("/users/:id", user.DeleteUser, IsLoggedIn)
	apiGroup.PUT("/users/:id", user.UpdateUser, IsLoggedIn)
	apiGroup.PUT("/users/:id/resetPassword", user.ResetPassword, IsLoggedIn)
	apiGroup.POST("/materials", material.Add, IsLoggedIn)

	e.POST("/private", h.Private, IsLoggedIn)
	e.GET("/admin", h.Private, IsLoggedIn, isAdmin)
	e.Logger.Fatal(e.Start(conf.APP.Addr))

}
