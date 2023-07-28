package main

import (
	"context"
	"encoding/json"
	"fmt"
	pb "go-payment"
	"log"
	"net"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/thmeitz/ksqldb-go"
	knet "github.com/thmeitz/ksqldb-go/net"
	"google.golang.org/grpc"
)

type SafeMap struct {
	mu  sync.RWMutex
	Map map[uint64]int64
}

func (sm *SafeMap) Get(key uint64) int64 {
	sm.mu.RLock()
	defer sm.mu.RUnlock()
	return sm.Map[key]
}

func (sm *SafeMap) Set(key uint64, value int64) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	sm.Map[key] = value
}

var records = SafeMap{Map: make(map[uint64]int64)}

var rowChannel = make(chan ksqldb.Row)
var headerChannel = make(chan ksqldb.Header, 1)
var f, _ = os.OpenFile("golog.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
var logger = log.New(f, "ksqldb-go ", log.LstdFlags)

func exec_ksql() {
	var options = knet.Options{BaseUrl: "http://localhost:8088",
		AllowHTTP: true}
	var kcl, _ = ksqldb.NewClientWithOptions(options)
	defer kcl.Close()
	var query = "SELECT * FROM BALANCE EMIT CHANGES;"
	ctx, cancel := context.WithTimeout(context.TODO(), 1800000*time.Second)
	defer cancel()
	e := kcl.Push(ctx, ksqldb.QueryOptions{Sql: query}, rowChannel, headerChannel)
	if e != nil {
		fmt.Println(e)
	}

}

func modify_map() {
	for row := range rowChannel {
		if row != nil {
			var id = uint64(row[0].(float64))
			var balance = int64(row[1].(float64))
			records.Set(id, balance)
		}
	}

}

type paymentDetail struct {
	Id     uint64 `json:"ID"`
	Amount int64  `json:"AMOUNT"`
}

var p, _ = kafka.NewProducer(&kafka.ConfigMap{
	"bootstrap.servers": "localhost:37029",
	"acks":              "all"})
var delivery_chan = make(chan kafka.Event, 10000)
var topic = "Payment"

func sendPayment(from int, to int, amount int) {
	if records.Get(uint64(from)) < int64(amount) {
		return
	}
	a := paymentDetail{Id: uint64(from), Amount: int64(-amount)}
	b, _ := json.Marshal(a)
	p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Key:            []byte(strconv.Itoa(from)),
		Value:          []byte(b)},
		delivery_chan,
	)
	c := paymentDetail{Id: uint64(to), Amount: int64(-amount)}
	d, _ := json.Marshal(c)
	p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Key:            []byte(strconv.Itoa(to)),
		Value:          []byte(d)},
		delivery_chan,
	)
}

type server struct {
	pb.UnimplementedPaymentServiceServer
}

func (s *server) TransferPayment(ctx context.Context, in *pb.TransferPaymentRequest) (*pb.TransferPaymentResponse, error) {
	sendPayment(int(in.GetFrom()), int(in.GetTo()), int(in.GetAmount()))
	sendPayment(int(in.GetTo()), int(in.GetFrom()), -int(in.GetAmount()))
	// fmt.Printf("Received: %v\n", in.GetFrom())
	return &pb.TransferPaymentResponse{State: 0, PaymentId: in.GetPaymentId()}, nil
}
func main() {
	go func() {
		fmt.Println("start")
		for e := range delivery_chan {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					logger.Printf("Failed to deliver message: %v\n", ev.TopicPartition)
				}
			}
		}
	}()
	go modify_map()
	go exec_ksql()
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		logger.Println(err)
	}
	s := grpc.NewServer()
	pb.RegisterPaymentServiceServer(s, &server{})
	fmt.Printf("Server is running on port %v\n", lis.Addr())
	if err := s.Serve(lis); err != nil {
		logger.Println(err)
	}
}
