
FROM golang:1.21.0-alpine3.18

WORKDIR /app

ADD . /app/
ADD docker/dockerize /bin/

RUN export GO111MODULE=on && GOPROXY=https://goproxy.io
RUN go mod download

#RUN make
RUN go build -o ./bin/ ./cmd/profile && go build -o ./bin/ ./cmd/consultation

CMD ["/app/bin/init"]
