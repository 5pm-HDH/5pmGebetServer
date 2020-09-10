FROM golang:latest

RUN apt-get update && apt-get install build-essential -y
RUN go get github.com/mattn/go-sqlite3
RUN go install github.com/mattn/go-sqlite3

RUN mkdir -p /go/src/5pmGebetServer
COPY . .

EXPOSE 80

CMD ["5pmGebetServer"]
