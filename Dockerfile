FROM golang:1.23-alpine AS builder

WORKDIR /usr/goutube

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o goutube-api


FROM alpine:3.20
WORKDIR /usr/goutube

ENV GIN_MODE="release"
EXPOSE 3000

COPY --from=builder /usr/goutube/goutube-api .

ENTRYPOINT ["./goutube-api"]

