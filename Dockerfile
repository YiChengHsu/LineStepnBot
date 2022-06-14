FROM --platform=linux/amd64 golang:1.16.4-alpine3.12 AS build-env

# go-ethreum requires gcc
RUN apk add --no-cache build-base linux-headers

# Force the go compiler to use modules
ENV GO111MODULE=on

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o line-stepn-bot

###

FROM alpine:3.12

WORKDIR /usr/local

# alpine does not contain ZONEINFO require by Go
RUN apk add --no-cache --update tzdata && rm -f /var/cache/apk/*

COPY --from=build-env /app/line-stepn-bot .

CMD ["/usr/local/line-stepn-bot"]
