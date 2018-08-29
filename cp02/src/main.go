package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/golang_cloud/cp02/src/eventserice/rest"
	"github.com/golang_cloud/cp02/src/lib/configuration"
	"github.com/golang_cloud/cp02/src/lib/persistence/dblayer"

	"github.com/streadway/amqp"
	msgqueue_amqp "todo.com/myevents/lib/msgqueue/amqp"
)

func main() {
	confPath := flag.String("conf", `.\configuration\config.json`, "flag to set the path to the configuration json file")
	flag.Parse
	config, _ := configuration.ExtractConfiguration(*confPath)

	conn, err := amqp.Dial(config.AMQPMessageBroker)
	if err != nil {
		panic(err)
	}
	emitter, err := msgqueue_amqp.NewAMQPEventEmitter(conn)
	if err != nil {
		panic(err)
	}

	fmt.Println("Conecting to database")
	dbhandler, _ := dblayer.NewPersistenceLayer(config.Databasetype, config.DBConnection)

	// restfull api start
	log.Fatal(rest.ServeAPI(config.RestfulEndpoint, dbhandler, emitter))
}
