package communication

import (
	"context"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net"
	"time"

	"github.com/capt4ce/grpc-stream-vs-unary/monitoring"
	pb "github.com/capt4ce/grpc-stream-vs-unary/proto"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Stream struct {
	Address string
	Stream  pb.StreamService_SendStreamRequestClient
}
type StreamServer struct {
	pb.UnimplementedStreamServiceServer
}

func CreateStreamClient(address string) CommunicationInterface {
	ctx := context.Background()

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		logrus.Fatalf("did not connect: %v", err)
	}

	client := pb.NewStreamServiceClient(conn)
	if err != nil {
		logrus.Println("error in contacting server: ", err)
	}
	stream, err := client.SendStreamRequest(ctx)
	if err != nil {
		logrus.Println("error in stream", err)
	}

	instance := &Stream{
		Address: address,
		Stream:  stream,
	}

	go func() {
		for {
			_, err := stream.Recv()
			if err == io.EOF {
				// close(done)
				continue
			}
			if err != nil {
				logrus.Fatalf("can not receive %v", err)
			}
			// logrus.Printf("========> new res %d received", resp.Res)
		}
	}()

	return *instance
}

func StartStreamServerInstance(port string) {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterStreamServiceServer(s, &StreamServer{})
	go func() {
		fmt.Println("Stream Server started at= ", port)
		err := s.Serve(lis)
		if err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
		fmt.Println("Unary Server started at= ", port)
	}()
}

func (s Stream) SendRequest() {
	err := s.Stream.Send(&pb.StreamRequest{Req: int64(rand.Intn(10))})
	if err != nil && err != io.EOF {
		logrus.Println("error sending stream data", err)
	}
}
func (s Stream) SendResponseResponse() {}

func (ss *StreamServer) SendStreamRequest(stream pb.StreamService_SendStreamRequestServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			logrus.Println("no more data")
			break
		}
		if err != nil {
			return status.Errorf(codes.Unknown, "cannot receive stream request: %v", err)
		}
		monitoring.IncrementCounter()
		// logrus.Println("request ", req.GetReq())
		go func() {
			time.Sleep(30 * time.Millisecond)
			defer monitoring.DecrementCounter()
			err = stream.Send(&pb.StreamReply{Res: req.GetReq()})
			if err != nil {
				logrus.Println(err)
			}
		}()
	}

	return nil
}
