FROM golang

RUN mkdir /helloworld_server
RUN mkdir -p /helloworld_server/api

WORKDIR /helloworld_server

COPY ./api/api.pb.go /helloworld_server/api
COPY ./main.go /helloworld_server
COPY ./main_test.go /helloworld_server

COPY go.mod .
COPY go.sum .

RUN go mod download

RUN go test *.go

RUN go build -o helloworld_server main.go

CMD ./helloworld_server

EXPOSE 8081

