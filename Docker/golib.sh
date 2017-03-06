#!/usr/bin/env bash
echo '-----------go get start-----------'
go get github.com/golang/protobuf/protoc-gen-go

go get -d github.com/lib/pq
go get -d gopkg.in/jackc/pgx.v2
go get -d github.com/garyburd/redigo/redis
go get -d gopkg.in/mgo.v2
go get -d github.com/coreos/etcd/client
go get -d github.com/nsqio/go-nsq

go get -d golang.org/x/crypto/bcrypt
go get -d google.golang.org/grpc

go get -d github.com/urfave/negroni
go get -d github.com/julienschmidt/httprouter
go get -d github.com/dgrijalva/jwt-go
go get -d github.com/pborman/uuid
go get -d github.com/googollee/go-socket.io

go get -d github.com/wothing/log
go get -d github.com/wothing/worc
go get -d github.com/wothing/worpc
go get -d github.com/wothing/wonaming/etcd
go get -d github.com/wothing/timing

go get -d github.com/elgs/gostrgen
go get -d qiniupkg.com/api.v7
go get -d github.com/ylywyn/jpush-api-go-client
go get -d github.com/pingplusplus/pingpp-go/pingpp
go get -d github.com/tealeg/xlsx
go get -d github.com/mozillazg/go-pinyin

go get -d github.com/PuerkitoBio/goquery
go get -d github.com/axgle/mahonia


go get -d github.com/smartystreets/goconvey

go get -d github.com/bitly/go-simplejson
echo '-----------go get end-----------'