FROM golang:1.24 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o ./build/executable/app ./cmd/run.go

FROM ubuntu:22.04

ENV CONFIG_PATH=./config/prod.yaml

COPY --from=builder /app/build/executable/app ./build/executable/app
COPY --from=builder /app/config ./config

EXPOSE 8080

CMD ["./build/executable/app"]
