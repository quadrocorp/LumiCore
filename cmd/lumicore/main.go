package main

// Author == quadrocorp
// Version == v1.0-alpha || Photon
// NOTE: Photon is the alpha version of this project, and implies an MVP.

// LumiCore is the center, the core of Lumi, the innovative and feature-rich
// Telegram Bot management platform

import (
	"LumiCore/internal/keychain"
	"fmt"
	"os"
)

func main() {
	//cfgManager := config.New(config.CoreConfig)
	//
	//cfg, err := cfgManager.Load()
	//if err != nil {
	//	fmt.Errorf("failed loading config file: %e", err)
	//}
	//
	//r := gin.Default()
	//
	//r.GET("/", func(c *gin.Context) {
	//	c.JSON(http.StatusOK, gin.H{
	//		"message": "Hello, Gin!",
	//	})
	//})
	//
	//address := cfg.Server.Host + ":" + strconv.Itoa(cfg.Server.Port)
	//r.Run(address)

	key, err := keychain.GenerateJWTSecret(32)
	if err != nil {
		fmt.Errorf("error: %e", err)
		os.Exit(0)
	}

	if err := keychain.WriteToFile(key, keychain.JWTMaster); err != nil {
		fmt.Printf("an error occurred: %s", err)
	}
}
