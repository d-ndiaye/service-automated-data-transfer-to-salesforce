package main

import (
	flag "github.com/spf13/pflag"
	"log"
	"os"
	"service-automatisierte-daten-in-salesforce/internal"
	"service-automatisierte-daten-in-salesforce/pkg/config"
)

var configFile string

const file = "config/service.yaml"

func main() {
	flag.Parse()
	err, conf := config.Load(configFile)
	if err != nil {
		log.Println("Read config error")
		os.Exit(1)
	}
	db, err := internal.InitDb(conf.Mysql)
	if err != nil {
		log.Printf("Error %s db connection", err)
		return
	}
	repository := internal.NewRepository(db)
	log.Printf("Successfully connected to database: %s ", db.Name())

	fileManager := internal.NewFileManager(repository)
	watch := internal.NewWatcher(fileManager, repository)
	_, err = watch.WatchFolder(conf)
	if err != nil {
		log.Printf("Error WatchFolder: %s", err)
		return
	}
}

func init() {
	flag.StringVarP(&configFile, "config", "c", file, "this is the path and filename to the config file")
}
