FROM golang:1.21 AS build

WORKDIR /usr/src/lpm

COPY go.mod go.sum .

RUN go mod download

COPY . ./

RUN go build -o /usr/local/bin/lpm 

FROM ubuntu:latest

WORKDIR /usr/local/bin

COPY --from=build /usr/local/bin/lpm .

RUN apt-get update && \
    apt-get install -y iputils-ping iperf3 tcpdump && \
    apt-get clean
    

ENTRYPOINT ["lpm"]