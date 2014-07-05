// This worker is responsible for processing the events
// that are received, sending them to the appropriate users
// and modifying the appropriate data structures
package worker

import (
		"fmt"
        "../client"
        "../event"
)

type EventProcessor struct {}

func NewEventProcessor() *EventProcessor {
        return new(EventProcessor)
} 

// This is the method that processes events
func (eventProc EventProcessor) ProcessEvent(e event.Event, clients map[int]client.Client) (map[int]client.Client) {
	eventType := e.GetEventType()
	switch {
		case eventType == event.FOLLOW:
			return eventProc.processFollowEvent(e, clients)
		case eventType == event.UNFOLLOW:
			return eventProc.processUnFollowEvent(e, clients)
		case eventType == event.BROADCAST:
			return eventProc.processBroadcastEvent(e, clients)
		case eventType == event.PRIVATE_MSG:
			return eventProc.processPrivateMsgEvent(e, clients)
		case eventType == event.STATUS_UPDATE:
			return eventProc.processStatusUpdateEvent(e, clients)
		default:
			fmt.Println("Unknown event type received " + e.GetPayLoad())
			return clients
	}
}

// Process a follow event. Add the follower to the appropriate client.
// Add the client if it is not already added.
func (eventProc EventProcessor) processFollowEvent(e event.Event, clients map[int]client.Client) (map[int]client.Client){
	toclnt, exists := clients[e.GetToClient()]
	if !exists {
		clients[e.GetToClient()] = client.NewBasicClient(e.GetToClient())
		toclnt, _ = clients[e.GetToClient()]
	} 

	_, exists = clients[e.GetFromClient()]
	if !exists {
		clients[e.GetFromClient()] = client.NewBasicClient(e.GetFromClient())
	}

	toclnt.AddFollower(e.GetFromClient())
	toclnt.SendEvent(e.GetPayLoad())
	return clients 
}

// Process the unfollow event. Remove the follower from the appropriate client.
func (eventProc EventProcessor) processUnFollowEvent(e event.Event, clients map[int]client.Client) (map[int]client.Client) {
	toclnt, exists := clients[e.GetToClient()]
	if !exists {
		clients[e.GetToClient()] = client.NewBasicClient(e.GetToClient())
		toclnt, _ = clients[e.GetToClient()]
	} 

	_, exists = clients[e.GetFromClient()]
	if !exists {
		clients[e.GetFromClient()] = client.NewBasicClient(e.GetFromClient())
	}

	toclnt.RemoveFollower(e.GetFromClient())
	return clients
}

// Method for processing a broadcast message. Send the message to all the clients.
func (eventProc EventProcessor) processBroadcastEvent(e event.Event, clients map[int]client.Client) (map[int]client.Client) {
	for _, clientObj := range clients {
    	clientObj.SendEvent(e.GetPayLoad())
	}
	return clients
}

// Process a private message event. Send the message to the right client.
func (eventProc EventProcessor) processPrivateMsgEvent(e event.Event, clients map[int]client.Client) (map[int]client.Client) {
	toclnt, exists := clients[e.GetToClient()]
	if !exists {
		clients[e.GetToClient()] = client.NewBasicClient(e.GetToClient())
		toclnt, _ = clients[e.GetToClient()]
	} 

	_, exists = clients[e.GetFromClient()]
	if !exists {
		clients[e.GetFromClient()] = client.NewBasicClient(e.GetFromClient())
	}

	toclnt.SendEvent(e.GetPayLoad()) 
	return clients
}

// Process a status update event. Send this message to all the followers.
func (eventProc EventProcessor) processStatusUpdateEvent(e event.Event, clients map[int]client.Client) (map[int]client.Client) {
	fromclnt, exists := clients[e.GetFromClient()]
	if !exists {
		clients[e.GetFromClient()] = client.NewBasicClient(e.GetFromClient())
		fromclnt, _ = clients[e.GetFromClient()]
	} 

	for clientId, _ := range fromclnt.GetFollowers() {
    	follower, _ := clients[clientId]
    	follower.SendEvent(e.GetPayLoad())
	}
	return clients
}
