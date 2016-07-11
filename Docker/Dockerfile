FROM alpine:latest

# build and runtime essential
RUN apk add --no-cache --update-cache ca-certificates bash curl make tmux go git

# glibc for protoc
RUN GLIBC_PKG_VERSION="2.23-r3" \
    && curl -Lo /etc/apk/keys/sgerrand.rsa.pub "https://raw.githubusercontent.com/sgerrand/alpine-pkg-glibc/master/sgerrand.rsa.pub" \
    && curl -Lo glibc-${GLIBC_PKG_VERSION}.apk "https://github.com/sgerrand/alpine-pkg-glibc/releases/download/${GLIBC_PKG_VERSION}/glibc-${GLIBC_PKG_VERSION}.apk" \
    && curl -Lo glibc-bin-${GLIBC_PKG_VERSION}.apk "https://github.com/sgerrand/alpine-pkg-glibc/releases/download/${GLIBC_PKG_VERSION}/glibc-bin-${GLIBC_PKG_VERSION}.apk" \
    && curl -Lo glibc-i18n-${GLIBC_PKG_VERSION}.apk "https://github.com/sgerrand/alpine-pkg-glibc/releases/download/${GLIBC_PKG_VERSION}/glibc-i18n-${GLIBC_PKG_VERSION}.apk" \
    && apk add --no-cache glibc-${GLIBC_PKG_VERSION}.apk glibc-bin-${GLIBC_PKG_VERSION}.apk glibc-i18n-${GLIBC_PKG_VERSION}.apk \
    && /usr/glibc-compat/bin/localedef --force --inputfile POSIX --charmap UTF-8 C.UTF-8 || true \
    && echo "export LANG=C.UTF-8" > /etc/profile.d/locale.sh \
    && apk del glibc-i18n \
    && rm glibc-*

# docker protoc
RUN curl -sSL -O "http://cdn.1024coder.com/docker-1.11.2.tgz" \
    && tar zxf docker-1.11.2.tgz -C /usr/local/bin/ \
    && rm docker-1.11.2.tgz \
    && curl -fsSL "http://cdn.1024coder.com/protoc-3.0.0-beta-3.tgz" -o protoc.tgz \
    && tar zxf protoc.tgz -C /usr/local/bin/ \
    && rm protoc.tgz

ENV CGO_ENABLED 0
ENV GOPATH /gopath
ENV PATH $PATH:$GOPATH/bin

# woci
COPY woci /usr/local/bin/

# go get -u
COPY golib.sh /
RUN /golib.sh

CMD ["/bin/bash"]