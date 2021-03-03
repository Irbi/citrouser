FROM golang:1.15-alpine AS base

RUN apk update && apk upgrade
RUN apk add --no-cache bash git openssh

WORKDIR /home/app/src
COPY . ./
RUN go mod download
WORKDIR /home/app/src

FROM base as dev

RUN apk add --no-cache autoconf automake libtool gettext gettext-dev make g++ texinfo curl

WORKDIR /root
RUN wget https://github.com/emcrisostomo/fswatch/releases/download/1.14.0/fswatch-1.14.0.tar.gz
RUN tar -xvzf fswatch-1.14.0.tar.gz
WORKDIR /root/fswatch-1.14.0
RUN ./configure
RUN make
RUN make install

WORKDIR /home/app/src

#FROM golang:1.15-alpine AS build
#
#RUN apk add --no-cache git g++ bash autoconf automake libtool gettext gettext-dev make texinfo curl
#
#COPY go.* ./
#RUN go mod download
#COPY . .
#RUN GOARCH=amd64 GOOSE=linux go build -o user-api .
#
#FROM alpine:latest
#
#WORKDIR /root/
#COPY --from=build /go/src/citroneer/user/user-api .
#RUN chmod +x ./user-api
#
#EXPOSE 8080
#
#CMD ["./user-api"]
