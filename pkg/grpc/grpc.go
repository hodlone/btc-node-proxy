package grpc

import (
	"fmt"
	"net"

	"github.com/NodeHodl/btc-node-proxy/config"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

// NewServer ...
func NewServer(logger *logrus.Entry) (*net.Listener, *grpc.Server) {
	logger.Infof("Starting on port %v...", config.GrpcServerPort)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", config.GrpcServerPort))

	if err != nil {
		logger.WithFields(logrus.Fields{"error": err}).Panic("Failed to start")
	}

	srv := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_logrus.UnaryServerInterceptor(logger),
		)),
	)

	return &lis, srv
}
