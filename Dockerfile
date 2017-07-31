FROM golang:1.8.0

RUN apt-get -q update && \
    apt-get install --no-install-recommends -y --force-yes -q \
    emacs \
    locales \
    task-japanese && \
    echo "ja_JP.UTF-8 UTF-8" >> /etc/locale.gen && \
    locale-gen

ENV LANG=ja_JP.UTF-8

RUN go get github.com/mattn/gom
ADD Gomfile /go/src/
WORKDIR /go/src/
RUN gom install
RUN go get -u gopkg.in/godo.v2/cmd/godo

WORKDIR /go/src/myapp

ENV PATH $PATH:$GOPATH/bin:/go/src/vendor/bin

EXPOSE 1323
