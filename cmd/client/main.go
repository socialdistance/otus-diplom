package main

import (
	"context"
	"flag"
	"io"
	"log"
	"os/signal"
	internalgrpc "static_collector/internal/server/grpc" //nolint:gci
	"syscall"

	"google.golang.org/grpc"
)

var (
	port string
	n, m int
)

func init() {
	flag.StringVar(&port, "port", ":50051", "Listen port")
	flag.IntVar(&n, "n", 5, "interval to get statistic (sec)")
	flag.IntVar(&m, "m", 5, "interval to average statistic (sec)")
}

func main() {
	flag.Parse()

	conn, err := grpc.Dial(port, grpc.WithInsecure()) //nolint:staticcheck
	if err != nil {
		log.Fatalf("can't connect with server %v", err)
	}

	client := internalgrpc.NewStreamServiceClient(conn)
	in := internalgrpc.GatherRequest{
		N: int64(n),
		M: int64(m),
	}

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	stream, err := client.ListGather(context.Background(), &in)
	if err != nil {
		log.Fatalf("open stream error: %s", err) //nolint:gocritic
	}

	done := make(chan struct{})

	go func() {
		for {
			resp, err := stream.Recv()
			if err == io.EOF { //nolint:errorlint
				close(done)
				return
			}
			if err != nil {
				log.Fatalf("cannot receive %v", err)
			}

			for _, t := range resp.Result {
				log.Println(t)
				log.Println("-------")
			}
		}
	}()

	<-done
	<-ctx.Done()
	log.Printf("finished")
}
