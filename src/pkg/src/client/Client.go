// This is the interface of all types of clients that associate with the server.
// In this package, a client can either have a socket connection associated with it or not.
package client

type Client interface {
	AddFollower(follower int)
	RemoveFollower(follower int) bool
	SendEvent(payLoad string)
	GetIsConnected() bool
	GetId() int
	GetFollowers() map[int]bool
}