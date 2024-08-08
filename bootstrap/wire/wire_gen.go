// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package wire

import (
	"github.com/chongyanovo/go-zzz/bootstrap"
	"github.com/chongyanovo/go-zzz/bootstrap/internal"
)

// Injectors from wire.go:

func InitApp() (bootstrap.Application, error) {
	viper := internal.NewViper()
	config := internal.NewConfig(viper)
	db := internal.NewMysql(config)
	cmdable := internal.NewRedis(config)
	logger := internal.NewZap(config)
	engine := internal.NewServer(config)
	application := bootstrap.NewApplication(config, viper, db, cmdable, logger, engine)
	return application, nil
}