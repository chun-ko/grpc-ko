package main

import (
	"context"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "redbelly.grpcrace/proto"
	"redbelly.grpcrace/race"
)

func startServer(listenOn string) {
	lis, err := net.Listen("tcp", listenOn)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	pb.RegisterChecksumServiceServer(s, &race.Server{})

	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func method1(total int, totalCount chan int) {
	log.Print("Creating Client connection")
	conn, err := grpc.Dial("localhost:50501", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewChecksumServiceClient(conn)

	log.Print("Connection Created")

	log.Print("Start Checking")
	count := 0
	for i := 0; i < total; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		_, err := c.Checksum(ctx, &pb.CheckRequest{Payload: "abc"})
		if err != nil {
			log.Fatalf("could not greet: %d %v", i, err)
		}
		//log.Printf("Greeting: %d", r.GetSize())
		count++
	}
	log.Print("End Checking")
	totalCount <- count
}

func method2(total int) {
	totalCount := make(chan int)
	go method1(total/2, totalCount)
	go method1(total/2, totalCount)

	sum := 0
	log.Print("Waiting all to be finished")
	for sum < total {
		result := <-totalCount
		sum = sum + result
	}
	log.Print("All Finished")

}

func main() {

	log.Print("Starting Server.")
	go startServer("localhost:50501")
	time.Sleep(2 * time.Second)
	log.Print("Server started.")

	total := 1000000

	method2(total)

}
