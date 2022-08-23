# syntax=docker/dockerfile:1

FROM golang:1.18-alpine AS builder
WORKDIR /internal
RUN apk add --no-cache bash
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN go build -o /server

FROM golang:1.18-alpine
WORKDIR /internal
COPY --from=builder server .
RUN chmod +x ./server

CMD [ "./server" ]