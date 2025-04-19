package infrastructure

import (
	"fmt"
	"os"
	"time"

	Huma "github.com/danielgtaylor/huma/v2"
	Humagin "github.com/danielgtaylor/huma/v2/adapters/humagin"
	Cors "github.com/gin-contrib/cors"
	Gin "github.com/gin-gonic/gin"

	Config "contact-me/src/config"
	Utils "contact-me/src/utils"
)

type Server struct {
	Router *Gin.Engine
	Api    Huma.API
}

type Infrastructure struct {
	Server Server
}

func (self *Infrastructure) Close() {
	fmt.Println("Shutting down server...")
}

var shouldServeDocs = (os.Getenv("ENV") == "local")

var corsMiddleware = Cors.New(Cors.Config{
	AllowMethods:     Config.AppConfig.CORS.AllowMethods,
	AllowHeaders:     Config.AppConfig.CORS.AllowHeaders,
	ExposeHeaders:    Config.AppConfig.CORS.ExposeHeaders,
	AllowCredentials: Config.AppConfig.CORS.AllowCredentials,
	AllowOriginFunc: func(requestOrigin string) bool {
		maybeValidOriginMatch := Utils.Find[string](Config.AppConfig.CORS.AllowOrigins, func(allowedOrigin string, index int32) bool {
			return allowedOrigin == requestOrigin
		})

		if maybeValidOriginMatch != nil {
			return true
		} else {
			return false
		}
	},
	MaxAge: 12 * time.Hour,
})

func UseInfrastructure(config Config.Config) (*Infrastructure, error) {
	ginRouter := Gin.Default()
	ginRouter.Use(Gin.Recovery())
	ginRouter.Use(corsMiddleware)

	openApiConfig := Huma.DefaultConfig(config.API.Name, config.API.Version)
	openApiConfig.Servers = []*Huma.Server{{URL: "http://127.0.0.1:" + config.API.Port}}

	if !shouldServeDocs {
		openApiConfig.DocsPath = ""
	}

	api := Humagin.New(ginRouter, openApiConfig)

	server := Server{
		Router: ginRouter,
		Api:    api,
	}

	return &Infrastructure{
		Server: server,
	}, nil
}
