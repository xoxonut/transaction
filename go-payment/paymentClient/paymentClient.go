package main

import (
	"context"
	"fmt"
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
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*100)
	defer cancel()
	start := time.Now()
	for i := 0; i < 100000; i++ {
		r, err := c.TransferPayment(ctx, &pb.TransferPaymentRequest{From: 777, To: 666, Amount: 1, PaymentId: 132})
		if err != nil {
			fmt.Println(err)
			fmt.Println(time.Since(start))
			break
		}
		fmt.Println(r.GetState(), r.GetPaymentId())
	}
	fmt.Println(time.Since(start))
}
