package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "HT05/proto"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedSaludoServiceServer
}

func (s *server) Saludar(ctx context.Context, req *pb.SaludoRequest) (*pb.SaludoResponse, error) {
	msg := fmt.Sprintf("¡Hola, %s!", req.Nombre)
	return &pb.SaludoResponse{Mensaje: msg}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("No se pudo escuchar: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterSaludoServiceServer(s, &server{})

	log.Println("Servidor escuchando en :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Fallo al servir: %v", err)
	}
}
