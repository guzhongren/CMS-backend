package main

import (
	"net/http"
	"os"

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
	// log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	// XSS
	e.Use(middleware.Secure())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: conf.APP.CORS_Origins,
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	auth := Auth{}
	user := User{}
	role := Role{}
	material := Material{}
	file := File{}
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
	apiGroup.Static(conf.APP.StaticPath.Http, conf.APP.StaticPath.Local)

	apiGroup.POST("/login", auth.Login)
	// 查询所有用户
	apiGroup.GET("/users", user.GetAllUsers, IsLoggedIn)
	apiGroup.GET("/users/:id", user.GetUser, IsLoggedIn)
	apiGroup.POST("/users", user.AddUser, IsLoggedIn)
	apiGroup.DELETE("/users/:id", user.DeleteUser, IsLoggedIn)
	apiGroup.PUT("/users/:id", user.UpdateUser, IsLoggedIn)
	apiGroup.PUT("/users/:id/resetPassword", user.ResetPassword, IsLoggedIn)
	apiGroup.GET("/roles", role.GetAll, IsLoggedIn)
	apiGroup.GET("/materials/types", material.GetMaterialType, IsLoggedIn)
	apiGroup.GET("/materials/types/:id", material.GetMaterialTypeById, IsLoggedIn)
	apiGroup.POST("/materials", material.Add, IsLoggedIn)
	apiGroup.POST("/file/upload", file.Upload)
	apiGroup.POST("/file/delete/:name", file.Delete)
	apiGroup.DELETE("/materials/:id", material.Delete, IsLoggedIn)
	apiGroup.PUT("/materials/:id", material.Update, IsLoggedIn)
	apiGroup.GET("/materials", material.GetAll, IsLoggedIn)
	apiGroup.GET("/materials/:id", material.GetOne, IsLoggedIn)

	e.POST("/private", h.Private, IsLoggedIn)
	e.GET("/admin", h.Private, IsLoggedIn, isAdmin)
	e.Logger.Fatal(e.Start(conf.APP.Addr))

}
