package config

import (
	"github.com/BMS/models"
	"os"
	"log"
	"encoding/json"
)
var Config models.Config

func LoadConfig() error{
	file , err := os.Open("config/config.json");
	if err != nil{
		log.Printf("error occured to open config json file %v",err);
		return err;
	}
	defer file.Close()
	decoder := json.NewDecoder(file);
	err = decoder.Decode(&Config);
	if err != nil {
		log.Printf("error occured in decoding the json %v", err);
		return err;
	}
	log.Println("Config loaded successfully");
	return nil;
}