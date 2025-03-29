```bash
# Crear imagen y subir a GCP
docker build -t gcr.io/so1-tarea4-2025-455121/so1-app .
gcloud auth configure-docker
docker push gcr.io/so1-tarea4-2025-455121/so1-app

# Crear clúster limpio sin NEG/Ingress
gcloud container clusters create so1-nodeport-cluster   --zone us-central1-a   --num-nodes=2   --no-enable-ip-alias   --no-enable-intra-node-visibility   --disk-type=pd-standard

# Conectar al clúster
gcloud container clusters get-credentials so1-nodeport-cluster --zone us-central1-a

# Aplicar archivos Kubernetes
kubectl apply -f deployment.yaml
kubectl apply -f service-nodeport.yaml

# Crear regla de firewall para NodePort
gcloud compute firewall-rules create allow-nodeport-30081   --allow tcp:30081   --direction=INGRESS   --priority=1000   --network=default
```

---

como ya tenemos la imagen en gcp

```bash
# Eliminar clúster antiguo con configuraciones automáticas no deseadas
gcloud container clusters delete so1-cluster --zone us-central1-a

# Verificar que no queden clústeres activos
gcloud container clusters list
```

```bash
# Crear clúster limpio
gcloud container clusters create so1-nodeport-cluster \
  --zone us-central1-a \
  --num-nodes=2 \
  --no-enable-ip-alias \
  --no-enable-intra-node-visibility \
  --disk-type=pd-standard
```

```bash
# Aplicar archivos YAML nuevamente
kubectl delete deployment so1-deployment
kubectl delete svc so1-service

kubectl apply -f deployment.yaml
kubectl apply -f service-nodeport.yaml
```
# Verificar acceso desde navegador
http://34.xx.xx.xx:30081
