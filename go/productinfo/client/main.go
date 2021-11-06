package main

import (
	"context"
	pb "productinfo/client/ecommerce"
	"time"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("can't connect to gRPC server: %s", err)
	}
	defer conn.Close()

	c := pb.NewProductInfoClient(conn)

	name := "Apple iPhone 13"
	description := `Meet Apple iPhone 13. All-new dual-camera system with
	Ultra Wide and Night mode.`
	price := float32(1000.0)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.AddProduct(ctx,
		&pb.Product{Name: name, Description: description, Price: price})
	if err != nil {
		log.Fatalf("Couldn't add product: %s", err)
	}
	log.Infof("Product ID: %s added successfully", r.Value)

	product, err := c.GetProduct(ctx, &pb.ProductID{Value: r.Value})
	if err != nil {
		log.Fatalf("Couldn't get product: %s", err)
	}
	log.Printf("Product: %s", product)
}
