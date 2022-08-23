# syntax=docker/dockerfile:1

FROM golang:1.18-alpine AS builder
WORKDIR /internal
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN go build -o /server

FROM golang:1.18-alpine
WORKDIR /internal
COPY --from=builder server .
RUN chmod +X ./server

CMD [ "./server" ]