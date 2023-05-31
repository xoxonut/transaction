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
	total := 100000
	rps := 100000
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err, 123)
	}
	defer conn.Close()
	c := pb.NewPaymentServiceClient(conn)
	start := time.Now()
	fmt.Println(start)
	wg := &sync.WaitGroup{}
	wg.Add(total)
	t := total / rps
	for i := 0; i < t; i++ {
		for j := 0; j < rps; j++ {
			go func(i int, j int, wg *sync.WaitGroup, t int) {
				ctx, cancel := context.WithTimeout(context.Background(), time.Second*18000)
				defer cancel()
				defer wg.Done()
				r, err := c.TransferPayment(ctx, &pb.TransferPaymentRequest{From: rand.Int63n(1000000), To: rand.Int63n(1000000), Amount: 1, PaymentId: int64(i*t + j)})
				if err != nil {
					fmt.Println(err)
					fmt.Println(i, time.Since(start))
					os.Exit(1)
				}
				fmt.Println(r.GetState(), r.GetPaymentId())
			}(i, j, wg, rps)
		}
		time.Sleep(time.Second * 1)
	}
	wg.Wait()
	fmt.Println(time.Since(start))
}
