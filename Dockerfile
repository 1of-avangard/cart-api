FROM golang:alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /cart-bin ./cmd/main.go

FROM alpine:3.18

WORKDIR /

COPY --from=builder /cart-bin /cart-bin
COPY --from=builder /app/config /config

EXPOSE 3000

CMD ["/cart-bin"]
