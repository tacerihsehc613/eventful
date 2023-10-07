package main

import (
	"flag"
	"rabbit/bookingservice/listener"
	"rabbit/eventservice/rest"
	"rabbit/lib/configuration"
	msgqueue_amqp "rabbit/lib/msgqueue/amqp"
	"rabbit/lib/persistence/dblayer"

	"github.com/streadway/amqp"
)

func main() {
	confPath := flag.String("config", "./configuration/config.json", "path to config file")
	flag.Parse()
	config, _ := configuration.ExtractConfiguration(*confPath)

	dbhandler, err := dblayer.NewPersistenceLayer(config.Databasetype, config.DBConnection)
	if err != nil {
		panic(err)
	}
	conn, err := amqp.Dial(config.AMQPMessageBroker)
	if err != nil {
		panic(err)
	}
	eventListener, err := msgqueue_amqp.NewAMQPEventListner(conn, "event")
	if err != nil {
		panic(err)
	}
	eventEmitter, err := msgqueue_amqp.NewAMQPEventEmitter(conn)
	if err != nil {
		panic(err)
	}
	processor := &listener.EventProcessor{EventListener: eventListener, Database: dbhandler}
	processor.ProcessEvents()

	rest.ServeAPI(config.RestfulEndpoint, config.RestfulTLSEndPoint, dbhandler, eventEmitter)
}
