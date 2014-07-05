// Interface for all the events that the server can recieve
package event

// Integer codes for the different possible events.
const (
	FOLLOW = 0
	UNFOLLOW = 1
	PRIVATE_MSG = 2
	STATUS_UPDATE = 3
	BROADCAST = 4
)

type Event interface {
	GetSequence () int
	GetEventType () int
	GetPayLoad () string
	GetToClient () int
	GetFromClient() int
}
