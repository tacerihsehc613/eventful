package main

import (
	"flag"
	"fmt"
	"rabbit/bookingservice/listener"
	"rabbit/bookingservice/rest"
	"rabbit/lib/configuration"
	msgqueue_amqp "rabbit/lib/msgqueue/amqp"
	"rabbit/lib/persistence/dblayer"

	"github.com/streadway/amqp"
	//amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	confPath := flag.String("config", "./configuration/config.json", "path to config file")
	flag.Parse()
	config, _ := configuration.ExtractConfiguration(*confPath)

	fmt.Println("config.AMQPMessageBroker: ", config.AMQPMessageBroker)
	fmt.Println("config.DBConnection: ", config.DBConnection)

	dbhandler, err := dblayer.NewPersistenceLayer(config.Databasetype, config.DBConnection)
	if err != nil {
		panic(err)
	}
	conn, err := amqp.Dial(config.AMQPMessageBroker)
	if err != nil {
		panic(err)
	}

	eventListener, err := msgqueue_amqp.NewAMQPEventListner(conn, "events", "booking")
	if err != nil {
		panic(err)
	}
	eventEmitter, err := msgqueue_amqp.NewAMQPEventEmitter(conn, "events")
	if err != nil {
		panic(err)
	}
	processor := &listener.EventProcessor{EventListener: eventListener, Database: dbhandler}

	//without go, the program will block here

	// Use sync.WaitGroup to wait for the goroutine to finish
	//var wg sync.WaitGroup
	//wg.Add(2)

	// Run the processor in a goroutine
	/*go func() {
		defer wg.Done()
		if err := processor.ProcessEvents(); err != nil {
			log.Printf("Error processing events: %s", err)
		}
	}() */

	// Use a channel to signal the exit
	//exit := make(chan struct{}, 1)

	//Run the processor in a goroutine
	/*go func() {
		defer close(exit)
		if err := processor.ProcessEvents(); err != nil {
			fmt.Printf("Error processing events: %s\n", err)
		}
	}()*/
	go processor.ProcessEvents()

	rest.ServeAPI("0.0.0.0:8282", dbhandler, eventEmitter)
	//<-exit

	//Run the REST API in a goroutine
	/*go func() {
		defer wg.Done()
		rest.ServeAPI("0.0.0.0:8282", dbhandler, eventEmitter)
	}()*/
	//rest.ServeAPI("localhost:8282", dbhandler, eventEmitter)
	//rest.ServeAPI(config.RestfulEndpoint, dbhandler, eventEmitter)

	// Wait for the goroutine to finish before exiting
	//wg.Wait()
}
