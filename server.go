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
	db *sql.DB
)

func main() {
	log.SetLevel(log.DebugLevel)
	// log.SetFormatter(&log.JSONFormatter{})
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	//CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))
	auth := auth{}
	var IsLoggedIn = middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte("secret"),
		Skipper:    auth.skipper,
	})
	db = getDB("47.95.247.139", 5432, "postgres", "000000", "cms")
	defer db.Close()

	e.GET("/", func(c echo.Context) error {
		user := User{}
		user.getAllUser(db)
		return c.String(http.StatusOK, "Hello, world")
	})

	h := &handler{}

	e.POST("/login", auth.Login)
	e.POST("/private", h.Private, IsLoggedIn)
	e.GET("/admin", h.Private, IsLoggedIn, isAdmin)
	e.Logger.Fatal(e.Start(":1234"))

}
