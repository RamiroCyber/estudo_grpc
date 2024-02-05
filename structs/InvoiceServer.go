package structs

import (
	"context"
	"fmt"
	"github.com/ramiroribeiro/estudo_grpc/proto"
)

type InvoiceServer struct {
	proto.UnimplementedInvoiceServer
}

func (is InvoiceServer) Create(ctx context.Context, req *proto.CreateRequest) (*proto.CreateResponse, error) {
	message := fmt.Sprintf("De %v para %v valor: %v", req.From, req.To, req.Amount.Amount)
	fmt.Println(message)
	return &proto.CreateResponse{
		Pdf:     []byte(message),
		Docx:    []byte(message),
		Success: true,
	}, nil
}
