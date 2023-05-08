package main

import (
	"context"
	"fmt"
	pb "go-payment"
	"math/rand"
	"os"
	"sync"
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
	start := time.Now()
	fmt.Println(start)
	wg := &sync.WaitGroup{}
	wg.Add(5000)
	for i := 0; i < 5; i++ {
		for j := 0; j < 1000; j++ {
			go func(i int, j int, wg *sync.WaitGroup) {
				ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
				defer cancel()
				defer wg.Done()
				// fmt.Print(ctx)
				r, err := c.TransferPayment(ctx, &pb.TransferPaymentRequest{From: rand.Int63n(100000), To: rand.Int63n(100000), Amount: 1, PaymentId: int64(i*1000 + j)})
				if err != nil {
					fmt.Println(err)
					fmt.Println(i, time.Since(start))
					os.Exit(1)
				}
				fmt.Println(r.GetState(), r.GetPaymentId())
			}(i, j, wg)
		}
		time.Sleep(time.Second * 1)
	}
	wg.Wait()
	fmt.Println(time.Since(start))
}
