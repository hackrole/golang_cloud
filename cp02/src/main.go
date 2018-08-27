package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/golang_cloud/cp02/src/eventserice/rest"
	"github.com/golang_cloud/cp02/src/lib/configuration"
	"github.com/golang_cloud/cp02/src/lib/persistence/dblayer"
)

func main() {
	confPath := flag.String("conf", `.\configuration\config.json`, "flag to set the path to the configuration json file")
	flag.Parse
	config, _ := configuration.ExtractConfiguration(*confPath)

	fmt.Println("Conecting to database")
	dbhandler, _ := dblayer.NewPersistenceLayer(config.Databasetype, config.DBConnection)

	// restfull api start
	log.Fatal(rest.ServeAPI(config.RestfulEndpoint, dbhandler))
}
