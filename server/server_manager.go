package server

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	// log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

//RunAndHold Hold the GRPC server running.
func RunAndHold(s *grpc.Server, portNumber int, f func(string)) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", portNumber))
	if err != nil {
		if f != nil {
			f(fmt.Sprintf("failed to listen: %v", err))
		}

		os.Exit(1)
	}
	reflection.Register(s)
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {

		if f != nil {
			f(fmt.Sprintf("waiting SIGTERM..."))
		}
		<-c
		if f != nil {
			f(fmt.Sprintf("do clean jobs..."))
		}
		s.Stop()
		// os.Exit(0)
	}()
	if f != nil {
		f(fmt.Sprintf("starting server tcp on %v", portNumber))
	}
	if err := s.Serve(lis); err != nil {
		if f != nil {
			f(fmt.Sprintf("failed to serve: %v", portNumber))
		}
		os.Exit(1)
	}
}
