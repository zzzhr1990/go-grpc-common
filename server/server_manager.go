package server

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

//RunAndHold Hold the GRPC server running.
func RunAndHold(s *grpc.Server, portNumber int) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", portNumber))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		os.Exit(1)
	}
	reflection.Register(s)
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		log.Infoln("Waiting SIGTERM...")
		<-c
		log.Infoln("Do clean jobs...")
		s.Stop()
		// os.Exit(0)
	}()
	log.Infof("Starting server tcp on %v", portNumber)
	if err := s.Serve(lis); err != nil {
		log.Errorf("failed to serve: %v", portNumber)
		os.Exit(1)
	}
}
