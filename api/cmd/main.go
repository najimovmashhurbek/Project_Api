package main

import (
	"fmt"

	"github.com/casbin/casbin/v2"
	defaultrolemanager "github.com/casbin/casbin/v2/rbac/default-role-manager"
	"github.com/casbin/casbin/v2/util"
	"github.com/gomodule/redigo/redis"
	"github.com/najimovmashhurbek/Project_Api/api-gateway_first/api"
	"github.com/najimovmashhurbek/Project_Api/api-gateway_first/config"
	"github.com/najimovmashhurbek/Project_Api/api-gateway_first/pkg/logger"
	"github.com/najimovmashhurbek/Project_Api/api-gateway_first/services"
	rds "github.com/najimovmashhurbek/Project_Api/api-gateway_first/storage/redis"
)

func main() {
	cfg := config.Load()
	log := logger.New(cfg.LogLevel, "api")

	// psqlString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s",
	// 	cfg.PostgresHost,
	// 	cfg.PostgresPort,
	// 	cfg.PostgresUser,
	// 	cfg.PostgresPassword,
	// 	cfg.PostgressDatabase,
	// )

	// _, err := gormadapter.NewAdapter("postgres", psqlString, true)
	// if err != nil {
	// 	log.Error("new adapter error", logger.Error(err))
	// 	return
	// }

	casbinEnforcer, err := casbin.NewEnforcer(cfg.CasbinConfigPath, "./config/policy_defenition.csv")
	if err != nil {
		log.Error("new enforcer error", logger.Error(err))
		return
	}
	err = casbinEnforcer.LoadPolicy()
	if err != nil {
		log.Error("new load policy error", logger.Error(err))
		return
	}

	pool := redis.Pool{
		MaxIdle:   80,
		MaxActive: 12000,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", fmt.Sprintf("%s:%d", cfg.RedisHost, cfg.RedisPort))
			if err != nil {
				panic(err.Error())
			}
			return c, err
		},
	}

	redisRepo := rds.NewRedisRepo(&pool)
	casbinEnforcer.GetRoleManager().(*defaultrolemanager.RoleManager).AddMatchingFunc("KeyMatch", util.KeyMatch)
	casbinEnforcer.GetRoleManager().(*defaultrolemanager.RoleManager).AddMatchingFunc("KeyMatch3", util.KeyMatch3)

	serviceManager, err := services.NewServiceManager(&cfg)
	if err != nil {
		log.Error("gRPC dial error", logger.Error(err))
	}

	server := api.New(api.Option{
		Conf:           cfg,
		Logger:         log,
		Casbin:         casbinEnforcer,
		ServiceManager: serviceManager,
		RedisRepo:      redisRepo,
	})

	if err := server.Run(cfg.HTTPPort); err != nil {
		log.Fatal("failed to run http server", logger.Error(err))
		panic(err)
	}
}
