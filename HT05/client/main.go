package main

import (
	"context"
	"log"
	"time"

	pb "HT05/proto"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("server:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("No se pudo conectar: %v", err)
	}
	defer conn.Close()

	client := pb.NewSaludoServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := client.Saludar(ctx, &pb.SaludoRequest{Nombre: "Eli"})
	if err != nil {
		log.Fatalf("Error al saludar: %v", err)
	}

	log.Printf("Respuesta del servidor: %s", resp.Mensaje)
}
