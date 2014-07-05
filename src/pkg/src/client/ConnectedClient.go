// This is a client which the server can communicate with.
// The server has a socket associated with this client
package client

import (
	"net"
	"bufio"
)

type ConnectedClient struct {
	BasicClient
	writer *bufio.Writer
}

// A client that is created when the server accepts the client's connection request
func NewConnectedClient(id int, connection net.Conn) *ConnectedClient {
	client := new(ConnectedClient)
	client.id = id
	client.isConnected = true
	followers := make(map[int]bool)
	client.followers = followers
	writer := bufio.NewWriter(connection)
	client.writer = writer
	return client
}

// Write the data to the client via the socket
func (client ConnectedClient) SendEvent(payLoad string) {
	client.writer.WriteString(payLoad)
	client.writer.Flush()
}
