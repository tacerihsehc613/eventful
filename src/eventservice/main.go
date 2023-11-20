package main

import (
	"flag"
	"fmt"
	"log"
	"rabbit/eventservice/rest"
	"rabbit/lib/configuration"
	msgqueue_amqp "rabbit/lib/msgqueue/amqp"
	"rabbit/lib/persistence/dblayer"

	"github.com/streadway/amqp"
)

func main() {
	confPath := flag.String("conf", `./configuration/config.json`, "flag to set the path to the configuration json file")
	flag.Parse()

	config, _ := configuration.ExtractConfiguration(*confPath)

	fmt.Println("config.AMQPMessageBroker: ", config.AMQPMessageBroker)
	fmt.Println("config.DBConnection: ", config.DBConnection)
	conn, err := amqp.Dial(config.AMQPMessageBroker)
	if err != nil {
		panic(err)
	}
	//emitter, err := msgqueue_amqp.NewAMQPEventEmitter(conn)
	emitter, err := msgqueue_amqp.NewAMQPEventEmitter(conn, "events")
	if err != nil {
		panic(err)
	}

	fmt.Println("Connecting to database")
	//fmt.Println("aa:", config.Databasetype, config.DBConnection)
	dbhandler, _ := dblayer.NewPersistenceLayer(config.Databasetype, config.DBConnection)

	httpErrChan, httptlsErrChan := rest.ServeAPI(config.RestfulEndpoint, config.RestfulTLSEndPoint, dbhandler, emitter)

	select {
	case err := <-httpErrChan:
		log.Fatal("HTTP Error: ", err)
	case err := <-httptlsErrChan:
		log.Fatal("HTTPS Error: ", err)
	}

}
