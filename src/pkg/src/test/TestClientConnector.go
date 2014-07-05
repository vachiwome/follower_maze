// this is used to test the client connector
package test

import (
	"net"
	"sync"
	"bufio"
	"strconv"
	"fmt"
	"../worker"
)

type TestClientConnector struct {
	AbstractTest
}

// This is the method for simulating a client that is trying to connect to the server
func (tester TestClientConnector) dial(clientId int, port string, waitGroup *sync.WaitGroup) {
	conn, _ := net.Dial("tcp", "localhost:"+port)
	writer := bufio.NewWriter(conn)
	id := strconv.Itoa(clientId) + "\n"
    writer.WriteString(id)
    writer.Flush()
    waitGroup.Done()
}

// This is the method that simulates a server that is listening for connection requests
func (tester TestClientConnector) accept(clientId int, port string, waitGroup *sync.WaitGroup) {
	fmt.Println("Running test #1")
	clientConn := worker.NewClientConnector(":"+port)
	client := clientConn.NextClient()
	if clientId == client.GetId() {
		fmt.Println("TestClientConnector : test passed")
	} else {
		fmt.Println("TestClientConnector : test failed")
	}
	waitGroup.Done()
} 

func (tester TestClientConnector) Run() {
	var waitGroup sync.WaitGroup
	waitGroup.Add(2)
	go tester.accept(1, "9099", &waitGroup)
	go tester.dial(1, "9099", &waitGroup)
	waitGroup.Wait()
}