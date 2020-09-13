FROM golang:latest

RUN apt-get update && apt-get install build-essential sqlite3 -y
RUN go get github.com/mattn/go-sqlite3
RUN go install github.com/mattn/go-sqlite3

RUN mkdir -p /go/src/5pmGebetServer
WORKDIR /go/src/5pmGebetServer
COPY . .

ARG MASTER_KEY
RUN sqlite3 pray.db < ddl.sql
RUN sqlite3 pray.db "INSERT INTO authorization(auth_key) VALUES($MASTER_KEY)"

RUN go build
RUN go install

ARG PORT
ENV PORT=${PORT}

EXPOSE 8080
CMD go run 5pmGebetServer -p $PORT
