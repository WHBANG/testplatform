#!/bin/bash
set -e 
set -x 

#usage 
#TAG=latest ./build-docker.sh videotask videoindexworker ...
if [ $# -ne 0 ]; then
    apps=("$@")
else 
    echo "usage:./build-docker.sh [mongo test_platform]"
    exit 1
fi 

VERSION=`date -u '+%Y%m%d'`-`git rev-parse --short HEAD`
# VERSION=latest
if [ -z "${TAG}" ]; then
    TAG=${VERSION}
fi

build_docker(){
    docker build -t reg.supremind.info/wuzhenzhou/$1:$TAG -f docker/$1/Dockerfile .
}

for ((i=0;i<${#apps[@]};i++)); do
    app=${apps[i]}
    build_docker $app
done
