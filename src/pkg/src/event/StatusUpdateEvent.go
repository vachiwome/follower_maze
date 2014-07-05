// This is a status update event
// For this event we need to know which client is sending the status updates. 
// so that we can know which clients follow the former.
package event

// This is the class that represents the status update events.
// fromClient - this is the client that is sending the status update events to its clients.
type StatusUpdateEvent struct {
	BasicEvent
	fromClient int
}

// This is the constructor for the status update events
func NewStatusUpdateEvent(sequence int, payLoad string, fromClient int) *StatusUpdateEvent {
	event := new(StatusUpdateEvent)
	event.sequence = sequence
	event.eventType = STATUS_UPDATE
	event.payLoad = payLoad	
	event.fromClient = fromClient
	return event
}

func (event StatusUpdateEvent) GetFromClient() int {
    return event.fromClient
}
