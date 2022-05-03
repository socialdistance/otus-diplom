//go:generate protoc --go_out=. --go-grpc_out=. ../../../api/gatherService.proto --proto_path=../../../api

package grpc

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"static_collector/internal/config"

	internalapp "static_collector/internal/app"
)

type Server struct {
	UnimplementedStreamServiceServer
	port       string
	grpcServer *grpc.Server
	config     config.Config
}

func NewServer(port string, config config.Config) *Server {
	server := &Server{
		port:       port,
		grpcServer: grpc.NewServer(),
		config:     config,
	}

	RegisterStreamServiceServer(server.grpcServer, server)

	return server
}

func (s *Server) Start() error {
	lsn, err := net.Listen("tcp", fmt.Sprintf(":%s", s.port))
	if err != nil {
		return fmt.Errorf("[-] Can't start GRPC server: %w", err)
	}

	fmt.Printf("[+] Start GRPC Server localhost:%s\n", s.port)

	return s.grpcServer.Serve(lsn)
}

func (s *Server) Stop() {
	s.grpcServer.GracefulStop()
}

func (s *Server) ListGather(req *GatherRequest, stream StreamService_ListGatherServer) error {
	log.Printf("Start gather resources for n = %d and m = %d", req.N, req.M)

	values := internalapp.Run(stream.Context(), req.N, req.M, s.config)
	for tick := range values {
		sender(tick, req.M, stream)
	}

	return nil
}

func sender(tick map[string][][]internalapp.Value, m int64, stream StreamService_ListGatherServer) {
	for key, value := range tick {
		keyCount := len(value[0])
		resp := GatherResponse{Result: make([]string, 0)}
		tmpRes := internalapp.CalculateRes(keyCount, value, key, m)
		resp.Result = tmpRes
		if err := stream.Send(&resp); err != nil {
			log.Printf("send error %v", err)
		}
	}
}
