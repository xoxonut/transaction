package main

import (
	"context"
	"fmt"
	pb "go-payment"
	"math/rand"
	"os"
	"time"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err, 123)
	}
	defer conn.Close()
	c := pb.NewPaymentServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
	defer cancel()
	start := time.Now()
	for i := 0; i < 10000; i++ {
		tmp := time.Now()
		for j := 0; j < 10; j++ {
			go func() {
				r, err := c.TransferPayment(ctx, &pb.TransferPaymentRequest{From: rand.Int63n(100000), To: rand.Int63n(100000), Amount: 1, PaymentId: 132})
				if err != nil {
					fmt.Println(err)
					fmt.Println(i, time.Since(start))
					os.Exit(1)
				}
				fmt.Println(time.Since(tmp))
				fmt.Println(r.GetState(), r.GetPaymentId())
			}()
		}
		time.Sleep(time.Second * 1)
	}
	fmt.Println(time.Since(start))
}
