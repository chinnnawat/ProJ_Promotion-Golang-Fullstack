package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/chinnnawat/ProJ_Promotion-Golang/backend/api/middlewares"
	"github.com/chinnnawat/ProJ_Promotion-Golang/backend/api/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func initViper() {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	viper.AddConfigPath(".")
	viper.SetEnvPrefix("demo")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("unable to initialize viper: %w", err))
	}
	log.Println("viper config initialized")
}

func main() {
	initViper()
	app := fiber.New(fiber.Config{
		AppName:      "Promotion",
		ServerHeader: "Fiber",
	})
	middlewares.InitFiberMiddlewares(app, routes.InitPublicRoutes, nil)

	var listenIp = viper.GetString("ListenIP")
	var listenPort = viper.GetString("ListenPort")

	log.Printf("will listen on %v:%v", listenIp, listenPort)

	err := app.Listen(fmt.Sprintf("%v:%v", listenIp, listenPort))
	log.Fatal(err)
}
