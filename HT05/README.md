## Instalaciones previas

### 1. Instalar Go
Descargar desde: https://go.dev/dl/  
Versión usada: `>= 1.21`

### 2. Instalar Docker 

### 3. Instalar `protoc` (Protocol Buffers Compiler)
Descargar desde: https://github.com/protocolbuffers/protobuf/releases  
Extraer y agregar `protoc/bin` al `PATH`

### 4. Instalar plugins de Go para Protocol Buffers
```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.31.0
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.3.0

protoc-gen-go --version
protoc-gen-go-grpc --version
```
Agregar `%USERPROFILE%\go\bin` al PATH si no está

---

####  Generar código gRPC desde .proto

1. Crear archivo `proto/saludo.proto`
2. Ejecutar desde la raíz del proyecto:
```bash
protoc --go_out=. --go-grpc_out=. proto/saludo.proto
```
Esto genera:
- `proto/saludo.pb.go`
- `proto/saludo_grpc.pb.go`

---

#### Configuración de Go Modules

Desde la carpeta `HT05/`, ejecutar:
```bash
go mod init HT05
go mod tidy
```

Agregar manualmente versiones compatibles en `go.mod`:
```go
require (
    google.golang.org/grpc v1.56.0
    google.golang.org/protobuf v1.31.0
    google.golang.org/genproto v0.0.0-20230822172742-b8732ec3820d
)

go 1.21
```

#### Ejecución

Desde la carpeta `HT05/`:

```bash
docker-compose up --build
```

#### Resultado :

```
grpc-server | Servidor escuchando en :50051
grpc-client | Respuesta del servidor: ¡Hola, someone!
```
