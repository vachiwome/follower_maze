// This is a basic event that implements the event interface.
// This event only has a sequence number and a payload. It is also labelled with its type
package event

// This is the constructor for a basic event. 
// sequence - the sequence number for this event.
// eventType - this represents the type of event in question (refer to file Event.go to see the list of possible values). 
type BasicEvent struct {
	Event
	sequence int
	eventType int
	payLoad string	
}

// constructor for a new basic event.
func NewBasicEvent(sequence int, eventType int, payLoad string) *BasicEvent {
	event := new(BasicEvent)
	event.sequence = sequence
	event.eventType = eventType
	event.payLoad = payLoad	
	return event
}

func (event BasicEvent) GetSequence() int {
    return event.sequence
}

func (event BasicEvent) GetEventType() int {
	return event.eventType
}

func (event BasicEvent) GetPayLoad() string {
	return event.payLoad
}


