#!/usr/bin/env bash
echo '-----------go get start-----------'
go get github.com/golang/protobuf/protoc-gen-go

go get -d github.com/lib/pq
go get -d github.com/garyburd/redigo/redis
go get -d github.com/coreos/etcd/client
go get -d github.com/nsqio/go-nsq

go get -d golang.org/x/crypto/bcrypt
go get -d google.golang.org/grpc

go get -d github.com/urfave/negroni
go get -d github.com/gorilla/mux
go get -d github.com/gorilla/context
go get -d github.com/dgrijalva/jwt-go
go get -d github.com/pborman/uuid

go get -d github.com/wothing/log
go get -d github.com/wothing/worc
go get -d github.com/wothing/wonaming/etcd

go get -d github.com/elgs/gostrgen
go get -d qiniupkg.com/api.v7/kodo
go get -d qiniupkg.com/x/url.v7
go get -d github.com/ylywyn/jpush-api-go-client
go get -d github.com/pingplusplus/pingpp-go/pingpp
go get -d github.com/tealeg/xlsx

go get -d github.com/smartystreets/assertions
go get -d github.com/smartystreets/goconvey

go get -d github.com/bitly/go-simplejson
echo '-----------go get end-----------'