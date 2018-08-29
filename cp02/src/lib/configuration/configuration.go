package configuration

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/golang_cloud/cp02/src/lib/persistence/dblayer"
)

var (
	DBTypeDefault       = dblayer.DBTYPE("mongodb")
	DBConnectionDefault = "mongodb://127.0.0.1"
	RestfulEPDefault    = "localhost:8181"
)

type ServiceConfig struct {
	DatabaseHandler   dblayer.DBTYPE `json:"databasetype"`
	DBConnection      string         `json:"dbconnection"`
	RestfulEndpoint   string         `json:"restfulapi_endpoint"`
	AMQPMessageBroker string         `json:"amqp_message_broker"`
}

func ExtractConfiguration(filename string) (ServiceConfig, error) {
	conf := ServiceConfig{
		DBTypeDefault,
		DBConnectionDefault,
		RestfulEndpoint,
	}

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Configuration file not found. Continuing with default values.")
		return conf, erro
	}

	err = json.NewDecoder(file).Decode(&conf)
	if broker := os.Getenv("AMQP_URL"); broker != "" {
		conf.AMQPMessageBroker = broker
	}
	return conf, err
}
