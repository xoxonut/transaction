package main

import (
	"context"
	"fmt"
	"net"

	pb "go-payment"

	"github.com/thmeitz/ksqldb-go"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedPaymentServiceServer
}

var kcl ksqldb.KsqldbClient

// var connections = 100
// var kcl_channel = make(chan *ksqldb.KsqldbClient, connections)

func transfer2Kafka(from int, to int, amount int) error {
	// var options = knet.Options{
	// 	BaseUrl:   "http://localhost:8088",
	// 	AllowHTTP: true,
	// }
	// kcl, _ := ksqldb.NewClientWithOptions(options)
	// defer kcl.Close()
	// kcl := <-kcl_channel
	stmt, err := ksqldb.QueryBuilder("INSERT INTO BALANCE_STREAM VALUES (?, ?);", from, -amount)
	_, err = kcl.Execute(ksqldb.ExecOptions{KSql: *stmt})
	stmt, err = ksqldb.QueryBuilder("INSERT INTO BALANCE_STREAM VALUES (?, ?);", to, amount)
	_, err = kcl.Execute(ksqldb.ExecOptions{KSql: *stmt})
	// kcl_channel <- kcl
	return err
}
func (s *server) TransferPayment(ctx context.Context, in *pb.TransferPaymentRequest) (*pb.TransferPaymentResponse, error) {
	// fmt.Println(in.GetFrom(), in.GetTo(), in.GetAmount())
	err := transfer2Kafka(int(in.GetFrom()), int(in.GetTo()), int(in.GetAmount()))
	if err != nil {
		fmt.Println(err)
	}
	return &pb.TransferPaymentResponse{State: 0, PaymentId: in.GetPaymentId()}, nil
}
func main() {
	// var options = knet.Options{
	// 	BaseUrl:   "http://localhost:8088",
	// 	AllowHTTP: true,
	// }
	// kcl, _ = ksqldb.NewClientWithOptions(options)
	// kcls := make([]ksqldb.KsqldbClient, connections)
	// for i := 0; i < connections; i++ {
	// 	kcls[i], _ = ksqldb.NewClientWithOptions(options)
	// }
	// for i := 0; i < connections; i++ {
	// 	kcl_channel <- &kcls[i]
	// }
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
