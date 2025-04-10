IMG_NAME=grpc-server
REGISTRY=
PROJECT=tarea5

docker build -t $REGISTRY/$PROJECT/$IMG_NAME .
docker push $REGISTRY/$PROJECT/$IMG_NAME