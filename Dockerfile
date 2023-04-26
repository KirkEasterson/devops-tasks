# syntax=docker/dockerfile:1

# create a build image
FROM golang:1.19 AS build-stage

WORKDIR /app

COPY go/ ./
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /gocounter

# create a lean image
FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /

COPY --from=build-stage /gocounter /gocounter

EXPOSE 8080
USER nonroot:nonroot
ENTRYPOINT ["/gocounter"]
