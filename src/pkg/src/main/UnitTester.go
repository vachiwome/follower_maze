// This program is responsoble for performing unit tests for the different workers in this project
package main

import (
	"../test"
)

func main() {
	// test whether the client connector works as expected
	testClientConn := new(test.TestClientConnector)
	testClientConn.Run()

	// check if the event generator parses the strings well and create correct event objects
	testEventGen := new(test.TestEventGenerator)
	testEventGen.Run()

	// check if the event scheduler delivers events in the right sequence number order
	testEventScheduler := new(test.TestEventScheduler)
	testEventScheduler.Run()

	// check if events are processed the right way
	testEventProcessor := new(test.TestEventProcessor)
	testEventProcessor.Run()
}