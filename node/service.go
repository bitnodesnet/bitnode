package node

import (
	"context"
	"log"
	"net"
	"os/exec"
	"strings"

	"github.com/bitnodesnet/bitnode/service"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func (n *Node) Ping(ctx context.Context, in *service.Empty) (*service.Pong, error) {
	log.Println("Ping request received")
	return &service.Pong{Message: "ok"}, nil
}

func (n *Node) Execute(ctx context.Context, in *service.ExecutionRequest) (*service.ExecutionResponse, error) {
	log.Println("Execute request received [" + in.Command + "]")

	command := strings.Replace(in.Command, "\n", "", -1)
	args := strings.Split(command, " ")
	baseCommand := args[0]
	args = args[1:]

	response := &service.ExecutionResponse{Command: command}
	result, err := exec.Command(baseCommand, args...).Output()
	if err != nil {
		response.StdErr = err.Error()
	}

	response.StdOut = string(result)

	return response, nil
}

func (n *Node) Details(ctx context.Context, in *service.Empty) (*service.NodeDetails, error) {
	log.Println("Details request received")
	return n.GetNodeDetails(), nil
}

func (n *Node) startAPIServer() {
	log.Printf("Starting node service on port %s", viper.GetString("rpc.port"))

	lis, err := net.Listen("tcp", ":"+viper.GetString("rpc.port"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Creates a new gRPC server
	s := grpc.NewServer()
	service.RegisterNodeServer(s, n)
	s.Serve(lis)
}
