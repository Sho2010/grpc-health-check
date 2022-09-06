package main

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"os/signal"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

func main() {

	port := 50051
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()

	healthSrv := health.NewServer()
	healthpb.RegisterHealthServer(s, healthSrv)
	healthSrv.SetServingStatus("mygrpc", healthpb.HealthCheckResponse_SERVING)

	reflection.Register(s)

	go func() {
		// tick毎にランダムにステータス変わる
		for range time.Tick(3 * time.Second) {
			if rand.Intn(2) == 1 {
				healthSrv.SetServingStatus("health-check-test", healthpb.HealthCheckResponse_SERVING)
			} else {
				healthSrv.SetServingStatus("health-check-test", healthpb.HealthCheckResponse_NOT_SERVING)
			}
		}
	}()

	go func() {
		log.Printf("start gRPC server port: %v", port)
		s.Serve(listener)
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("stopping gRPC server...")
	s.GracefulStop()
}
