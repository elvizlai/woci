# woci

### Usage
woci.yaml使用YAML格式撰写。

#### 保留关键字

[TRACER] - 预留

[KEY] - alias名称

[PLACEHOLDER] - param名称

#### ENV

ENV用于变量定义，定义后的变量可在当前ENV或后续生命周期中传递。变量是顺序随意。使用$+变量名称调用定义的变量。

！！！调用未定义的变量不会报错。

#### woci 的生命周期

INIT --> BEFORE --> BUILD --> START --> TEST --> CLEAN

为空则略过相应的步骤。

BUILD 与 START 步骤较为特殊，可使用字符串数组或者数组对象。数组对象支持前缀alias，可使用[KEY]获取对应的alias占位，方便单步调试。



###Commands
```
NAME:
   woci - make coding joyful!

USAGE:
   woci [global options] command [command options] [arguments...]

VERSION:
   2.0, go1.7.1

AUTHOR(S):
   elvizlai

COMMANDS:
     all, a    Evil Mode.👿
     init, i   initial stage
     build, b  build stage
     start, s  start stage
     test, t   test stage
     clean, c  clean stage
     plugin    plugin center
     help, h   Shows a list of commands or help for one command

GLOBAL OPTIONS:
   -c file        load config from file (default: "/woci.yaml")
   -f             run in force mode
   -d             run in debug mode
   --help, -h     show help
   --version, -v  print the version
```
