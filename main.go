package main

import (
	//"bufio"
	"context"
	//"crypto/rand"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"math/rand"
	"time"
	protocol "github.com/Casper2411/DISYS-HandIn4-Ben_Dover/grpc"
	"google.golang.org/grpc"
)

type peer struct {
	protocol.UnimplementedRicartAgrawalaServiceServer
	id            int32
	//amountOfPings map[int32]int32
	clients       map[int32]protocol.RicartAgrawalaServiceClient
	ctx           context.Context
	//maintain a list of requests
	//requests	  []*protocol.Request
	isRequesting  bool
	ownRequest    *protocol.Request
	isInCriticalSection bool
	lamportTimestamp int32
}

func main() {
	arg1, _ := strconv.ParseInt(os.Args[1], 10, 32)
	ownPort := int32(arg1) + 5000

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	//make a slice of requests
	//var requestSlice []*protocol.Request

	p := &peer{
		id:            ownPort,
		//amountOfPings: make(map[int32]int32),
		clients:       make(map[int32]protocol.RicartAgrawalaServiceClient),
		ctx:           ctx,
		//requests: 	   requestSlice,
		isRequesting: false,
		isInCriticalSection: false,
		lamportTimestamp: 1,
	}

	// Create listener tcp on port ownPort
	list, err := net.Listen("tcp", fmt.Sprintf(":%v", ownPort))
	if err != nil {
		log.Fatalf("Failed to listen on port: %v", err)
	}
	grpcServer := grpc.NewServer()
	protocol.RegisterRicartAgrawalaServiceServer(grpcServer, p)

	go func() {
		if err := grpcServer.Serve(list); err != nil {
			log.Fatalf("failed to server %v", err)
		}
	}()

	for i := 0; i < 3; i++ {
		port := int32(5000) + int32(i)

		if port == ownPort {
			continue
		}

		var conn *grpc.ClientConn
		fmt.Printf("Trying to dial: %v\n", port)
		conn, err := grpc.Dial(fmt.Sprintf(":%v", port), grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			log.Fatalf("Could not connect: %s", err)
		}
		defer conn.Close()
		c := protocol.NewRicartAgrawalaServiceClient(conn)
		p.clients[port] = c
	}

	/*scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		p.SendRequests()
	}*/

	for{
		rand.NewSource(time.Now().UnixNano())
		randomInt := rand.Intn(10)
		//fmt.Println(randomInt)
		if(randomInt<2){
			//send request
			p.SendRequests()
			//enter critical section
			p.isInCriticalSection = true
			fmt.Printf("Client %d entering critical section. It's a critical hit! Lamport: %v\n", p.id, p.lamportTimestamp)
			//wait
			time.Sleep(3 * time.Second)
			//exit critical section
			fmt.Printf("Client %d exiting critical section.\n", p.id)
			p.isInCriticalSection = false
		}
	}
	
}


func (p *peer) RicartAgrawala(ctx context.Context, req *protocol.Request) (*protocol.Reply, error) {
	p.lamportTimestamp += 1
	fmt.Printf("Client %d received request from Client %d with Lamport time %v\n", p.id, req.Id, req.LamportTimestamp)
	for !p.shouldIReply(req){
		//wait
	}
	fmt.Printf("Client %d replied to Client %d\n", p.id, req.Id)
	rep := &protocol.Reply{ Message: "OK"}
	return rep, nil
}

func (p *peer) SendRequests() {
	request := &protocol.Request{Id: p.id, LamportTimestamp: p.lamportTimestamp}
	fmt.Printf("Client %v is sending requests at Lamport time %v\n", p.id, p.lamportTimestamp)
	p.isRequesting = true
	p.ownRequest = request
	for id, client := range p.clients {
		reply, err := client.RicartAgrawala(p.ctx, request)
		if err != nil {
			fmt.Println("something went wrong")
		}
		fmt.Printf("Got reply from id %v: %v\n", id, reply.Message)
	}
	//check that all replies have been received
	//p is no longer requesting after receiving all replies and entering critical section
	p.isRequesting = false
	p.lamportTimestamp += 1	
}

func (p *peer) shouldIReply(req *protocol.Request) bool{
	if(p.isInCriticalSection){
		return false
	}
	if(p.isRequesting){
		if(p.ownRequest.LamportTimestamp<req.LamportTimestamp){
			return false
		}else if(p.ownRequest.LamportTimestamp==req.LamportTimestamp){
			if(p.ownRequest.Id<req.Id){
				return false
			}
		}
	} 
	
	return true
}