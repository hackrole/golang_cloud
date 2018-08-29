package main

import (
	"flag"

	"github.com/streadway/amqp"
	"todo.com/myevents/lib/configuration"
	msgqueue_amqp "todo.com/myevents/lib/msgqueue/amqp"
)

func main() {
	confPath := flag.String("config", "./configuration/config.json", "path to config file")
	flag.Parse
	config := configuration.ExtractConfiguration(*confPath)

	dblayer, err := dblayer.NewPersistenceLayer(config.Databasetype, config.DBConnection)
	if err != nil {
		panic(err)
	}

	conn, err := amqp.Dial(config.AMQPMessageBroker)
	if err != nil {
		panic(err)
	}

	EventListner, err := msgqueue.NewEventListener(conn)
	if err != nil {
		panic(err)
	}

	processor := &Listener.EventProcessor{EventListner, dblayer}
	go processor.ProcessEvents()

	rest.ServeAPI(config.RestfulEndpoint, dbhandler, eevntEmitter)
}
