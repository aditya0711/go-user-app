package main

import (
	"fmt"

	employee "api/pkg/api"

	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("dev")         // name of config file (without extension)
	viper.AddConfigPath("$PWD/config") // path to look for the config file in
	err := viper.ReadInConfig()        // Find and read the config file
	if err != nil {                    // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})
	setupLogging()

	r := gin.Default()

	r.GET("/employees", employee.FindAll)
	r.GET("/employee", employee.FindAll)

	r.Run()
}

func setupLogging() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	log.Print("hello world")
	log.Print("hello world")
	log.Printf("JSON", "Logging to a file in Go!")
}
