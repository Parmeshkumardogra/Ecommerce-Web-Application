package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"github.com/BMS/database"
	"github.com/BMS/routes"
	"github.com/BMS/config"
)

func main() {
	//trying to setup connection
	err := config.LoadConfig()
	if err != nil{
		log.Fatal("Error in loading config",err.Error());
	}
	os.Setenv("GIN_MODE", "release");
	err = database.ConnectRedis();
	if err !=nil {
		log.Fatalf("Error in connecting to Redis: %v",err);
	}
	err = database.ConnectDB();
	if err != nil {
		log.Fatalf("Error in connecting to MongoDB: %v",err);
	}

	router := routes.SetRoutes();
	// start the server in a separate goroutine
	go func(){
		log.Println("Listenning and Serving on port: 8080")
		err = router.Run(":8080")
		if err != nil {
			fmt.Println("Shutting down the server due to error");
		}
	}()
	
	quit := make(chan os.Signal,1);
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM);
	<-quit

	log.Println("Shutting down gracefully...")
	database.DisconnectRedis();
	database.DisconnectDB();
}
