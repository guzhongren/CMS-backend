package main

import (
	"database/sql"
	"net/http"
	"os"
	"time"

	"github.com/guzhongren/CMS-backend/docs"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	echoSwagger "github.com/swaggo/echo-swagger"
	// _ "github.com/guzhongren/CMS-backend/docs"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /api/v1

var db = new(sql.DB)
var conf = new(Conf)

func main() {
	utils := Utils{}
	utils.LoadConfig()
	// log.Info(conf)
	// log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
	e := echo.New()
	e.Server.ReadTimeout = time.Second * 5
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

	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:   conf.APP.StaticPath.Http,
		Browse: true,
	}))

	auth := Auth{}
	user := User{}
	role := Role{}
	material := Material{}
	system := System{}
	file := File{}
	var IsLoggedIn = middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(conf.Secret),
		// Skipper:    auth.skipper,
		Skipper: func(c echo.Context) bool {
			return true
		},
	})
	dbInfo := conf.DB
	db = getDB(dbInfo.Host, dbInfo.Port, dbInfo.Username, dbInfo.Password, dbInfo.Db)
	defer db.Close()

	docs.SwaggerInfo.Title = "CMS"
	docs.SwaggerInfo.Description = "This is a sample server Petstore server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:1234"
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.GET("/test", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Hello! Welcome to CMS!")
	})
	e.Static("/", "App")

	h := &handler{}
	apiGroup := e.Group("/api/" + conf.Version)
	apiGroup.Use(IsLoggedIn)
	// apiGroup.Static(conf.APP.StaticPath.Http, conf.APP.StaticPath.Local)
	// apiGroup.GET(conf.APP.StaticPath.Http+"/:id", file.Look)

	apiGroup.POST("/login", auth.Login)
	apiGroup.GET("/statistic", system.Statistic, IsLoggedIn)
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
	apiGroup.DELETE("/file/delete/:name", file.Delete)
	apiGroup.DELETE("/materials/:id", material.Delete, IsLoggedIn)
	apiGroup.PUT("/materials/:id", material.Update, IsLoggedIn)
	apiGroup.GET("/materials", material.GetAll, IsLoggedIn)
	apiGroup.GET("/materials/:id", material.GetOne, IsLoggedIn)

	e.POST("/private", h.Private, IsLoggedIn)
	e.GET("/admin", h.Private, IsLoggedIn, isAdmin)
	e.Logger.Fatal(e.Start(conf.APP.Addr))

}
