package structs

import (
	"context"
	"github.com/ramiroribeiro/estudo_grpc/proto"
)

type InvoiceServer struct {
	proto.UnimplementedInvoiceServer
}

func (is InvoiceServer) Create(context.Context, *proto.CreateRequest) (*proto.CreateResponse, error) {
	return &proto.CreateResponse{
		Pdf:     []byte("teste"),
		Docx:    []byte("teste"),
		Success: true,
	}, nil
}
