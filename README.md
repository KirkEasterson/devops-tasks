# Tasks
- Create a webserver in Go, which answers with a number that increases per request
- Create a /metrics endpoint that shows this number in a request_sum metric
- Build a docker image with the application
- Use k3d image import to make it available in a k3d cluster
    - NOTE: The following repo can be used as a template https://github.com/tenstad/k8s-infra/
    - You can choose if you want to follow his repo or setup your own k3d cluster. (k3d is a simplified k8s, running with docker)
- Make a k8s/k3d deployment, service and ingress, and add the hostname in Windwos/mac DNS
- Make a dashboard in Grafana, which shows how many requests that have been sent to the webserver


# Results

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

- For reference:
    - https://docs.docker.com/language/golang/build-images/

```
docker build -t gocounter .
docker run -p 8080:8080 gocounter
```

Usage:
```
curl http://localhost:8080/metrics/
```

### kubernetes

- For reference:
    - https://www.middlewareinventory.com/blog/deploy-docker-image-to-kubernetes/
    - https://docs.docker.com/get-started/kube-deploy/

```
minikube start
minikube addons enable metrics-server # optional
kubectl create -f gocounter.yaml
```
TODO: finish this task
