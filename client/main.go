package main

import (
	"context"
	"fmt"
	pd "go-grpc/pb"
	"log"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:5051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to Connect: %v", err)
	}
	defer conn.Close()

	client := pd.NewBookServiceClient(conn)
	callBook(client)

}

func callBook(client pd.BookServiceClient) {
	res, err := client.Book(context.Background(), &pd.BookRequest{})
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(res)
	// title:"title1" price:1000
	fmt.Println(res.GetTitle())
	// title1
	fmt.Println(res.GetPrice())
	// 1000
}