// This worker is responsible for receiving events from the 
// events source and passing them to the master program.
package worker

import (
        "net"
        "bufio"
        "strings"
        "strconv"
        "fmt"
        "../event"
)

// connector - this is the instance that connects this class to the sockets that produces the stream of events
// reader - reader that is used for reading the strings that encode different events
type EventGenerator struct {
        connector net.Conn
        reader *bufio.Reader
}

// This is the constructor for a new event generator
func NewEventGenerator(eventSource string) *EventGenerator {
        eventGen := new(EventGenerator)
        listener, err := net.Listen("tcp", eventSource)
        if err == nil {
        	connector, _ := listener.Accept()
        	eventGen.connector = connector
        	reader := bufio.NewReader(connector)
        	eventGen.reader = reader
        	return eventGen
        } else {
        	fmt.Println("connection to event source could not be established")
        	return nil
        }
} 

// create an event object using a string encoding of that event
func (eventGen EventGenerator) createEvent(payLoad string) event.Event {
	tokens := strings.Split(payLoad[:len(payLoad)-1], "|")
	if false {
		fmt.Println(tokens)
	}
	sequence, _ := strconv.Atoi(tokens[0])
	switch {
		case tokens[1] == "F":
			fromClient, _ := strconv.Atoi(tokens[2])
			toClient, _ := strconv.Atoi(tokens[3])
			return event.NewFollowEvent(sequence, payLoad, fromClient, toClient)
		case tokens[1] == "U":
			fromClient, _ := strconv.Atoi(tokens[2])
			toClient, _ := strconv.Atoi(tokens[3])
			return event.NewUnFollowEvent(sequence, payLoad, fromClient, toClient)
		case tokens[1] == "B":
			return event.NewBroadcastEvent(sequence, payLoad)
		case tokens[1] == "P":
			fromClient, _ := strconv.Atoi(tokens[2])
			toClient, _ := strconv.Atoi(tokens[3])
			return event.NewPrivateMsgEvent(sequence, payLoad, fromClient, toClient)
		case tokens[1] == "S":
			fromClient, _ := strconv.Atoi(tokens[2])
			return event.NewStatusUpdateEvent(sequence, payLoad, fromClient)
	}
	return nil
}

// Method for getting the next available event
func (eventGen EventGenerator) NextEvent() event.Event {
		line, _ := eventGen.reader.ReadString('\n')
		if len(line) > 0 {
			return eventGen.createEvent(line)
		}
        return nil
}
