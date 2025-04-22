FROM --platform=linux/arm64 golang:1.23 as builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy

COPY . .

WORKDIR /app/cmd/simpleserver
RUN go build -o /simpleserver .

FROM --platform=linux/arm64 debian:bookworm
RUN apt-get update && apt-get install -y \
    ca-certificates \
 && rm -rf /var/lib/apt/lists/*

WORKDIR /root/
COPY --from=builder /simpleserver .
COPY var/config.yaml /var/config.yaml

EXPOSE 443
CMD ["./simpleserver"]
