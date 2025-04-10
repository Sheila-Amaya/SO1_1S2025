package main

import (
    "context"
    "log"
    "time"

    pb "proto"
    "google.golang.org/grpc"
)

func main() {
    conn, err := grpc.Dial("server:50051", grpc.WithInsecure())
    if err != nil {
    log.Fatalf("No se pudo conectar: %v", err)
    }
    defer conn.Close()

    c := pb.NewSaludoServiceClient(conn)
    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel()

    resp, err := c.Saludar(ctx, &pb.SaludoRequest{Nombre: "Elizabeth"})
    if err != nil {
    log.Fatalf("Error al llamar: %v", err)
    }
    log.Printf("Respuesta del servidor: %s", resp.Mensaje)
}