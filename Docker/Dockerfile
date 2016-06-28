FROM alpine:latest

# glibc for protoc
ENV GLIBC_PKG_VERSION=2.23-r3
RUN apk add --no-cache --update-cache bash curl ca-certificates make tmux git && \
  curl -Lo /etc/apk/keys/sgerrand.rsa.pub "https://raw.githubusercontent.com/sgerrand/alpine-pkg-glibc/master/sgerrand.rsa.pub" && \
  curl -Lo glibc-${GLIBC_PKG_VERSION}.apk "https://github.com/sgerrand/alpine-pkg-glibc/releases/download/${GLIBC_PKG_VERSION}/glibc-${GLIBC_PKG_VERSION}.apk" && \
  curl -Lo glibc-bin-${GLIBC_PKG_VERSION}.apk "https://github.com/sgerrand/alpine-pkg-glibc/releases/download/${GLIBC_PKG_VERSION}/glibc-bin-${GLIBC_PKG_VERSION}.apk" && \
  curl -Lo glibc-i18n-${GLIBC_PKG_VERSION}.apk "https://github.com/sgerrand/alpine-pkg-glibc/releases/download/${GLIBC_PKG_VERSION}/glibc-i18n-${GLIBC_PKG_VERSION}.apk" && \
  apk add glibc-${GLIBC_PKG_VERSION}.apk glibc-bin-${GLIBC_PKG_VERSION}.apk glibc-i18n-${GLIBC_PKG_VERSION}.apk && rm glibc-* && \
  curl -fsSL "http://cdn.1024coder.com/docker-1.11.1.tgz" -o docker-1.11.1.tgz && tar zxf docker-1.11.1.tgz -C /usr/local/bin/ && rm docker-1.11.1.tgz && \
  curl -fsSL "http://cdn.1024coder.com/protoc.tgz" -o protoc.tgz && tar zxf protoc.tgz -C /usr/local/bin/ && rm protoc.tgz && \
  curl -fsSL "http://cdn.1024coder.com/woci.tgz" -o woci.tgz && tar zxf woci.tgz -C /usr/local/bin/ && rm woci.tgz

ENV GO_VERSION=1.6.2
RUN curl -fsSL "https://storage.googleapis.com/golang/go${GO_VERSION}.linux-amd64.tar.gz" -o go.tgz && tar zxf go.tgz -C /usr/local && rm go.tgz

ENV CGO_ENABLED 0
ENV GOROOT /usr/local/go
ENV GOPATH /gopath
ENV PATH $PATH:$GOROOT/bin:$GOPATH/bin

# go get -u
COPY golib.sh /
RUN /golib.sh

CMD ["/bin/bash"]