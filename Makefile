# Ensure GOPATH is set before running build process.
ifeq "$(GOPATH)" ""
  $(error Please set the environment variable GOPATH before running `make`)
endif

temp := $(GOPATH)
export GOPATH = $(temp):/Users/Elvizlai/IdeaProjects/WoCI

linux:
	CGO_ENABLED=0 GOOS=linux go build -o woci

osx:
	CGO_ENABLED=0 go build -o woci_darwin
