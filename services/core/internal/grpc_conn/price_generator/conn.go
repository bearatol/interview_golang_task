package price_generator

import (
	pb "github.com/bearatol/interview_golang_task/proto/price_generator"
	"github.com/bearatol/interview_golang_task/sevices/core/internal/mapping"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

//go:generate mockgen -source=conn.go -destination=mocks/mock.go

type grpcConn struct {
	conn   *grpc.ClientConn
	client pb.PriceGeneratorClient
}

func Close(conn *grpc.ClientConn) {
	conn.Close()
}

type Service struct {
	Conn      *grpc.ClientConn
	PriceFile *priceFile
}

func NewConn(addr string) (*Service, error) {
	conn, err := createConn(addr)
	if err != nil {
		return nil, err
	}

	return &Service{
		Conn:      conn.conn,
		PriceFile: NewPriceFile(conn),
	}, nil
}

func createConn(addr string) (*grpcConn, error) {
	conn, err := grpc.Dial(
		addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(mapping.MaxMsgSize)),
	)

	if err != nil {
		return nil, err
	}

	c := pb.NewPriceGeneratorClient(conn)

	return &grpcConn{
		conn:   conn,
		client: c,
	}, nil
}
