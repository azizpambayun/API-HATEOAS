package config

import (
	"log"

	"github.com/spf13/viper"
)

func Key() []byte {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	jwtSecret := viper.GetString("JWT_SECRET")
	log.Printf("Loaded JWT Secret: %s\n", jwtSecret)
	
	var JwtKey = []byte(jwtSecret)
	return JwtKey
}