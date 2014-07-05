// this is used to test the client connector
package test

import (
	"net"
	"sync"
	"bufio"
	"strconv"
	"fmt"
	"../client"
	"../worker"
	"../event"
)

type TestEventProcessor struct {
	AbstractTest
}

func (tester TestEventProcessor) dial(clientId int, payLoad string, port string, waitGroup *sync.WaitGroup) bool {
	conn, _ := net.Dial("tcp", "localhost:"+port)
	writer := bufio.NewWriter(conn)
	id := strconv.Itoa(clientId) + "\n"
    writer.WriteString(id)
    writer.Flush()

    reader := bufio.NewReader(conn)
    p, _ := reader.ReadString('\n')
    error := false
    payLoad = payLoad + "\n"
    if payLoad != p {
    	fmt.Println("Expected : " + payLoad + " but received : " + p)
    	error = true
    }
    waitGroup.Done()
    return error
}

func (tester TestEventProcessor) accept(clientId int, port string, waitGroup *sync.WaitGroup) bool {
	clientConn := worker.NewClientConnector(":"+port)
	clnt := clientConn.NextClient()

	clients := make(map[int]client.Client)
	clients[clnt.GetId()] = clnt

	eventProc := new(worker.EventProcessor)
	e := event.NewFollowEvent(1,"1|F|"+strconv.Itoa(clientId)+"|"+strconv.Itoa(clientId)+"\n", clientId, clientId)
	clients = eventProc.ProcessEvent(e, clients)

	error := false
	if len(clients[clientId].GetFollowers()) != 1 {
		fmt.Println("Follower was not successfully added")
		error = true
	}
	waitGroup.Done()
	return error
} 

func (tester TestEventProcessor) Run() {
	var waitGroup sync.WaitGroup
	waitGroup.Add(2)
	go tester.accept(1, "9099", &waitGroup)
	go tester.dial(1, "1|F|1|1", "9099", &waitGroup)
	waitGroup.Wait()
}