// This class is used to check that the event scheduler delivers the events
// in the order that the client expects.
package test

import (
	"net"
	"sync"
	"bufio"
	"fmt"
	"strconv"
	"../worker"
)

type TestEventScheduler struct {
	AbstractTest
}

// This method is used to simulate an event source that produces events in the opposite sequence number order
// These events have to be reordered by the scheduler.
func (tester TestEventScheduler) sendEvents(numEvents int, port string, waitGroup *sync.WaitGroup) {
	conn, _ := net.Dial("tcp", "localhost:"+port)
	writer := bufio.NewWriter(conn)
	for i:=numEvents; i >=1; i-- {
 		writer.WriteString(strconv.Itoa(i)+"|B\n")
    	writer.Flush()
	}
    waitGroup.Done()
}

// This method simulates a client to the event scheduler that expects events
// in increasing sequence number order
func (tester TestEventScheduler) receiveEvents(testId int, numEvents int, port string, waitGroup *sync.WaitGroup) {
	fmt.Println("Running test #" + strconv.Itoa(testId))
	eventGen := worker.NewEventGenerator(":"+port)
	eventScheduler := worker.NewEventScheduler(eventGen, 10000000)
	error := false
	for i:=1; i < numEvents; i++ {
		e, _ := eventScheduler.NextEvent(i)
		if e.GetSequence() != i {
			fmt.Println("TestEventScheduler : unexpected sequence number received")
			error = true
		}

	}
	if !error {
		fmt.Println("TestEventScheduler : test passed")
	}
	waitGroup.Done()
} 

func (tester TestEventScheduler) Run() {
	var waitGroup sync.WaitGroup
	waitGroup.Add(2)
	go tester.receiveEvents(7, 10, "9095", &waitGroup)
	go tester.sendEvents(10, "9095", &waitGroup)
	waitGroup.Wait()
}




