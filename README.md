### go server
```
go mod download
go run .
```

Usage:
```
curl http://localhost:8080/metrics/
```

### docker

- For reference https://docs.docker.com/language/golang/build-images/

```
docker build -t gocounter .
docker run -p 8080:8080 gocounter
```

Usage:
```
curl http://localhost:8080/metrics/
```

### kubernetes
```
minikube start
minikube addons enable metrics-server # optional
kubectl create -f gocounter.yaml
```
TODO: finish this task