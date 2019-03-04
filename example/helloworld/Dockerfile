FROM alpine:latest
MAINTAINER aaronpei2015@foxmail.com

RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
ADD ./bundles/helloworld /usr/local/bin
ENTRYPOINT ["/usr/local/bin/helloworld"]

EXPOSE 50051
