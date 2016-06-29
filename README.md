# woci

### Usage

#### step0
```
docker pull quay.io/coreos/etcd
docker pull nsqio/nsq
docker pull postgres
docker pull redis
```

#### step1
```
docker pull index.tenxcloud.com/sdrzlyz/woci:1.1
```
#### step2
```
docker network create test
```

#### step3
开启docker in docker
按需编辑woci.json与run.sh，执行run.sh; 将其与17mei项目挂载。
```
#!/usr/bin/env bash
curDir=$(cd `dirname $0`; pwd)

docker run -it --rm --net=test -v /var/run/docker.sock:/var/run/docker.sock -v app:/app -v $curDir/woci.json:/woci.json -v /Users/Elvizlai/IdeaProjects/17Mei/src/github.com/wothing/17mei:/gopath/src/github.com/wothing/17mei index.tenxcloud.com/sdrzlyz/woci:1.1
```

#### step4
```
woci all
```

###Commands
```
NAME:
   woci - make coding joyful!

USAGE:
   woci [global options] command [command options] [arguments...]

VERSION:
   0.1

AUTHOR(S):
   elvizlai

COMMANDS:
     all, a          Evil Mode.👿
     init, i         Init etcd, postgres, redis, nsq
     build, b        Build all app and start
     rebuild, r, re  Rebuild app. split by ',' if multiple
     test, t         Run test case
     clean, c        Clean all app and data
     logs, l         Log by app name
     help, h         Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --debug, -d             debug mode
   --config file, -c file  load config from file (default: "/woci.json")
   --help, -h              show help
   --version, -v           print the version
```

###黑魔法总是写在最后

```
woci build -o appway  //只编译,不运行.一般用于编译检查
woci rebuild appway,interway  //重新编译并运行多个模块,当然啦 woci rebuild appway interway 也是可行的
woci rebuild all  //重新编译并运行所有模块
woci test wechat,version  //单独测试一个或多个模块
woci logs all  //输出所有模块的日志
```