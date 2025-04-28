## 1. Limpiar minikube (opcional)

```bash
minikube delete --all --purge
docker system prune -af
```

## 2. Iniciar minikube

```bash
minikube start --driver=docker --cpus=4 --memory=4096mb
```

## 3. Habilitar addons necesarios

```bash
minikube addons enable ingress
```


---

## 4. Construir imágenes locales

```bash
eval $(minikube docker-env)

docker build -t kafka-consumer:latest ./src/kafka-consumer
docker build -t rabbitmq-consumer:latest ./src/rabbitmq-consumer
```

---

## 5. Aplicar los manifiestos de Kubernetes

Desde la raíz del proyecto:

```bash
kubectl apply -f k8s/
```

Esto crea:
- RabbitMQ
- Kafka
- Valkey
- Kafka-consumer
- RabbitMQ-consumer
- Servicios necesarios

---

## 6. Configurar RabbitMQ (una sola vez)

Entrar al pod:

```bash
kubectl exec -it rabbitmq-0 -- bash
```

Dentro:

```bash
rabbitmqctl set_permissions -p / user ".*" ".*" ".*"
exit
```

En tu máquina local:

```bash
curl -LO http://localhost:15672/cli/rabbitmqadmin
chmod +x rabbitmqadmin
```

Declarar la cola:

```bash
./rabbitmqadmin --vhost=/ --username=user --password=<PASSWORD_DEL_SECRET> declare queue name=message durable=true
```

---

## 7. Crear topic en Kafka

```bash
kubectl run kafka-client --rm -it --restart=Never --image=bitnami/kafka:4.0.0-debian-12-r3 -- \
  /opt/bitnami/kafka/bin/kafka-topics.sh \
  --bootstrap-server kafka.default.svc.cluster.local:9092 \
  --create --topic message --partitions 1 --replication-factor 1
```

---

## 8. Actualizar Passwords en RabbitMQ (si cambian)

Editar `k8s/rabbitmq-consumer-deploy.yaml`:

```yaml
env:
- name: RABBITMQ_URL
  value: "amqp://user:<PASSWORD>@rabbitmq.default.svc.cluster.local:5672/"
```

Aplicar cambios:

```bash
kubectl apply -f k8s/rabbitmq-consumer-deploy.yaml
kubectl rollout restart deployment rabbitmq-consumer
```

---

## 9. Probar envío de mensajes

**Kafka:**

```bash
kubectl exec -it kafka-controller-0 --container kafka -- bash -ilc "
echo '{\"country\":\"GT\",\"weather\":\"Soleado\"}' | /opt/bitnami/kafka/bin/kafka-console-producer.sh --bootstrap-server localhost:9092 --topic message
"
```

**RabbitMQ:**

```bash
./rabbitmqadmin --vhost=/ --username=user --password=<PASSWORD> publish routing_key=message payload='{"country":"US","weather":"Nublado"}'
```

---

## 10. Exponer Valkey para Grafana (Port-Forward)

```bash
kubectl port-forward svc/valkey 6379:6379
```

---

## 11. Configurar Grafana

1. Entrar a Grafana: [http://localhost:3000](http://localhost:3000)
2. Usuario: `admin`, Contraseña: `admin`
3. Crear nueva Data Source:
   - Tipo: **Redis**
   - URL: `localhost:6379`
   - ACL: Desactivado
   - Sin contraseña
4. Crear dashboard:
   - Query de tipo **Redis**
   - Comando **HGETALL**
   - Key: `tweets_vk`
   - Visualizar como **Bar Chart**

---

# Mini-Resumen de carpetas

| Carpeta | Contenido |
|:---|:---|
| `src/kafka-consumer` | Código fuente del Kafka consumer |
| `src/rabbitmq-consumer` | Código fuente del RabbitMQ consumer |
| `k8s/` | Manifiestos Kubernetes |
| `rabbitmqadmin` | Cliente CLI para RabbitMQ |