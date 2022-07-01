package v1

import (
	_"errors"
	_"net/http"

	
	_"github.com/dgrijalva/jwt-go"
	_"github.com/gin-gonic/gin"
	"github.com/najimovmashhurbek/Project_Api/api-gateway_first/api/auth"
	_"github.com/najimovmashhurbek/Project_Api/api-gateway_first/api/model"
	"github.com/najimovmashhurbek/Project_Api/api-gateway_first/config"
	"github.com/najimovmashhurbek/Project_Api/api-gateway_first/pkg/logger"
	"github.com/najimovmashhurbek/Project_Api/api-gateway_first/services"
	"github.com/najimovmashhurbek/Project_Api/api-gateway_first/storage/repo"
)

type handlerV1 struct {
	log            logger.Logger
	serviceManager services.IServiceManager
	cfg            config.Config
	redisStorage   repo.RepositoryStorage
	jwtHandler     auth.JwtHandler
}

// HandlerV1Config ...
type HandlerV1Config struct {
	Logger         logger.Logger
	ServiceManager services.IServiceManager
	Cfg            config.Config
	Redis          repo.RepositoryStorage
	jwtHandler     auth.JwtHandler
}

// New ...
func New(c *HandlerV1Config) *handlerV1 {
	return &handlerV1{
		log:            c.Logger,
		serviceManager: c.ServiceManager,
		cfg:            c.Cfg,
		redisStorage:   c.Redis,
		jwtHandler:     c.jwtHandler,
	}
}

/*func CheckClaims(h *handlerV1, c *gin.Context) jwt.MapClaims {
	var (
		ErrUnauthotized = errors.New("unauthorized")
		authorization   model.JwtRequestModel
		claims          jwt.MapClaims
		err             error
	)

	authorization.Token = c.GetHeader("Authorization")
	if c.Request.Header.Get("Authorization") == "" {
		c.JSON(http.StatusUnauthorized, ErrUnauthotized)
		h.log.Error()
	}
}
*/