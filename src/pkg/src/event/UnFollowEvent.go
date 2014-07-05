// This is the Unfollow event. 
// We need to know which client is "unfollowing" which client
package event

// This is the class that is used to represent the unfollow event.
// fromClient - this is the client that wants to stop following.
// toClient - this is the client that was being followed but will not be anymore.
type UnFollowEvent struct {
	BasicEvent
	fromClient int
	toClient int
}

// this is the constructor for a new unfollow event
func NewUnFollowEvent(sequence int, payLoad string, fromClient int, toClient int) *UnFollowEvent {
	event := new(UnFollowEvent)
	event.sequence = sequence
	event.eventType = UNFOLLOW
	event.payLoad = payLoad	
	event.fromClient = fromClient
	event.toClient = toClient
	return event
}

func (event UnFollowEvent) GetFromClient() int {
    return event.fromClient
}

func (event UnFollowEvent) GetToClient() int {
	return event.toClient
}


