FROM golang:alpine3.10 as builder
WORKDIR /root
COPY go.mod go.sum ./
RUN GOPROXY=https://goproxy.cn go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o ./bin/app ./cmd/web/main.go

FROM alpine:3.10
WORKDIR /root
COPY --from=builder /root/bin/app .
CMD ["sh"]