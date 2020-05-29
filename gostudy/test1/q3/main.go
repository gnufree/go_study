package main

import (
"context"
"log"
"os"
"time"
)

var logg *log.Logger
func someHandler()  {
	//ctx, cancel := context.WithCancel(context.Background())
	ctx,cancel := context.WithDeadline(context.Background(), time.Now().Add(5 * time.Second))
	go doStuff(ctx)
	time.Sleep(10 * time.Second)
	cancel()
}

func doStuff(ctx context.Context)  {
	for {
		time.Sleep(1 * time.Second)
		select {
		case <- ctx.Done():
			logg.Println("done")
			return
		default:
			logg.Print("work")
		}
	}
}

func main()  {
	logg = log.New(os.Stdout, "",log.Ltime)
	someHandler()
	logg.Printf("down")
}