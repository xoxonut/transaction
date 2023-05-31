package main

import (
	"context"
	"fmt"
	pb "go-payment"
	"net"
	"strconv"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"google.golang.org/grpc"
)

var p, err = kafka.NewProducer(&kafka.ConfigMap{
	"bootstrap.servers": "localhost:9092",
	"acks":              "all"})
var delivery_chan = make(chan kafka.Event, 10000)
var topic = "balance"

func sendPayment(from int, to int, amount int) {
	err = p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Key:            []byte(strconv.Itoa(from)),
		Value:          []byte(`{to:` + strconv.Itoa(to) + `,amount:` + strconv.Itoa(amount) + `}`)},
		delivery_chan,
	)
	return
}

type server struct {
	pb.UnimplementedPaymentServiceServer
}

func (s *server) TransferPayment(ctx context.Context, in *pb.TransferPaymentRequest) (*pb.TransferPaymentResponse, error) {
	sendPayment(int(in.GetFrom()), int(in.GetTo()), int(in.GetAmount()))
	// fmt.Printf("Received: %v\n", in.GetFrom())
	return &pb.TransferPaymentResponse{State: 0, PaymentId: in.GetPaymentId()}, nil
}
func main() {
	go func() {
		fmt.Println("start")
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Failed to deliver message: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Successfully produced record to topic %s partition [%d] @ offset %v\n",
						*ev.TopicPartition.Topic, ev.TopicPartition.Partition, ev.TopicPartition.Offset)
				}
			}
		}
	}()
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
