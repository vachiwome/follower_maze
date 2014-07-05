// This is Broadcast event
package event

type BroadcastEvent struct {
	BasicEvent
}

// Constructor for a new broadcast event
func NewBroadcastEvent(sequence int, payLoad string) *BroadcastEvent {
	event := new(BroadcastEvent)
	event.sequence = sequence
	event.eventType = BROADCAST
	event.payLoad = payLoad	
	return event
}

