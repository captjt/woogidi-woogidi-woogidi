FROM alpine:3.5

MAINTAINER Jordan Taylor <jtaylor007.jt@gmail.com>

LABEL name="woogidi-woogidi-woogidi" \

COPY woogidi-woogidi-woogidi .

ENTRYPOINT ["./woogidi-woogidi-woogidi"]
