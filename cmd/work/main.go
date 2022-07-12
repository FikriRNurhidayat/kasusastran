package main

import (
	"flag"
	"io/ioutil"
	"os"
	"os/signal"
	"syscall"

	"github.com/fikrirnurhidayat/kasusastran/app/config"
	"github.com/fikrirnurhidayat/kasusastran/app/domain/event"
	"github.com/fikrirnurhidayat/kasusastran/app/driver/nsq"
	"github.com/fikrirnurhidayat/kasusastran/app/listener"
	"google.golang.org/grpc/grpclog"
)

var queue = flag.String("queue", "*", "The event to listen, default: '*'")

func main() {
	var consumers []nsq.Consumer

	log := grpclog.NewLoggerV2(os.Stdout, ioutil.Discard, ioutil.Discard)
	broker := nsq.NewConnection(config.GetNSQLookupAddress(), config.GetNSQAddress())

	seratListener := listener.NewSeratListener(log)

	// Set consumer listener
	if *queue == event.SERAT_CREATED_TOPIC || *queue == "*" {
		onercc, err := broker.AddEventListener(event.SERAT_CREATED_TOPIC, listener.ECHO_CHANNEL, seratListener.CreatedEventListener)

		if err != nil {
			log.Fatalf("cannot listen to user consumer: %v", err)
		}

		consumers = append(consumers, onercc)
	}

	if *queue == event.SERAT_RETRIEVED_TOPIC || *queue == "*" {
		onercc, err := broker.AddEventListener(event.SERAT_RETRIEVED_TOPIC, listener.ECHO_CHANNEL, seratListener.RetrievedEventListener)

		if err != nil {
			log.Fatalf("cannot listen to user consumer: %v", err)
		}

		consumers = append(consumers, onercc)
	}

	if *queue == event.SERAT_UPDATED_TOPIC || *queue == "*" {
		onercc, err := broker.AddEventListener(event.SERAT_UPDATED_TOPIC, listener.ECHO_CHANNEL, seratListener.UpdatedEventListener)

		if err != nil {
			log.Fatalf("cannot listen to user consumer: %v", err)
		}

		consumers = append(consumers, onercc)
	}

	if *queue == event.SERAT_DELETED_TOPIC || *queue == "*" {
		onercc, err := broker.AddEventListener(event.SERAT_DELETED_TOPIC, listener.ECHO_CHANNEL, seratListener.DeletedEventListener)

		if err != nil {
			log.Fatalf("cannot listen to user consumer: %v", err)
		}

		consumers = append(consumers, onercc)
	}

	for _, consumer := range consumers {
		consumer.ConnectToNSQD(broker.Address)
		consumer.ConnectToNSQLookupd(broker.LookupAddress)
	}

	// wait for signal to exit
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	<-sigChan

	for _, consumer := range consumers {
		consumer.Stop()
	}
}
