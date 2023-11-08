package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"

	protocol "github.com/Casper2411/DISYS-HandIn4-Ben_Dover/grpc"
	"google.golang.org/grpc"
)

func main() {
	arg1, _ := strconv.ParseInt(os.Args[1], 10, 32)
	ownPort := int32(arg1) + 5000

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	p := &peer{
		id:            ownPort,
		amountOfPings: make(map[int32]int32),
		clients:       make(map[int32]protocol.RicartAgrawalaServiceClient),
		ctx:           ctx,
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

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		p.SendRequests()
	}
}

type peer struct {
	protocol.UnimplementedRicartAgrawalaServiceServer
	id            int32
	amountOfPings map[int32]int32
	clients       map[int32]protocol.RicartAgrawalaServiceClient
	ctx           context.Context
}

func (p *peer) RicartAgrawala(ctx context.Context, req *protocol.Request) (*protocol.Reply, error) {
	id := req.Id
	p.amountOfPings[id] += 1

	rep := &protocol.Reply{ Message: "OK"}
	return rep, nil
}

func (p *peer) SendRequests() {
	request := &protocol.Request{Id: p.id}
	for id, client := range p.clients {
		reply, err := client.RicartAgrawala(p.ctx, request)
		if err != nil {
			fmt.Println("something went wrong")
		}
		fmt.Printf("Got reply from id %v: %v\n", id, reply.Message)
	}
}
