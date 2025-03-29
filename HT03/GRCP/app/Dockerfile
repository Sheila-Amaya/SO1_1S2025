# Imagen base
FROM golang:1.21-alpine

# Directorio de trabajo
WORKDIR /app

# Descargar dependencias
COPY go.mod ./
RUN go mod download

# Copiar todo el código
COPY . .

# Compilar la app
RUN go build -o main .

# Exponer el puerto
EXPOSE 4000

# Ejecutar la aplicación
CMD ["./main"]
