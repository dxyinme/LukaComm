package main

import (
	"context"
	pb "github.com/dxyinme/LukaComm/TestPb"
	"google.golang.org/grpc"
	"log"
	"sync"
	"time"
)

func main() {
	startTime := time.Now()
	var reqCount int = 25000
	var wg sync.WaitGroup
	wg.Add(reqCount)
	for i := 0 ; i < reqCount ; i ++ {
		go func() {
			func() {
				conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
				if err != nil {
					log.Println(err)
					return
				}
				defer conn.Close()
				c := pb.NewTestServerClient(conn)
				resp,err := c.GrpcCall(context.Background(), &pb.TestMsg{Name:"UUUUU"})
				if err != nil {
					log.Println(err)
					return
				}
				if resp.Ans != "echoUUUUU" {
					log.Println("err response")
				}
			}()
			wg.Done()
		}()
	}
	wg.Wait()
	log.Printf("use time : %v ms", time.Now().Sub(startTime).Milliseconds())
}
