package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/merryChris/gameLord/api"
	"github.com/merryChris/gameLord/hangman"
	"github.com/merryChris/gameLord/utils"
	"github.com/spf13/viper"
)

var (
	configFile = flag.String("config", "", "Config File Path")
)

func main() {
	flag.Parse()

	config := viper.New()
	config.SetConfigType("yaml")
	config.SetConfigFile(*configFile)
	if err := config.ReadInConfig(); err != nil {
		log.Fatal(fmt.Sprintf("Loading Config Error: `%s`.", err.Error()))
	}
	dbConf := config.Sub("database")
	srvConf := config.Sub("server")
	modConf := config.Sub("module")

	utils.InitOrm(dbConf.Sub("mysql"))
	if err := api.RoutingUserHandler.Init(dbConf.Sub("redis")); err != nil {
		log.Fatal(fmt.Sprintf("Initializing `RoutingUserHandler` Error: `%s`.", err.Error()))
	}
	defer api.RoutingUserHandler.Close()
	router := api.NewRouter()

	if ok := modConf.GetBool("hangman"); ok {
		if err := hangman.RoutingHangmanHandler.Init(dbConf.Sub("redis")); err != nil {
			log.Fatal(fmt.Sprintf("Initializing `RoutingHangmanHandler` Error: `%s`.", err.Error()))
		}
		api.AddRouter(router, hangman.V1HangmanRoutes)
		defer hangman.RoutingHangmanHandler.Close()
	}

	server := &http.Server{Addr: srvConf.GetString("address"), Handler: router}
	go func() {
		log.Println("`GameLord` Server Started.")
		if err := server.ListenAndServe(); err != nil {
			log.Fatal(fmt.Sprintf("Server Listening Error: `%s`.", err.Error()))
		}
	}()

	stopChan := make(chan os.Signal)
	defer close(stopChan)
	signal.Notify(stopChan, os.Interrupt)

	<-stopChan
	log.Println("Shutting Down Server.")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	server.Shutdown(ctx)
	log.Println("Server Gracefully Stopped.")
}
