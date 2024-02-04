package main

import (
	"github.com/ramiroribeiro/estudo_grpc/proto"
	"github.com/ramiroribeiro/estudo_grpc/structs"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", ":8089")
	if err != nil {
		log.Fatalf("error listen : %s", err)
	}

	server := grpc.NewServer()
	proto.RegisterInvoiceServer(server, &structs.InvoiceServer{})

}
