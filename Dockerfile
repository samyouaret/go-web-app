FROM golang:1.18-alpine AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY *.go ./

RUN go build -o /go-api

# FROM gcr.io/distroless/base-debian11

# WORKDIR /

# COPY --from=builder /go-api /go-api

# EXPOSE 8080

# USER nonroot:nonroot

ENTRYPOINT [ "/go-api" ]
