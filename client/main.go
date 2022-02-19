package main

import (
	"context"
	"log"
	"time"

	pb "github.com/kidsan/contacts-api/protobuffer"
	"google.golang.org/grpc"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial("0.0.0.0:3000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewContactsClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	newContact := &pb.ContactRequest{
		Name: "John",
	}
	res, err := c.Save(ctx, newContact)
	if err != nil {
		log.Fatalf("could not SAVE: %v", err)
	}
	log.Printf("Saved: %s", res.Id)
	r, err := c.Get(ctx, &pb.GetRequest{})
	if err != nil {
		log.Fatalf("could not GET: %v", err)
	}
	log.Printf("All contacts: %s", r.GetContacts())

	found, err := c.Find(ctx, &pb.ContactId{Id: res.Id})
	if err != nil {
		log.Fatalf("could not FIND: %v", err)
	}
	log.Printf("Found: %s", found.Name)
}
