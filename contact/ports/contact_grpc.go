package ports

import (
	"context"
	"fmt"

	"go.uber.org/zap"

	"github.com/google/uuid"
	contactsapi "github.com/kidsan/contacts-api"
	pb "github.com/kidsan/contacts-api/protobuffer"
)

type GRPCHandler struct {
	pb.UnimplementedContactsServer
	service ContactService
	logger  *zap.Logger
}

func NewContactGRPCHandler(logger *zap.Logger, s ContactService) *GRPCHandler {
	return &GRPCHandler{
		logger:  logger,
		service: s,
	}
}

func (g GRPCHandler) Get(ctx context.Context, _ *pb.GetRequest) (*pb.ContactListReply, error) {
	contacts, err := g.service.Get(ctx)
	r := &pb.ContactListReply{}
	for _, v := range contacts {
		r.Contacts = append(r.Contacts, &pb.ContactReply{
			Id:   v.ID.String(),
			Name: v.Name,
		})
	}
	if err != nil {
		return nil, fmt.Errorf("ports(contacts): could not get all contacts %w", err)
	}
	return r, nil
}

func (g GRPCHandler) Save(ctx context.Context, newContactRequest *pb.ContactRequest) (*pb.ContactReply, error) {
	newContact := contactsapi.Contact{
		ID:   uuid.New(),
		Name: newContactRequest.GetName(),
	}
	result, err := g.service.Save(ctx, newContact)
	if err != nil {
		return &pb.ContactReply{}, fmt.Errorf("ports(contacts): could save new contact %w", err)
	}
	return &pb.ContactReply{Id: result.ID.String(), Name: result.Name}, nil
}

func (g GRPCHandler) Find(ctx context.Context, id *pb.ContactId) (*pb.ContactReply, error) {
	result, err := g.service.Find(ctx, id.GetId())
	if err != nil {
		return &pb.ContactReply{}, fmt.Errorf("ports(contact): could find contact %w", err)
	}
	return &pb.ContactReply{Id: result.ID.String(), Name: result.Name}, nil
}
