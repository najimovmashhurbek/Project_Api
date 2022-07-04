package api

import (
	"github.com/gin-gonic/gin"
	"github.com/najimovmashhurbek/Project_Api/api-gateway_first/api/auth"
	"github.com/najimovmashhurbek/Project_Api/api-gateway_first/api/casbin"
	v1 "github.com/najimovmashhurbek/Project_Api/api-gateway_first/api/handlers/v1"
	"github.com/najimovmashhurbek/Project_Api/api-gateway_first/config"
	"github.com/najimovmashhurbek/Project_Api/api-gateway_first/pkg/logger"
	"github.com/najimovmashhurbek/Project_Api/api-gateway_first/services"
	"github.com/najimovmashhurbek/Project_Api/api-gateway_first/storage/repo"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	casbinN "github.com/casbin/casbin/v2"
	_ "github.com/najimovmashhurbek/Project_Api/api-gateway_first/api/docs"
)

// Option ...
type Option struct {
	Conf           config.Config
	Logger         logger.Logger
	ServiceManager services.IServiceManager
	RedisRepo      repo.RepositoryStorage
	Casbin         *casbinN.Enforcer
}

// New ...
func New(option Option) *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	jwtHandler := &auth.JwtHandler{
		SigninKey: option.Conf.SigninKey,
		Log:       option.Logger,
	}
	router.Use(casbin.NewJwtRoleStruct(option.Casbin, option.Conf, *jwtHandler))

	handlerV1 := v1.New(&v1.HandlerV1Config{
		Logger:         option.Logger,
		ServiceManager: option.ServiceManager,
		Cfg:            option.Conf,
		Redis:          option.RedisRepo,
	})

	api := router.Group("/v1")
	api.POST("/users", handlerV1.CreateUser)
	api.GET("/users/:id", handlerV1.GetUser)
	api.GET("/users", handlerV1.ListUsers)
	api.PUT("/users/:id", handlerV1.UpdateUser)
	api.DELETE("/users/:id", handlerV1.DeleteUser)
	api.POST("/users/register", handlerV1.Register)
	api.POST("/users/verfication", handlerV1.VerifyUser)
	api.POST("/users/login/:email/:password", handlerV1.Login)

	url := ginSwagger.URL("swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return router
}
