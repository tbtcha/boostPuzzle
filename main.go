package main

import (
	"boostPuzzle/server/rpc"
	"boostPuzzle/server/xendpoint"
	"boostPuzzle/server/xservice"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"net/http"
	"os"
)

var (
	port           = ":8080"
	configPath     = ".env"
	version        = "0.0.1"
	boostyImgUrl   = ""
	boostyAlbumUrl = ""
)

func parseEnvFile() {
	// Parse config file (.env) if path to it specified and populate env vars
	err := godotenv.Overload(configPath)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	// Read config from env vars
	port = os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	}

	boostyImgUrl = os.Getenv("BOOSTY_IMG")
	boostyAlbumUrl = os.Getenv("BOOSTY_MEDIA")
}

func setupRouter(factory xendpoint.EndpointFactory) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.StaticFile("/favicon.ico", "server/static/favicon.ico")
	router.GET("/album/:username", factory.GetAlbumHTML())
	v1 := router.Group("/api/v1")
	{
		v1.GET("/album/:username", factory.GetAlbumRest())
	}
	return router
}

func main() {
	parseEnvFile()
	client := http.Client{}
	fmt.Println("App started...")
	newRpc := rpc.NewRpc(boostyAlbumUrl, &client)
	newService := xservice.NewService(newRpc)
	newEndpointFactory := xendpoint.NewEndpointFactory(newService)
	router := setupRouter(newEndpointFactory)
	err := router.Run(port)
	if err != nil {
		return
	}
}
