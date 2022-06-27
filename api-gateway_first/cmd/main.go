package main

import (
	"github.com/najimovmashhurbek/Project_Api/api-gateway_first/api"
	"github.com/najimovmashhurbek/Project_Api/api-gateway_first/config"
	"github.com/najimovmashhurbek/Project_Api/api-gateway_first/pkg/logger"
	"github.com/najimovmashhurbek/Project_Api/api-gateway_first/services"
)

func main() {
	cfg := config.Load()
	log := logger.New(cfg.LogLevel, "api_gateway")

	serviceManager, err := services.NewServiceManager(&cfg)
	if err != nil {
		log.Error("gRPC dial error", logger.Error(err))
	}

	server := api.New(api.Option{
		Conf:           cfg,
		Logger:         log,
		ServiceManager: serviceManager,
	})

	if err := server.Run(cfg.HTTPPort); err != nil {
		log.Fatal("failed to run http server", logger.Error(err))
		panic(err)
	}
}

/*package token

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/najimovmashhurbek/Project_Api/api-gateway_first/pkg/logger"
)

type JWTHandler struct {
	Sub       string
	Iss       string
	Exp       string
	Iat       string
	Aud       []string
	Role      string
	SigninKey string
	log       logger.Logger
	Token     string
}

func (jwthandler *JWTHandler) GenerateAuthJWT() (access, refresh string, err error) {
	var(
	accessToken  *jwt.Token
	refreshToken *jwt.Token
	claims jwt.MapClaims
	)
	accessToken=jwt.New(jwt.SigningMethodHS256)
	refreshToken=jwt.New(jwt.SigningMethodHS256)

	claims=accessToken.Claims.(jwt.MapClaims)
	claims["iss"]=jwthandler.Iss
	claims["sub"]=jwthandler.Sub

	if os.Getenv("ENVIRONMENT")=="production"{
		claims["exp"]=time.Now().Add(time.Hour*500).Unix()
	}else if os.Getenv("ENVIRONMENT")=="staging"{
		claims["exp"]=time.Now().Add(time.Hour*500).Unix()
	}else{
		claims["exp"]=time.Now().Add(time.Hour*500).Unix()
	}

	claims["iat"]=time.Now().Unix()
	claims["role"]=jwthandler.Role
	claims["aud"]=jwthandler.Aud

	access,err=accessToken.SignedString([]byte(jwthandler.SigninKey))
	if err!=nil{
		jwthandler.log.
	}

}
*/