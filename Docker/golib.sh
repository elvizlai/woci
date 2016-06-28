#!/usr/bin/env bash
echo '-----------go get start-----------'
go get -u github.com/lib/pq
go get -u github.com/garyburd/redigo/redis
go get -u github.com/coreos/etcd/client
go get -u github.com/nsqio/go-nsq

go get -u golang.org/x/crypto/bcrypt
go get -u google.golang.org/grpc
go get -u github.com/golang/protobuf/protoc-gen-go

go get -u github.com/urfave/negroni
go get -u github.com/gorilla/mux
go get -u github.com/gorilla/context
go get -u github.com/dgrijalva/jwt-go
go get -u github.com/pborman/uuid

go get -u github.com/wothing/log
go get -u github.com/wothing/worc
go get -u github.com/wothing/wonaming/etcd

go get -u github.com/elgs/gostrgen
go get -u qiniupkg.com/api.v7/kodo
go get -u qiniupkg.com/x/url.v7
go get -u github.com/ylywyn/jpush-api-go-client
go get -u github.com/pingplusplus/pingpp-go/pingpp
go get -u github.com/tealeg/xlsx

go get -u github.com/smartystreets/assertions
go get -u github.com/smartystreets/goconvey

go get -u github.com/bitly/go-simplejson
echo '-----------go get end-----------'