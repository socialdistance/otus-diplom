package main

import (
	"context"
	"flag"
	"google.golang.org/grpc"
	"io"
	"log"
	internalgrpc "static_collector/internal/server/grpc"
)

var port string
var n, m int

func init() {
	flag.StringVar(&port, "port", ":50051", "Listen port")
	flag.IntVar(&n, "n", 5, "interval to get statistic (sec)")
	flag.IntVar(&m, "m", 15, "interval to average statistic (sec)")
}

func main() {
	flag.Parse()

	conn, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("can't connect with server %v", err)
	}

	client := internalgrpc.NewStreamServiceClient(conn)
	in := internalgrpc.GatherRequest{
		N: int64(n),
		M: int64(m),
	}

	stream, err := client.ListGather(context.Background(), &in)
	if err != nil {
		log.Fatalf("open stream error: %s", err)
	}

	done := make(chan bool)

	go func() {
		resp, err := stream.Recv()
		if err == io.EOF {
			done <- true
			return
		}
		if err != nil {
			log.Fatalf("cannot receive %v", err)
		}
		log.Printf("Resp received: %s", resp.Result)
	}()

	<-done
	log.Printf("finished")

}
