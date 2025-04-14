package infrastructure

import (
	"fmt"
	"time"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humagin"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"contact-me/src/config"
	"contact-me/src/utils"
)

type Server struct {
	Router *gin.Engine
	Api    huma.API
}

type Infrastructure struct {
	Server Server
}

func (self *Infrastructure) Close() {
	fmt.Println("Shutting down server...")
}

func UseInfrastructure(config config.Config) (*Infrastructure, error) {
	ginRouter := gin.Default()
	ginRouter.Use(gin.Recovery())
	ginRouter.Use(cors.New(cors.Config{
		AllowMethods:     config.CORS.AllowMethods,
		AllowHeaders:     config.CORS.AllowHeaders,
		ExposeHeaders:    config.CORS.ExposeHeaders,
		AllowCredentials: config.CORS.AllowCredentials,
		AllowOriginFunc: func(requestOrigin string) bool {
			maybeValidOriginMatch := utils.Find[string](config.CORS.AllowOrigins, func(allowedOrigin string, index int32) bool {
				return allowedOrigin == requestOrigin
			})

			if maybeValidOriginMatch != nil {
				return true
			} else {
				return false
			}
		},
		MaxAge: 12 * time.Hour,
	}))

	openApiConfig := huma.DefaultConfig(config.API.Name, config.API.Version)
	openApiConfig.Servers = []*huma.Server{{URL: "http://127.0.0.1:" + config.API.Port}}

	api := humagin.New(ginRouter, huma.DefaultConfig("Contact Me API", "1.0.0"))

	server := Server{
		Router: ginRouter,
		Api:    api,
	}

	return &Infrastructure{
		Server: server,
	}, nil
}
