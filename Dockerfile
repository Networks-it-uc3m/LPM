FROM golang:1.21 AS build
ARG TARGETOS
ARG TARGETARCH

WORKDIR /usr/src/lpm

COPY go.mod go.sum .

RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=${TARGETOS:-linux} GOARCH=${TARGETARCH} go build -a -o /usr/local/bin/lpm 

FROM ubuntu:latest

WORKDIR /usr/local/bin

COPY --from=build /usr/local/bin/lpm .

RUN apt-get update && \
    apt-get install -y iputils-ping iperf3 tcpdump && \
    apt-get clean
    

ENTRYPOINT ["lpm"]