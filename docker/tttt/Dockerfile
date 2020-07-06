FROM golang:1.14.3-stretch as builder
WORKDIR /go/src
# RUN apt-get update && \
#     apt-get  -y install wget

ARG VERSION=dev
ENV GO111MODULE=on
# ENV GOSUMDB=off
ARG GOPATH=/go
ARG GOPROXY=https://goproxy.cn,direct
ARG GOPRIVATE=git.supremind.info

ARG VERSION=dev
COPY ./ /go/src/git.supremind.info/testplatform/

