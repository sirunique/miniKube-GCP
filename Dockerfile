FROM golang:latest

RUN mkdir  /build
WORKDIR    /build

RUN export GO111MODULE=on
RUN go get github.com/sirunique/miniKube-GCP/
RUN cd /build && git clone https://github.com/sirunique/miniKube-GCP.git

RUN cd /build/miniKube-GCP && go build

EXPOSE 8080

ENTRYPOINT [ "/build/miniKube-GCP" ]