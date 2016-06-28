#!/usr/bin/env bash
curDir=$(cd `dirname $0`; pwd)

if [ ! -d "$curDir/17mei" ]; then
    git clone git@github.com:ElvizLai/17mei.git
else
    pushd $curDir/17mei
    git pull
fi

docker run -it --rm --net=test -v /var/run/docker.sock:/var/run/docker.sock -v app:/app -v $curDir/woci.json:/woci.json -v $curDir/17mei:/gopath/src/github.com/wothing/17mei woci