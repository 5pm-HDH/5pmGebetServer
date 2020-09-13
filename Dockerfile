FROM golang:latest

RUN apt-get update && apt-get install build-essential -y
RUN go get github.com/mattn/go-sqlite3
RUN go install github.com/mattn/go-sqlite3

RUN mkdir -p /go/src/5pmGebetServer
WORKDIR /go/src/5pmGebetServer
COPY . .

RUN go build
RUN go install

EXPOSE 8080
CMD ["go", "run", "5pmGebetServer"]
