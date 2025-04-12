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


# 🐳 Configuración de Harbor

### 1. Cuenta en [Google Cloud Platform](https://console.cloud.google.com/)
### 2. Docker desktop instalado (máquina local)
### 3. Proyecto HT05: gRPC en Go (client/server)

---

### Paso 01: Crear la Máquina Virtual en GCP

1. Ve a **Compute Engine > Crear Instancia**
2. Configura:
   - **Nombre:** `harbor-vm`
   - **Zona:** `us-central1-a`
   - **Tipo de máquina:** `e2-medium (2 vCPU, 4 GB)`
   - **Sistema operativo:** Ubuntu 20.04 LTS
   - **Disco:** 30GB estándar
3. En sección **Redes**, activa:
   - ☑️ Permitir tráfico HTTP
   - ☑️ Permitir tráfico HTTPS
4. Haz clic en **Crear**

---

### Paso 02: Instalar Docker y Docker Compose

```bash
sudo apt update && sudo apt upgrade -y
sudo apt install docker.io=26.1.3-0ubuntu1~20.04.1 -y
sudo systemctl enable docker
sudo usermod -aG docker $USER
```

_Verificación:_

```bash
docker --version
```

Instalar Docker Compose:

```bash
sudo curl -L "https://github.com/docker/compose/releases/download/1.29.2/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
sudo chmod +x /usr/local/bin/docker-compose
```

_Verificación:_

```bash
docker-compose --version
```

---

### Paso 03: Descargar Harbor

```bash
wget https://github.com/goharbor/harbor/releases/download/v2.12.0/harbor-online-installer-v2.12.0.tgz
tar -xvzf harbor-online-installer-v2.12.0.tgz
cd harbor
cp harbor.yml.tmpl harbor.yml
```

---

### Paso 04: Instalar Certbot y generar certificados



```bash
sudo apt install certbot -y
sudo fuser -k 80/tcp
sudo certbot certonly --standalone -d MI_IP.nip.io
```

Verifica archivos generados:

```bash
sudo ls /etc/letsencrypt/live/MI_IP.nip.io/
```

Debes ver:
- `fullchain.pem`
- `privkey.pem`

---

### Paso 05: Configurar Harbor.yml

Edita el archivo `harbor.yml` con:

```bash
# Verifica el contenido del archivo de configuración
nano harbor.yml
```

```yaml
hostname: MI_IP.nip.io

https:
  port: 443
  certificate: /etc/letsencrypt/live/MI_IP.nip.io/fullchain.pem
  private_key: /etc/letsencrypt/live/MI_IP.nip.io/privkey.pem
```

---

### Paso 06: Instalar Harbor

```bash
cd ~/harbor
sudo ./install.sh
```

Verifica contenedores:

```bash
sudo docker ps --format "table {{.Names}}\t{{.Status}}\t{{.Ports}}"
```

---

### Paso 07: Acceder a Harbor desde navegador

```
https://MI_IP.nip.io
```

- Usuario: `admin`
- Contraseña: definida en `harbor.yml` (por defecto: `Harbor12345`)

Desde la UI, crea un proyecto llamado `ht05-grpc`.

---

## Paso 08: Subir imágenes desde máquina local


```bash
# verificar antes tener Construidas las imagen del servidor y client
docker build -t grpc-server -f server/Dockerfile .
```


```bash

docker login https://MI_IP.nip.io

docker tag grpc-client MI_IP.nip.io/ht05-grpc/grpc-client
docker tag grpc-server MI_IP.nip.io/ht05-grpc/grpc-server

docker push MI_IP.nip.io/ht05-grpc/grpc-client
docker push MI_IP.nip.io/ht05-grpc/grpc-server

docker pull MI_IP.nip.io/ht05-grpc/grpc-client
docker pull MI_IP.nip.io/ht05-grpc/grpc-server
```

Verificar:
```bash
docker compose up
```