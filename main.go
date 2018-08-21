package main

import (
	"flag"
	"github.com/spf13/viper"
	"log"
	"gpstrace/db"
	"gpstrace/web"
	"net/http"
)

func main() {
	
	configPath := flag.String("c", "", "Configuration file path")
	flag.Parse()
	
	if *configPath != "" {
		viper.SetConfigFile(*configPath)
		
		if err := viper.ReadInConfig(); err != nil {
			log.Fatal("Unable to load config:", err)
			
		}
	}
	
	d, err := db.GetConnection()

	if err != nil {
		log.Fatal("Unable to set up database:", err)
	}

	app := web.New(d)
	
	err = http.ListenAndServe(viper.GetString("http.listen"), app)
	
	if err != nil {
		log.Fatal("Unable to start HTTP Service:", err)
	}
	
	
}
