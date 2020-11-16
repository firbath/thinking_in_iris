FROM base_go

WORKDIR /go/src/app
COPY . .
RUN go get -d -v ./... &&\
    go install -v main/hello_iris.go

CMD ["hello_iris"]