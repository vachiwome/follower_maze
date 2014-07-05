// This is the central server.
// It runs two threads, one that is accepting new clients and another one that is
// receiving and processing events.

// to run the server, do "go run Server.go"
package main

import (
	"sync"
	"fmt"
	"strconv"
	"../worker"
	"../client"
)

// This is the class that is used to represent a server.
// mutex - this is the mutex that is used to enforce mutual exclusion between the two threads that run in the server process.
// clients - this is the map of the available clients. A map was used, rather than a list or other data structures to enforce constant time access.
// clientConnector - this is the worker whose job is to accept connection requests from user clients.
// eventGenerator - this is the worker that listens for events and creates Event objects.
// eventScheduler - this is used to implement any scheduling algorithm for the order of the delivery of the event.
type Server struct {
	mutex  sync.Mutex
	clients map[int]client.Client
	clientConnector *worker.ClientConnector
	eventGenerator *worker.EventGenerator
	eventScheduler *worker.EventScheduler
	eventProcessor *worker.EventProcessor
}

// this is the constructor for a server object.
func NewServer(clientPort string, eventSource string) *Server {
	server := new(Server)
	clientConnector := worker.NewClientConnector(clientPort)
	server.clientConnector = clientConnector

	eventGenerator := worker.NewEventGenerator(eventSource)
	server.eventGenerator = eventGenerator
	
	eventScheduler := worker.NewEventScheduler(eventGenerator, 10000000)
	server.eventScheduler = eventScheduler
	
	eventProcessor := worker.NewEventProcessor()
	server.eventProcessor = eventProcessor
	
	clients := make(map[int]client.Client)
	server.clients = clients
	
	return server
}

// this function runs in its own thread, accepting clients as they come.
func (server Server) clientConnectorThread(waitGroup sync.WaitGroup) {
	fmt.Println("client connector started")
	for {
		client := server.clientConnector.NextClient()
		server.mutex.Lock()
		server.clients[client.GetId()] = client
		server.mutex.Unlock()
	}
	waitGroup.Done()
	fmt.Println("client connector terminated")
}

// This function runs in its own thread, accepting events and scheduling them and delivering them in the right order
func (server Server) eventListenerThread(waitGroup sync.WaitGroup) {
	fmt.Println("event listerner started")
	expected := 1
	for {
		event, isExpected := server.eventScheduler.NextEvent(expected)
		if event != nil {
			if isExpected {
				expected = expected + 1
			}
			server.mutex.Lock()
			newClients := server.eventProcessor.ProcessEvent(event, server.clients)
			server.clients = newClients
			server.mutex.Unlock()
			if event.GetSequence() == event.GetSequence() {
				fmt.Println("Received event : " + event.GetPayLoad())
			}
		}
	}
	waitGroup.Done()
	fmt.Println("event listener terminated")
}

func main() {
	server := NewServer(":9099", ":9090")
	var waitGroup sync.WaitGroup
	waitGroup.Add(2)
	go server.clientConnectorThread(waitGroup)
	go server.eventListenerThread(waitGroup)
	waitGroup.Wait()	
}
