# stage 1: build the application
FROM golang:1.24 AS build-stage

WORKDIR /app

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /go-app

# stage 2: deploy the application binary
FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /

COPY --from=build-stage /go-app /go-app

EXPOSE 1323

USER nonroot:nonroot

ENTRYPOINT ["/go-app"]