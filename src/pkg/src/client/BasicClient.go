// This is a basic client that has an identifier and followers
// This client may or may not be registered via a socket
package client

// id - unique integer identifier for a client. 
// isConnected - true if the client has an associated socket connection
// followers - map that contains the identifiers of the clients that follow this client. A map was used for constant access.
type BasicClient struct {
	Client
	id int
	isConnected bool
	followers map[int]bool
}

// Constructor for a new un-connected client
func NewBasicClient(id int) *BasicClient {
	client := new(BasicClient)
	client.id = id
	client.isConnected = false
	followers := make(map[int]bool)
	client.followers = followers
	return client
}

// register a new follower with identifier "id". It is called when a follow event is being processed.
func (client BasicClient) AddFollower(follower int ) {
	client.followers[follower] = true
}

// Remove s follower that has identifier "id" from the this client's followers.
// Returns true if this client was being followed by 'follower' else returns false.
func (client BasicClient) RemoveFollower(follower int) bool {
	_, exists := client.followers[follower]
	if exists {
		delete(client.followers, follower)
	}
	return exists
}

// a stub for sending an event to a client that does not have an associated socket connection
func (client BasicClient) SendEvent(payLoad string) {
	// do nothing
}

func (client BasicClient) GetIsConnected() bool {
	return client.isConnected 
}

func (client BasicClient) GetId() int {
	return client.id
}

func (client BasicClient) GetFollowers() map[int]bool {
	return client.followers
}