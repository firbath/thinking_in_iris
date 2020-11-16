FROM golang:1.14

WORKDIR /go/src/app
COPY . .

# 容器环境变量添加，会覆盖默认的变量值
ENV GO111MODULE=on
ENV GOPROXY=http://172.16.101.58:9998

RUN go get -d -v ./... &&\
    go build -o hello_iris -v main/hello_iris.go

CMD ["./hello_iris"]