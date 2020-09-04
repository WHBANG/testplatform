#!/bin/bash

init(){
    if [ ! -d "./mongo-data/db" ];then
         mkdir -p ./mongo-data/db
        chmod -R 777 mongo-data/db
    fi

    if [ ! -d "./mongo-data/log" ];then
         mkdir -p ./mongo-data/log
        chmod -R 777 mongo-data/log
    fi
}
start(){
    docker-compose up -d
}

restart(){
    docker-compose up --force-recreate -d
}

stop(){
    docker-compose down 
}

if [ ! $# -ne 0 ];then
    echo "usage: $0 [ start | restart | stop ]"
else
    case "$1" in 
        start)
            init
            start
            ;;
        stop)
            stop
            ;;
        restart)
            restart
            ;;
        *)
        echo "Usage: $0 [ start | restart | stop ]"
        exit 1
    esac
fi