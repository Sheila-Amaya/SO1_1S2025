
_Crear el proyecto en GCP_
Crear un nuevo proyecto: so1-tarea4-2025

_Habilitar APIs necesarias_
Kubernetes Engine API
Container Registry API
Compute Engine API

_Instalar herramientas_

```
gcloud init
gcloud auth login
gcloud config set project  so1-tarea4-2025
gcloud components install kubectl
```

_Crear cluster kubernetes_
```
gcloud container clusters create so1-cluster \
  --zone us-central1-a \
  --num-nodes=2

```
_Construir y subir la imagen a Container Registry_
```
docker build -t gcr.io/so1-tarea4-2025/so1-app .
docker push gcr.io/so1-tarea4-2025/so1-app
```
_Conectarse al cluster kubernetes_
```
gcloud container clusters get-credentials so1-cluster --zone us-central1-a
```

_Instalar NGINX Ingress Controller_
```
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/controller-v1.10.0/deploy/static/provider/cloud/deploy.yaml
```

verificar si se instalo
kubectl get pods -n ingress-nginx

_Aplicar los YAMLs_
```
kubectl apply -f k8s/deployment.yaml
kubectl apply -f k8s/service.yaml
kubectl apply -f k8s/ingress.yaml
```

_obtener ip externa del ingress_
```
kubectl get ingress
```

