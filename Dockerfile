# syntax=docker/dockerfile:1

# create a build image
FROM golang:1.20 AS build-stage

WORKDIR /app

COPY go/ ./
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /go-counter

# create a lean image
FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /

COPY --from=build-stage /go-counter /go-counter

EXPOSE 8080
USER nonroot:nonroot
ENTRYPOINT ["./go-counter -p 8080"]

