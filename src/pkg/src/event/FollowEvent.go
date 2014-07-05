// This is a follow event. For this event, we need to know 
// who wants to follow whom

package event

// This is the event of type FOLLOW.
// fromClient- this is the client that wants to follow.
// toClient - this is the client that is being followed.
type FollowEvent struct {
	BasicEvent
	fromClient int
	toClient int
}

func NewFollowEvent(sequence int, payLoad string, fromClient int, toClient int) *FollowEvent {
	event := new(FollowEvent)
	event.sequence = sequence
	event.eventType = FOLLOW
	event.payLoad = payLoad	
	event.fromClient = fromClient
	event.toClient = toClient
	return event
}

func (event FollowEvent) GetFromClient() int {
    return event.fromClient
}

func (event FollowEvent) GetToClient() int {
	return event.toClient
}


