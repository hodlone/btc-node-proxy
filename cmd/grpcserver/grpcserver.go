package grpcserver

import (
	"context"
	"fmt"

	"github.com/btcsuite/btcd/rpcclient"
	"github.com/btcsuite/btcutil"

	"github.com/NodeHodl/btc-node-proxy/pkg/bitcoin"
	"github.com/NodeHodl/btc-node-proxy/pkg/grpc"
	"github.com/NodeHodl/btc-node-proxy/pkg/logger"
	"github.com/NodeHodl/gRPC/dist/protos/services/nodeproxypb"
)

type server struct {
	bitcoinClient *rpcclient.Client
	nodeproxypb.UnimplementedNodeProxyServiceServer
}

var grpcLogger = logger.New("grpcserver")

// Start ...
func Start() {
	lis, srv := grpc.NewServer(grpcLogger)

	nodeproxypb.RegisterNodeProxyServiceServer(srv, &server{
		bitcoinClient: bitcoin.NewBitcoinRPCClient(),
	})

	srv.Serve(*lis)
}

// GenerateAddress Recieves a string of comma separated pairs and a supported exchange
func (s *server) GenerateAddress(ctx context.Context, req *nodeproxypb.GenerateAddressRequest) (*nodeproxypb.GenerateAddressResponse, error) {
	var err error
	var address btcutil.Address

	if address, err = s.bitcoinClient.GetNewAddress(""); err != nil {
		fmt.Printf("KOKO error=%v", err.Error())
		return nil, err
	}

	fmt.Printf("EncodeAddress=%v", address.EncodeAddress())
	fmt.Printf("ScriptAddress=%v", address.ScriptAddress())
	fmt.Printf("String=%v", address.String())
	// create new crypto client
	// make request to node and return response

	return &nodeproxypb.GenerateAddressResponse{Address: address.String()}, nil
}
