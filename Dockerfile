# STAGE 1: Go build
FROM golang:buster AS build

RUN apt-get update \
 && apt-get install build-essential sqlite3 -y \
 && go get github.com/mattn/go-sqlite3 \
 && go install github.com/mattn/go-sqlite3

WORKDIR /5pmGebetServer
COPY . .

RUN go build -o 5pm.bin main.go database.go


# STAGE 2: UI build
FROM node:current-buster-slim AS ui

WORKDIR /5pmUI
COPY web/static/ .

RUN npm install && npm start


# STAGE 3: Final image
FROM debian:stable-slim

WORKDIR /5pm/web/static
COPY --from=ui /5pmUI .

WORKDIR /5pm
COPY --from=build /5pmGebetServer/5pm.bin .
COPY --from=build /5pmGebetServer/ddl.sql .

EXPOSE 8080

RUN chown -R 1000:1000 /5pm

USER 1000
ENTRYPOINT "./5pm.bin"
