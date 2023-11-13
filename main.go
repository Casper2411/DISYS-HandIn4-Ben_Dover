package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"strconv"
	"time"

	protocol "github.com/Casper2411/DISYS-HandIn4-Ben_Dover/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

//Peer struct used to represent a node in the system
type peer struct {
	protocol.UnimplementedRicartAgrawalaServiceServer
	id                  int32
	//A map of all the other peers in the system	
	clients             map[int32]protocol.RicartAgrawalaServiceClient
	ctx                 context.Context
	isRequesting        bool
	ownRequest          *protocol.Request
	isInCriticalSection bool
	lamportTimestamp    int32
}

func main() {
	//Parse the port number from the command line
	arg1, _ := strconv.ParseInt(os.Args[1], 10, 32)
	ownPort := int32(arg1) + 5000

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	//Create a new peer/node with the given port number
	p := &peer{
		id:                  ownPort,
		clients:             make(map[int32]protocol.RicartAgrawalaServiceClient),
		ctx:                 ctx,
		isRequesting:        false,
		isInCriticalSection: false,
		lamportTimestamp:    1,
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
	
	//Look for the other peers/nodes in the system and connect to them
	for i := 0; i < 3; i++ {
		port := int32(5000) + int32(i)

		if port == ownPort {
			continue
		}

		var conn *grpc.ClientConn
		fmt.Printf("Trying to dial: %v\n", port)
		conn, err := grpc.Dial(fmt.Sprintf(":%v", port), grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
		if err != nil {
			log.Fatalf("Could not connect: %s", err)
		}
		defer conn.Close()
		c := protocol.NewRicartAgrawalaServiceClient(conn)
		//Add the peer to the map of peers
		p.clients[port] = c
	}
	
	//Keep generating random numbers to decide if we should send a request to enter critical section
	for {
		rand.NewSource(time.Now().UnixNano())
		randomInt := rand.Intn(10)
		if randomInt < 2 {
			//Send request to all other peers
			p.SendRequests()

			//Enter critical section
			p.isInCriticalSection = true
			fmt.Printf("Peer %d entering critical section. Lamport: %v\n", p.id, p.lamportTimestamp)
			
			//Wait (simulates doing work in critical section)
			time.Sleep(3 * time.Second)
			
			//Exit critical section
			fmt.Printf("Peer %d exiting critical section.\n", p.id)
			p.isInCriticalSection = false
			p.isRequesting = false
		}
	}

}

func (p *peer) RicartAgrawala(ctx context.Context, req *protocol.Request) (*protocol.Reply, error) {
	p.lamportTimestamp += 1
	fmt.Printf("Peer %d received request: {Peer %d, Lamport time %v}\n", p.id, req.Id, req.LamportTimestamp)
	
	for !p.shouldIReply(req) {
		//Waiting until we can reply (which is when shouldIReply() returns true)
	}

	fmt.Printf("Peer %d replied to request: {Peer %d, Lamport time %v}\n", p.id, req.Id, req.LamportTimestamp)
	rep := &protocol.Reply{Message: "OK"}
	return rep, nil
}

func (p *peer) SendRequests() {
	//Create a new request with own id and current lamport timestamp
	request := &protocol.Request{Id: p.id, LamportTimestamp: p.lamportTimestamp}
	fmt.Printf("Peer %v is sending requests: {Peer %d, Lamport time %v} \n", p.id, p.id, p.lamportTimestamp)
	
	//Set the peer to be requesting and set the "ownRequest" to the current request
	p.isRequesting = true
	p.ownRequest = request
	
	//Send the request to all other peers
	for id, client := range p.clients {
		//Wait for reply from peer
		reply, err := client.RicartAgrawala(p.ctx, request)
		if err != nil {
			fmt.Println("something went wrong")
		}
		fmt.Printf("Got reply from id %v: %v to request: {Peer %d, Lamport time %v}\n", id, reply.Message, p.ownRequest.Id, p.ownRequest.LamportTimestamp)
	}
	p.lamportTimestamp += 1
}

//A helper function to decide if peer should reply to a request
func (p *peer) shouldIReply(req *protocol.Request) bool {
	if p.isInCriticalSection {
		return false
	}
	if p.isRequesting {
		if p.ownRequest.LamportTimestamp < req.LamportTimestamp {
			return false
		} else if p.ownRequest.LamportTimestamp == req.LamportTimestamp {
			if p.ownRequest.Id < req.Id {
				return false
			}
		}
	}
	return true
}
