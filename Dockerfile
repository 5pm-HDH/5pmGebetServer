# STAGE 1: Go build
FROM golang:buster-slim AS build

RUN apt-get update \
 && apt-get install build-essential sqlite3 -y \
 && go get github.com/mattn/go-sqlite3 \
 && go install github.com/mattn/go-sqlite3

WORKDIR /5pmGebetServer
COPY . .

ARG MASTER_KEY
RUN sqlite3 pray.db < ddl.sql \
# && sqlite3 pray.db "INSERT INTO authorization(auth_key) VALUES($MASTER_KEY)"

RUN CGO_ENABLED=0 go build -o 5pm.bin main.go


# STAGE 2: UI build
FROM node:current-buster-slim AS ui

WORKDIR /5pmUI
COPY web/static/ .

RUN npm install && npm start


# STAGE 3: Final image
FROM alpine:latest

WORKDIR /5pm
COPY --from=build /5pmGebetServer/5pm.bin .
COPY --from=build /5pmGebetServer/pray.db .

WORKDIR /5pm/web/static
COPY --from=ui /5pmUI .

EXPOSE 8080

USER 1000
ENTRYPOINT "/5pm/5pm.bin -p ${PORT} -m ${MASTERKEY}"
