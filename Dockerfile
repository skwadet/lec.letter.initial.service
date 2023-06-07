FROM golang:alpine as builder

RUN set -ex &&\
    apk add --no-progress --no-cache \
    gcc \
    musl-dev

WORKDIR /src

COPY go.mod /src
COPY go.sum /src

RUN go mod download

COPY . /src

RUN GO111MODULE=on GOOS=linux go build -a  -o main /src/cmd/main.go

FROM alpine

RUN apk add --no-cache tzdata
ENV TZ=Europe/Moscow

COPY --from=builder /src/main /

ARG ENV_STRING

RUN echo "${ENV_STRING}" > /.env

CMD ["/main"]