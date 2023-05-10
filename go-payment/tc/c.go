package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	pb "go-payment"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err, 123)
	}
	defer conn.Close()
	c := pb.NewPaymentServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*18000)
	defer cancel()
	start := time.Now()
	for i := 0; i < 10000; i++ {
		r, err := c.TransferPayment(ctx, &pb.TransferPaymentRequest{From: rand.Int63n(100000), To: rand.Int63n(100000), Amount: 1, PaymentId: 132})
		if err != nil {
			fmt.Println(err)
			fmt.Println(time.Since(start))
			break
		}
		fmt.Println(r.GetState(), r.GetPaymentId())
	}
	fmt.Println(time.Since(start))
}
