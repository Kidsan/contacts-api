package main

import (
	"context"

	"github.com/google/uuid"
	contactsapi "github.com/kidsan/contacts-api"
	pb "github.com/kidsan/contacts-api/protobuffer"
	"google.golang.org/grpc"
)

type Client struct {
	apiClient pb.ContactsClient
}

func NewClient(host string) (*Client, error) {
	conn, err := grpc.Dial(host, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	return &Client{
		apiClient: pb.NewContactsClient(conn),
	}, nil
}

func (c *Client) GetContact(ctx context.Context, id string) (contactsapi.Contact, error) {
	found, err := c.apiClient.Find(ctx, &pb.ContactId{Id: id})
	if err != nil {
		return contactsapi.Contact{}, err
	}
	contactId, err := uuid.Parse(found.GetId())
	if err != nil {
		return contactsapi.Contact{}, err
	}

	return contactsapi.Contact{
		ID:   contactId,
		Name: found.GetName(),
	}, nil

}
