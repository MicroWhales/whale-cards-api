package main

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

func loadDBConfig() (DB_CONNECTION_STRING string) {
	viper.SetConfigName("db") // config file name without extension
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config/") // config file path
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("fatal error config file: default \n", err)
		os.Exit(1)
	}

	// Declare var
	DB_NAME := viper.GetString("DB_NAME")
	DB_USERNAME := viper.GetString("DB_USERNAME")
	DB_PASSWORD := viper.GetString("DB_PASSWORD")
	DB_CONNECTION_STRING = "user=" + DB_USERNAME + " password=" + DB_PASSWORD + " dbname=" + DB_NAME + " sslmode=disable"

	// Print
	// fmt.Println("---------- DB Config Loaded ----------")
	// fmt.Println("URI :", DB_URI)
	// fmt.Println("DB_USERNAME :", DB_USERNAME)
	// fmt.Println("DB_PASSWORD :", DB_PASSWORD)
	// fmt.Println("DB_CONNECTION_STRING :", DB_CONNECTION_STRING)
	return
}
