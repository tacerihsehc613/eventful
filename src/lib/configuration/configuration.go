package configuration

import (
	"os"
	"rabbit/lib/persistence/dblayer"
)

var (
	DBTypeDefault       = dblayer.DBTYPE("mongodb")
	DBConnectionDefault = "mongodb://127.0.0.1:27017"
	//RestfulEPDefault    = "localhost:8181"
	RestfulEPDefault = "0.0.0.0:8181"
	//RestfulEPDefault = "172.19.0.4:8181"
	//RestfulTLSEPDefault = "localhost:9191"
	RestfulTLSEPDefault = "0.0.0.0:9191"
	//RestfulTLSEPDefault      = "172.19.0.4:9191"
	AMQPMessageBrokerDefault = "amqp://guest:guest@localhost:5672"
)

type ServiceConfig struct {
	Databasetype       dblayer.DBTYPE `json:"databasetype"`
	DBConnection       string         `json:"dbconnection"`
	RestfulEndpoint    string         `json:"restfulapi_endpoint"`
	RestfulTLSEndPoint string         `json:"restfulapi_tlsendpoint"`
	AMQPMessageBroker  string         `json:"amqp_message_broker"`
}

func ExtractConfiguration(filename string) (ServiceConfig, error) {
	conf := ServiceConfig{
		DBTypeDefault,
		DBConnectionDefault,
		RestfulEPDefault,
		RestfulTLSEPDefault,
		AMQPMessageBrokerDefault,
	}
	//fmt.Println("filename: ", filename)
	//file, err := os.Open(filename)
	// err := godotenv.Load()
	// if err != nil {
	// 	fmt.Println("Configuration file not found. Continuing with default values.")
	// 	return conf, err
	// }
	err := error(nil)
	//err = json.NewDecoder(file).Decode(&conf)
	if broker := os.Getenv("AMQP_URL"); broker != "" {
		conf.AMQPMessageBroker = broker
	}
	if amqp := os.Getenv("AMQP_BROKER_URL"); amqp != "" {
		conf.AMQPMessageBroker = amqp
	}
	if mongo := os.Getenv("MONGO_URL"); mongo != "" {
		conf.DBConnection = mongo
	}
	return conf, err
}
