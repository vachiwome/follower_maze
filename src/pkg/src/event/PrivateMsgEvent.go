// This is a private message event.
// For this event, we need to know which client wants to send which client a message
package event

// This is the class that encodes a Private Message instance.
// fromClient - this is the client that is sending the private message
// toClient - this is the client to which the private message is being sent
type PrivateMsgEvent struct {
	BasicEvent
	fromClient int
	toClient int
}

func NewPrivateMsgEvent(sequence int, payLoad string, fromClient int, toClient int) *PrivateMsgEvent {
	event := new(PrivateMsgEvent)
	event.sequence = sequence
	event.eventType = PRIVATE_MSG
	event.payLoad = payLoad	
	event.fromClient = fromClient
	event.toClient = toClient
	return event
}

func (event PrivateMsgEvent) GetFromClient() int {
    return event.fromClient
}

func (event PrivateMsgEvent) GetToClient() int {
	return event.toClient
}




