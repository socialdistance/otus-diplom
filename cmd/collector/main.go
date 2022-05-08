package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os/signal"
	"syscall"

	internalconfig "static_collector/internal/config"    //nolint:gci
	internalgrpc "static_collector/internal/server/grpc" //nolint:gci
)

var (
	configFile string
	port       string
)

func init() {
	flag.StringVar(&configFile, "config", "./configs/configs.yaml", "Path to configuration file")
	flag.StringVar(&port, "port", "50051", "Listen port")
}

func main() {
	flag.Parse()

	config, err := internalconfig.LoadConfig(configFile)
	if err != nil {
		log.Fatalf("Failed load config %s", err)
	}

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	grpc := internalgrpc.NewServer(port, *config)

	go func() {
		if err := grpc.Start(); err != nil {
			log.Fatalf("failed to start grpc server: %s", err)
		}
	}()

	<-ctx.Done()
	fmt.Println("Graceful shutdown")
	grpc.Stop()
	fmt.Println("Graceful shutdown contex")
}
