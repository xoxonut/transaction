package main

import (
	"context"
	"fmt"
	"net"

	pb "go-payment"

	"github.com/thmeitz/ksqldb-go"
	knet "github.com/thmeitz/ksqldb-go/net"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedPaymentServiceServer
}

func transfer2Kafka(from int, to int, amount int) error {
	var options = knet.Options{
		BaseUrl:   "http://localhost:8088",
		AllowHTTP: true,
	}
	kcl, err := ksqldb.NewClientWithOptions(options)
	defer kcl.Close()
	stmt, err := ksqldb.QueryBuilder("INSERT INTO BALANCE_STREAM VALUES (?, ?);", from, -amount)
	resp, err := kcl.Execute(ksqldb.ExecOptions{KSql: *stmt})
	stmt, err = ksqldb.QueryBuilder("INSERT INTO BALANCE_STREAM VALUES (?, ?);", to, amount)
	resp, err = kcl.Execute(ksqldb.ExecOptions{KSql: *stmt})
	fmt.Println(resp)
	return err
}
func (s *server) TransferPayment(ctx context.Context, in *pb.TransferPaymentRequest) (*pb.TransferPaymentResponse, error) {
	fmt.Println(in.GetFrom(), in.GetTo(), in.GetAmount())
	err := transfer2Kafka(int(in.GetFrom()), int(in.GetTo()), int(in.GetAmount()))
	if err != nil {
		fmt.Println(err)
	}
	return &pb.TransferPaymentResponse{State: 0, PaymentId: in.GetPaymentId()}, nil
}
func main() {

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		fmt.Println(err)
	}
	s := grpc.NewServer()
	pb.RegisterPaymentServiceServer(s, &server{})
	fmt.Printf("Server is running on port %v\n", lis.Addr())
	if err := s.Serve(lis); err != nil {
		fmt.Println(err)
	}
}
