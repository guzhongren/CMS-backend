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
	log.SetLevel(log.DebugLevel)
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))
	utils := Utils{}
	utils.LoadConfig()
	log.Info(conf)
	auth := auth{}
	user := User{}
	var IsLoggedIn = middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(conf.Secret),
		Skipper:    auth.skipper,
	})
	dbInfo := conf.DB
	db = getDB(dbInfo.Host, dbInfo.Port, dbInfo.Username, dbInfo.Password, dbInfo.Db)
	defer db.Close()

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Hello! Welecom to CMS!")
	})

	h := &handler{}
	apiGroup := e.Group("/api/" + conf.Version)
	apiGroup.POST("/login", auth.Login)
	apiGroup.GET("/users", user.getAllUsers, IsLoggedIn)
	e.POST("/private", h.Private, IsLoggedIn)
	e.GET("/admin", h.Private, IsLoggedIn, isAdmin)
	e.Logger.Fatal(e.Start(":1234"))

}
