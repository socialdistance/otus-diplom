//go:generate protoc --go_out=. --go-grpc_out=. ../../../api/gatherService.proto --proto_path=../../../api

package grpc

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	internalapp "static_collector/internal/app"
	"strconv"
	"time"
)

type GRPCServer struct {
	UnimplementedStreamServiceServer
	port       string
	grpcServer *grpc.Server
	app        *internalapp.App
}

func NewServer(port string) *GRPCServer {
	server := &GRPCServer{
		port:       port,
		grpcServer: grpc.NewServer(),
	}

	RegisterStreamServiceServer(server.grpcServer, server)

	return server
}

func (s *GRPCServer) Start() error {
	lsn, err := net.Listen("tcp", fmt.Sprintf(":%s", s.port))
	if err != nil {
		return fmt.Errorf("[-] Can't start GRPC server: %w", err)
	}

	fmt.Printf("[+] Start GRPC Server localhost:%s\n", s.port)

	return s.grpcServer.Serve(lsn)
}

func (s *GRPCServer) Stop() {
	s.grpcServer.GracefulStop()
}

func (s *GRPCServer) ListGather(req *GatherRequest, stream StreamService_ListGatherServer) error {
	log.Printf("Start gather resources for n = %d and m = %d", req.N, req.M)

	n, err := strconv.Atoi(strconv.FormatInt(req.N, 10))
	if err != nil {
		return err
	}

	m, err := strconv.Atoi(strconv.FormatInt(req.M, 10))
	if err != nil {
		return err
	}

	//ticker := time.NewTicker(time.Duration(n) * time.Second)
	//tickerM := time.NewTicker(time.Duration(m) * time.Second)
	result := make(chan []string)

	gather := func(done <-chan struct{}) <-chan []string {
		log.Printf("gathering... %d", m)
		go func() {
			res, err := s.app.Run()
			fmt.Println("res:", res)
			if err != nil {
				for _, err := range err {
					log.Printf("Error gathering resources %s", err)
				}
			}
			select {
			case <-done:
				return
			case result <- res:
			}
		}()

		return result
	}

	sender := func(done <-chan struct{}, gatherData <-chan []string) {
		for range time.Tick(time.Duration(n) * time.Second) {
			fmt.Println("test1")
			go func() {
				select {
				case <-done:
					return
				case data := <-gatherData:
					fmt.Println("test2")
					resp := GatherResponse{Result: data}
					if err := stream.Send(&resp); err != nil {
						log.Printf("send error %v", err)
					}
					log.Printf("sending data1...")
				}
			}()
			log.Printf("sending data2...")
		}
		log.Printf("sending data3...")

	}

	done := make(chan struct{})

	for range time.Tick(time.Duration(m) * time.Second) {
		sender(done, gather(done))
	}

	<-done

	return nil
}
