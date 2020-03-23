FROM golang

RUN mkdir /helloworld_client
RUN mkdir -p /helloworld_client/api

WORKDIR /helloworld_client

COPY ./api/api.pb.go /helloworld_client/api
COPY ./main.go /helloworld_client

COPY go.mod .
COPY go.sum .

RUN go mod download

RUN go build -o helloworld_client main.go

CMD ./helloworld_client
