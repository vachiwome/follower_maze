// This thread is responsible for receiving client
// connection requests
package worker

import (
        "net"
        "bufio"
        "fmt"
        "strconv"
        "../client"
)

// listener - specifies an object that waits for connection requests
type ClientConnector struct {
        listener net.Listener
}

// This is the constructor for a new client connector
func NewClientConnector(port string) *ClientConnector {
        cc := new(ClientConnector)
        listener, err := net.Listen("tcp", port)
        if err != nil {
                fmt.Println("client connector error")
        }
        cc.listener = listener
        return cc
} 

// This is the method for getting the next client that has requested to connect
func (clientConnector ClientConnector) NextClient() *client.ConnectedClient {
        conn, err := clientConnector.listener.Accept()
        if err == nil {
                reader := bufio.NewReader(conn)
                line, _ := reader.ReadString('\n')
                clientId, _ := strconv.Atoi(line[:len(line)-1])
                return client.NewConnectedClient(clientId, conn)
        } else {
                fmt.Println("could not accept a new client")
        }

        return nil
}
