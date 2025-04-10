package main

import (
    "context"
    "fmt"
    "log"
    "net"

    pb "proto"
    "google.golang.org/grpc"
)

type saludoServer struct {
    pb.UnimplementedSaludoServiceServer
}

func (s *saludoServer) Saludar(ctx context.Context, req *pb.SaludoRequest) (*pb.SaludoResponse, error) {
    mensaje := fmt.Sprintf("Hola %s, ¡bienvenida a gRPC!", req.Nombre)
    return &pb.SaludoResponse{Mensaje: mensaje}, nil
}

func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
    log.Fatalf("Error al escuchar: %v", err)
    }
    grpcServer := grpc.NewServer()
    pb.RegisterSaludoServiceServer(grpcServer, &saludoServer{})
    fmt.Println("Servidor gRPC escuchando en puerto 50051...")
    grpcServer.Serve(lis)
}