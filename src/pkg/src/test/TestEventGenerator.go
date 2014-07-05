//This is used to test that the event generator reads the strings that encode
// events correctly and creates the correct Event objects.
package test

import (
	"net"
	"sync"
	"bufio"
	"fmt"
	"strconv"
	"../event"
	"../worker"
)

type TestEventGenerator struct {
	AbstractTest
}

// This is the method that is used to simulate the event source
func (tester TestEventGenerator) sendEvent(payLoad string, port string, waitGroup *sync.WaitGroup) {
	conn, _ := net.Dial("tcp", "localhost:"+port)
	writer := bufio.NewWriter(conn)
    writer.WriteString(payLoad)
    writer.Flush()
    waitGroup.Done()
}

// This is the method that simulates the EventGenerator
func (tester TestEventGenerator) receiveEvent(testId int, sequence int, eventType int, fromClient int, toClient int, port string, waitGroup *sync.WaitGroup) {
	fmt.Println("Running test #" + strconv.Itoa(testId))
	eventGen := worker.NewEventGenerator(":"+port)
	e := eventGen.NextEvent()
	error := false
	if sequence != e.GetSequence() {
		fmt.Println("TestEventGenerator : sequence numbers do not match")
		error = true
	} 
	if eventType != e.GetEventType() {
		fmt.Println("TestEventGenerator : event types do not match")
		error = true
	}
	if fromClient != -1 && e.GetFromClient() != fromClient {
		fmt.Println("TestEventGenerator : fromClient does not match")
		error = true
	}
	if toClient != -1 && e.GetToClient() != toClient {
		fmt.Println("TestEventGenerator : toClient does not match")
		error = true
	}
	if !error {
		fmt.Println("TestEventGenerator test passed")
	}
	waitGroup.Done()
} 

// test various events
func (tester TestEventGenerator) Run() {
	var waitGroup sync.WaitGroup
	waitGroup.Add(2)
	go tester.receiveEvent(2, 1, event.BROADCAST, -1, -1, "9090", &waitGroup)
	go tester.sendEvent("1|B\n", "9090", &waitGroup)
	waitGroup.Wait()
	
	waitGroup.Add(2)
	go tester.receiveEvent(3, 1, event.FOLLOW, 12, 13, "9091", &waitGroup)
	go tester.sendEvent("1|F|12|13\n", "9091", &waitGroup)
	waitGroup.Wait()

	waitGroup.Add(2)
	go tester.receiveEvent(4, 1, event.UNFOLLOW, 12, 13, "9092", &waitGroup)
	go tester.sendEvent("1|U|12|13\n", "9092", &waitGroup)
	waitGroup.Wait()

	waitGroup.Add(2)
	go tester.receiveEvent(5, 1, event.PRIVATE_MSG, 12, 13, "9093", &waitGroup)
	go tester.sendEvent("1|P|12|13\n", "9093", &waitGroup)
	waitGroup.Wait()

	waitGroup.Add(2)
	go tester.receiveEvent(6, 1, event.STATUS_UPDATE, 12, -1, "9094", &waitGroup)
	go tester.sendEvent("1|S|12|13\n", "9094", &waitGroup)
	waitGroup.Wait()
}




