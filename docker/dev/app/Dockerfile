# syntax=docker/dockerfile:1

FROM golang:1.18

WORKDIR /internal
RUN go install github.com/cespare/reflex@latest
COPY reflex.conf /

COPY start.sh /

ENTRYPOINT ["reflex", "-c", "/reflex.conf"]