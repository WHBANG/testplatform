FROM golang:1.14.3-stretch as builder
# WORKDIR /go/src
# RUN apt-get update && \
#     apt-get  -y install wget

ARG VERSION=dev
ENV GO111MODULE=on
ENV GOSUMDB=off
ARG GOPATH=/go
ARG GOPROXY=https://goproxy.cn,direct
ARG GOPRIVATE=git.supremind.info

ARG VERSION=dev
COPY ./ /go/src/git.supremind.info/testplatform/
# RUN GOOS=linux GOARCH=amd64 go install -tags netgo  git.supremind.info/testplatform/app/test_platform
RUN cd /go/src/git.supremind.info/testplatform/ && \
    GOOS=linux GOARCH=amd64 go build -tags netgo -o /go/bin/test_platform  ./app/test_platform

FROM alpine

RUN apk add --no-cache curl  busybox-extras  wget  tzdata && \
    ln -sf /usr/share/zoneinfo/Asia/Shanghai  /etc/localtime && \
    echo "Asia/Shanghai" > /etc/timezone

# http server port 
EXPOSE 80 

RUN mkdir -p /workspace/config 

COPY --from=builder /go/bin/test_platform  /workspace/test_platform
COPY --from=builder /go/src/git.supremind.info/testplatform/app/test_platform/  /workspace/config/

WORKDIR /workspace

CMD /workspace/test_platform -f /workspace/config/service.conf
