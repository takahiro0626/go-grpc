package main

import (
	"context"
	"fmt"
	"go-grpc/pb"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	pd.UnimplementedBookServiceServer
}

func (*server) Book(context context.Context, request *pd.BookRequest) (*pd.BookResponse, error) {
	fmt.Println("Book was Invoked")

	response := &pd.BookResponse{
		Title: "title1",
		Price: 1000,
	}

	return response, nil
}

func main() {
	listen, err := net.Listen("tcp", "localhost:5051")
	if err != nil {
		log.Fatalf("Failed to Listen: %v", err)
	}

	s := grpc.NewServer()
	pd.RegisterBookServiceServer(s, &server{})

	fmt.Println("server is running")

	if err := s.Serve(listen); err != nil {
		log.Fatalf("Failed to Listen: %v", err)
	}
}
