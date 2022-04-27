FROM golang:1.18
WORKDIR /app

RUN GO111MODULE=off go get github.com/cosmtrek/air

CMD air