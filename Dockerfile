#https://hub.docker.com/repository/docker/ygdxdyj/httpserver
FROM golang:1.17-alpine AS build
WORKDIR /
COPY main.go main.go
RUN go env -w GO111MODULE=auto && export GOPROXY=https://goproxy.cn,direct
RUN CGO_ENABLED=0 GOARCH=amd64 go build -a -o main .

FROM alpine:3.7
WORKDIR /
COPY --from=build / .
RUN ls
EXPOSE 9100
CMD ["./main"]
