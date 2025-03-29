
_Crear el proyecto en GCP_  
Crear un nuevo proyecto: **so1-tarea4-2025**

_Habilitar APIs necesarias_  
- Kubernetes Engine API  
- Container Registry API  
- Compute Engine API  

---

_instalar gcloud_

```bash
curl -fsSL https://packages.cloud.google.com/apt/doc/apt-key.gpg | \
  sudo gpg --dearmor -o /usr/share/keyrings/cloud.google.gpg

echo "deb [signed-by=/usr/share/keyrings/cloud.google.gpg] https://packages.cloud.google.com/apt cloud-sdk main" | \
  sudo tee /etc/apt/sources.list.d/google-cloud-sdk.list

sudo apt update
gcloud version
sudo apt install google-cloud-sdk -y
```

---

_Instalar herramientas GCloud en Ubuntu_

```bash
gcloud init
gcloud components install kubectl
```

---
  
_Crear cluster Kubernetes_

```bash
gcloud container clusters create so1-cluster \
  --zone us-central1-a \
  --num-nodes=2
```

---

_Construir y subir la imagen a Container Registry_

```bash
cd ~/Escritorio/SO1_1S2025/HT04/GRCP/app

docker build -t gcr.io/so1-tarea4-2025-455121/so1-app .

gcloud auth configure-docker

docker push gcr.io/so1-tarea4-2025-455121/so1-app
```

---

_Conectarse al cluster Kubernetes_

```bash
gcloud container clusters get-credentials so1-cluster --zone us-central1-a
```

---

_Instalar NGINX Ingress Controller_

```bash
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/controller-v1.10.0/deploy/static/provider/cloud/deploy.yaml
```

_verificar si se instaló_

```bash
kubectl get pods -n ingress-nginx
kubectl get svc -n ingress-nginx
```

---

_Aplicar los YAMLs_

```bash
cd ~/Escritorio/SO1_1S2025/HT04/GRCP/K8s

kubectl apply -f deployment.yaml
kubectl apply -f service.yaml
kubectl apply -f ingress.yaml
```

cargar todos los yamls sin estar en la carpeta k8s
```
kubectl apply -f ~/Escritorio/SO1_1S2025/HT04/GRCP/K8s/
```


---

_Obtener IP externa del Ingress_

```bash
kubectl get ingress
```


