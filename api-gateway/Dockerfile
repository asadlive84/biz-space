# build stage
FROM golang:alpine AS builder
RUN apk add --no-cache git gcc libc-dev make


WORKDIR /go/src/app
COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .
RUN go mod tidy
RUN go mod vendor


EXPOSE 3000

CMD ["make", "server"]
