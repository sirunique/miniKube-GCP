FROM golang:latest

RUN mkdir  /build
WORKDIR    /build

RUN export GO111MODULE=on
RUN go get github.com/sirunique/miniKube-GCP/main