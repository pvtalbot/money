# syntax=docker/dockerfile:1

FROM golang:1.18
WORKDIR /internal
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN go build -o /server
RUN chmod +x server

CMD [ "/server" ]