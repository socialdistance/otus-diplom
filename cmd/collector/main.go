package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	internalgrpc "static_collector/internal/server/grpc"
	"syscall"
)

var configFile string
var port string

func init() {
	flag.StringVar(&configFile, "config", "./configs/config.yaml", "Path to configuration file")
	flag.StringVar(&port, "port", "50051", "Listen port")
}

func main() {
	flag.Parse()

	//config, err := internalconfig.LoadConfig(configFile)
	//if err != nil {
	//	log.Fatalf("Failed load config %s", err)
	//}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGKILL, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()

	grpc := internalgrpc.NewServer(port)

	go func() {
		if err := grpc.Start(); err != nil {
			log.Fatalf("failed to start grpc server: %s", err)
		}
	}()

	go func() {
		<-ctx.Done()
		grpc.Stop()
	}()

	<-ctx.Done()
}
