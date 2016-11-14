FROM golang:1.7.1

WORKDIR /go/src/github.com/caiguanhao/dlfiles

ADD upx-3.91-amd64_linux/upx /usr/bin/upx
ADD build.sh /
ADD *.go ./
ADD vendor ./vendor
ENTRYPOINT ["/build.sh"]
