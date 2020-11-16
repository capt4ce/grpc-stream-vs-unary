package communication

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/capt4ce/grpc-stream-vs-unary/monitoring"
	pb "github.com/capt4ce/grpc-stream-vs-unary/proto"
	"google.golang.org/grpc"
)

type Unary struct {
	Address string
	Client  pb.MainServiceClient
}
type UnaryServer struct {
	pb.UnimplementedMainServiceServer
}

func StartUnaryServerInstance(port string) {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterMainServiceServer(s, &UnaryServer{})
	go func() {
		fmt.Println("Unary Server started at= ", port)
		err := s.Serve(lis)
		if err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
		fmt.Println("Unary Server started at= ", port)
	}()
}

func CreateUnaryClient(address string) CommunicationInterface {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	client := pb.NewMainServiceClient(conn)

	return Unary{
		Address: address,
		Client:  client,
	}
}

func (u Unary) SendRequest() {
	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := u.Client.SendUnaryRequest(ctx, &pb.UnaryRequest{Req: 1})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
}
func (u Unary) SendResponseResponse() {}

func (us *UnaryServer) SendUnaryRequest(ctx context.Context, in *pb.UnaryRequest) (*pb.UnaryReply, error) {
	monitoring.IncrementCounter()
	defer monitoring.DecrementCounter()

	// logrus.Println("request", in.GetReq())

	time.Sleep(30 * time.Millisecond)
	return &pb.UnaryReply{Res: 123}, nil
}
